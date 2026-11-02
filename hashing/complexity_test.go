package hashing

import (
	"fmt"
	"testing"

	"alGorithms/internal/complexity"
)

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	if p < lo || p > hi {
		t.Errorf("%s: measured exponent %.2f, expected in [%.2f, %.2f]", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

func TestFNV1a(t *testing.T) {
	// Known FNV-1a test vectors (32-bit).
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
}

func TestChainMapComplexity(t *testing.T) {
	// n inserts + n lookups; with capped load factor this stays linear.
	assertExponent(t, "chain map O(1) amortized",
		complexity.FitExponent(1000, 5, func(n int) {
			m := NewChainMap(16)
			for i := range n {
				m.Put(fmt.Sprintf("key-%d", i), i)
			}
			for i := range n {
				m.Get(fmt.Sprintf("key-%d", i))
			}
		}), 0.8, 1.4)
}

func TestOpenMapComplexity(t *testing.T) {
	assertExponent(t, "open map O(1) amortized",
		complexity.FitExponent(1000, 5, func(n int) {
			m := NewOpenMap(16)
			for i := range n {
				m.Put(fmt.Sprintf("key-%d", i), i)
			}
			for i := range n {
				m.Get(fmt.Sprintf("key-%d", i))
			}
		}), 0.8, 1.4)
}

func TestBloomFilterComplexity(t *testing.T) {
	assertExponent(t, "bloom filter O(1) per op",
		complexity.FitExponent(1000, 5, func(n int) {
			b := NewBloomFilter(n, 0.01)
			for i := range n {
				b.Add(fmt.Sprintf("key-%d", i))
			}
			for i := range n {
				b.MayContain(fmt.Sprintf("key-%d", i))
			}
		}), 0.8, 1.4)
}

func TestRingLookupComplexity(t *testing.T) {
	// Lookups are binary searches over ring positions: exponent ~0.
	const probes = 20_000
	assertExponent(t, "ring lookup O(log n)",
		complexity.FitExponent(64, 5, func(n int) {
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
