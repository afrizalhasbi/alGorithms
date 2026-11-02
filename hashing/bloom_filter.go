package hashing

// [hard] BloomFilter is a probabilistic set that can say "definitely not present"
// or "maybe present", using a fixed bit budget.
//
// Multiple hash functions set bits; membership checks only read.
// Steps:
//  1. Construction: choose the number of bits from the expected item
//     count and the target false-positive rate using the standard
//     formula, and from that the number of hash functions per item.
//     Allocate the bit array, everything clear.
//  2. Adding: derive the needed hash values from a single base hash
//     of the item — combine the base hash with the function's number
//     to get each independent variant (two different hashes can be
//     mixed with the function index to produce all of them; this is
//     the standard trick and avoids k separate hash functions). Each
//     derived hash picks one bit, reduced into range by the remainder;
//     set every one of those bits.
//  3. Membership: recompute exactly the same set of bit positions for
//     the query item. If every one of them is set, answer "maybe
//     present"; if even one is clear, the item cannot have been added
//     (adding never clears a bit), so answer "definitely not present".
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
