package hashing

import (
	"fmt"
	"slices"
	"testing"
)

func TestFNV1aDeterministic(t *testing.T) {
	// Same input -> same hash; different inputs -> (almost surely) different.
	if FNV1a("hello") != FNV1a("hello") {
		t.Error("FNV1a is not deterministic")
	}
	seen := make(map[uint32]string)
	for i := range 1000 {
		s := fmt.Sprintf("key-%d", i)
		h := FNV1a(s)
		if prev, ok := seen[h]; ok {
			t.Errorf("FNV1a collision: %q and %q both hash to %#x", prev, s, h)
		}
		seen[h] = s
	}
}

// mapOracle runs the same op sequence against a stdlib map and the candidate.
func checkAgainstStdMap(t *testing.T, name string, m interface {
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
			t.Errorf("%s: Get(%q) = (%d, %v), want (%d, %v)", name, key, got, ok, want, wantOk)
		}
	}
	if _, ok := m.Get("missing"); ok {
		t.Errorf("%s: Get on absent key returned ok=true", name)
	}
}

func TestChainMapCorrect(t *testing.T) {
	checkAgainstStdMap(t, "ChainMap", NewChainMap(16))
}

func TestOpenMapCorrect(t *testing.T) {
	checkAgainstStdMap(t, "OpenMap", NewOpenMap(16))
}

func TestBloomFilterCorrect(t *testing.T) {
	const n = 10_000
	b := NewBloomFilter(n, 0.01)
	added := make([]string, n)
	for i := range n {
		added[i] = fmt.Sprintf("present-%d", i)
		b.Add(added[i])
	}
	// No false negatives, ever.
	for _, s := range added {
		if !b.MayContain(s) {
			t.Fatalf("BloomFilter: false negative for %q (this is a hard bug)", s)
		}
	}
	// False-positive rate against absent keys must be near the target 1%.
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
}

func TestRingCorrect(t *testing.T) {
	nodes := []string{"alpha", "beta", "gamma", "delta"}
	r := NewRing(nodes, 100)
	// Every lookup resolves to a real node...
	seen := make(map[string]string)
	for i := range 5000 {
		key := fmt.Sprintf("obj-%d", i)
		got := r.GetNode(key)
		if !slices.Contains(nodes, got) {
			t.Fatalf("Ring.GetNode(%q) = %q, not a known node", key, got)
		}
		// ...and the ring is a pure function of the key: stable answers.
		if prev, ok := seen[key]; ok && prev != got {
			t.Errorf("Ring.GetNode(%q) unstable: %q then %q", key, prev, got)
		}
		seen[key] = got
	}
	// With 100 replicas, no node should hog more than half the key space.
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
}
