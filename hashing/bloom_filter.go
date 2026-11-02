package hashing

// [hard] BloomFilter is a probabilistic set that can say "definitely not present"
// or "maybe present", using a fixed bit budget.
//
// Multiple hash functions set bits; membership checks only read.
// Steps:
//  1. NewBloomFilter(n, fpRate): size the bit array m from the standard
//     estimate m = -n*ln(p) / (ln2)^2 and k = (m/n)*ln2 hash functions.
//     Allocate the bit slice, all false.
//  2. Add(s): for each of the k hash functions, derive an independent hash
//     from FNV1a (e.g. combine FNV1a(s) with a salt per function, or split
//     one 32-bit hash into two halves and mix: h_i = h1 + i*h2) and set
//     bit (h_i mod m).
//  3. MayContain(s): recompute the k positions; return false if any bit is
//     clear, true otherwise.
//
// Notes:
//   - False positives are inherent (bits are shared); false negatives are
//     impossible — a bit set by Add is never cleared.
//   - Deriving k hashes from one base hash (the h1 + i*h2 trick) is a real
//     technique (Kirsch–Mitzenmacher); no crypto needed.
//
// Expected: O(k) = O(1) per op
// Input: not mutated
type BloomFilter struct {
	// TODO: fields
}

func NewBloomFilter(n int, fpRate float64) *BloomFilter {
	// TODO: implement
	return nil
}

func (b *BloomFilter) Add(s string) {
	// TODO: implement
}

func (b *BloomFilter) MayContain(s string) bool {
	// TODO: implement
	return false
}
