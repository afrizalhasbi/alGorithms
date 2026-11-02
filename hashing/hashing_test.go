package hashing

import (
	"fmt"
	"slices"
	"testing"

	"alGorithms/internal/complexity"
)

// ---- shared helpers ----

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	// Being faster than expected is fine (measurement noise, machine
	// variance); only being slower than the upper bound fails.
	if p > hi {
		t.Errorf("%s: measured exponent %.2f, expected at most %.2f", name, p, hi)
	} else if p < lo {
		t.Logf("%s: exponent %.2f below expected range [%.2f, %.2f] — faster than expected, ok", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

// mapOracle runs the same op sequence against a stdlib map and the candidate.
func mapOracle(t *testing.T, name string, m interface {
	Put(string, int)
	Get(string) (int, bool)
}) {
	t.Helper()
	ref := make(map[string]int)
	for i := range 2000 {
		key := fmt.Sprintf("key-%d", i%500) // 500 distinct keys, 4 overwrites each
		m.Put(key, i)
		ref[key] = i
	}
	for i := range 500 {
		key := fmt.Sprintf("key-%d", i)
		got, ok := m.Get(key)
		want, wantOk := ref[key]
		if ok != wantOk || got != want {
			t.Fatalf("%s: Get(%q) = (%d, %v), want (%d, %v) (further errors suppressed)", name, key, got, ok, want, wantOk)
		}
	}
	if _, ok := m.Get("missing"); ok {
		t.Errorf("%s: Get on absent key returned ok=true", name)
	}
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestFNV1a(t *testing.T) {
	// Correctness: known FNV-1a test vectors (32-bit).
	vectors := map[string]uint32{
		"":       2166136261,
		"a":      0xe40c292c,
		"foobar": 0xbf9cf968,
	}
	for in, want := range vectors {
		if got := FNV1a(in); got != want {
			t.Errorf("FNV1a(%q) = %#x, want %#x", in, got, want)
		}
	}
	// Determinism: same input -> same hash, no collisions on 1000 keys.
	if FNV1a("hello") != FNV1a("hello") {
		t.Error("FNV1a is not deterministic")
	}
	seen := make(map[uint32]string)
	for i := range 1000 {
		s := fmt.Sprintf("key-%d", i)
		h := FNV1a(s)
		if prev, ok := seen[h]; ok {
			t.Fatalf("FNV1a collision: %q and %q both hash to %#x (further collisions suppressed)", prev, s, h)
		}
		seen[h] = s
	}
	// No complexity test: O(len(s)) per key, no shape dependence.
}

func TestChainMap(t *testing.T) {
	// Correctness.
	mapOracle(t, "ChainMap", NewChainMap(16))
	// Complexity: n inserts + n lookups; with capped load factor, linear.
	assertExponent(t, "O(1) amortized", complexity.FitExponent(1000, 5, func(n int) {
		m := NewChainMap(16)
		for i := range n {
			m.Put(fmt.Sprintf("key-%d", i), i)
		}
		for i := range n {
			m.Get(fmt.Sprintf("key-%d", i))
		}
	}), 0.8, 1.4)
}

func TestOpenMap(t *testing.T) {
	// Correctness.
	mapOracle(t, "OpenMap", NewOpenMap(16))
	// Complexity.
	assertExponent(t, "O(1) amortized", complexity.FitExponent(1000, 5, func(n int) {
		m := NewOpenMap(16)
		for i := range n {
			m.Put(fmt.Sprintf("key-%d", i), i)
		}
		for i := range n {
			m.Get(fmt.Sprintf("key-%d", i))
		}
	}), 0.8, 1.4)
}

func TestBloomFilter(t *testing.T) {
	// Correctness: no false negatives, ever; FP rate near the 1% target.
	const n = 10_000
	b := NewBloomFilter(n, 0.01)
	added := make([]string, n)
	for i := range n {
		added[i] = fmt.Sprintf("present-%d", i)
		b.Add(added[i])
	}
	for _, s := range added {
		if !b.MayContain(s) {
			t.Fatalf("BloomFilter: false negative for %q (this is a hard bug)", s)
		}
	}
	absent := make([]string, n)
	for i := range absent {
		absent[i] = fmt.Sprintf("absent-%d", i)
	}
	pos := 0
	for _, s := range absent {
		if b.MayContain(s) {
			pos++
		}
	}
	rate := float64(pos) / float64(n)
	if rate < 0.002 || rate > 0.04 {
		t.Errorf("BloomFilter false-positive rate = %.4f, expected ~0.01", rate)
	} else {
		t.Logf("false-positive rate %.4f (target 0.01)", rate)
	}
	// Complexity.
	assertExponent(t, "O(1) per op", complexity.FitExponent(1000, 5, func(n int) {
		b := NewBloomFilter(n, 0.01)
		for i := range n {
			b.Add(fmt.Sprintf("key-%d", i))
		}
		for i := range n {
			b.MayContain(fmt.Sprintf("key-%d", i))
		}
	}), 0.8, 1.4)
}

func TestRing(t *testing.T) {
	// Correctness: resolves to real nodes, stable, balanced with replicas.
	nodes := []string{"alpha", "beta", "gamma", "delta"}
	r := NewRing(nodes, 100)
	seen := make(map[string]string)
	for i := range 5000 {
		key := fmt.Sprintf("obj-%d", i)
		got := r.GetNode(key)
		if !slices.Contains(nodes, got) {
			t.Fatalf("Ring.GetNode(%q) = %q, not a known node", key, got)
		}
		if prev, ok := seen[key]; ok && prev != got {
			t.Errorf("Ring.GetNode(%q) unstable: %q then %q", key, prev, got)
		}
		seen[key] = got
	}
	count := make(map[string]int, len(nodes))
	for _, v := range seen {
		count[v]++
	}
	for _, node := range nodes {
		share := float64(count[node]) / float64(len(seen))
		if share > 0.5 {
			t.Errorf("node %q holds %.2f of keys with replicas=100 — virtual nodes broken", node, share)
		} else {
			t.Logf("node %q holds %.2f of keys", node, share)
		}
	}
	// Complexity: lookups are binary searches over ring positions, exponent ~0.
	const probes = 20_000
	assertExponent(t, "O(log n)", complexity.FitExponent(64, 5, func(n int) {
		nodes := make([]string, n)
		for i := range nodes {
			nodes[i] = fmt.Sprintf("node-%d", i)
		}
		r := NewRing(nodes, 1)
		for i := range probes {
			r.GetNode(fmt.Sprintf("key-%d", i))
		}
	}), -0.3, 0.4)
}
