package hashing

// [hard] BloomFilter is a probabilistic set that can say "definitely not present"
// or "maybe present", using a fixed bit budget.
//
// Background. A Bloom filter is a set that gives up exactness in one
// direction only. Asking whether an item was added can produce a
// false positive (the answer is "maybe present" for an item that was
// never added), but never a false negative (an item that was added
// is always reported as present). The reason is structural: adding
// sets bits and nothing ever clears a bit, so any added item keeps
// finding its own bits set forever.
//
// The filter is a plain array of bits, all clear at the start. An
// item does not claim one bit but several: several hash functions
// each pick one bit, and adding the item sets all of them. A query
// recomputes the same positions and is only willing to say "present"
// if every one of those bits is set. A single clear bit is proof the
// item was never added, because if it had been, that bit would have
// been set.
//
// False positives happen when all of an item's bits happen to be set
// by a mixture of other items. The interesting part is that the rate
// of that accident is predictable from the size of the bit array,
// the number of items, and the number of hash functions, which is
// what the sizing formula encodes. More hash functions mean fewer
// false positives but a faster fill-up; there is an optimum, and the
// formula gives it.
//
// Deriving several hash functions from one base hash is standard
// (the Kirsch-Mitzenmacher trick): take one hash of the item, split
// or mix it into two values, and combine them with the function's
// index to produce each variant. No separate hash functions needed.
//
// Steps:
//  1. Construction: from the expected item count and the target
//     false-positive rate, compute the bit array length and the
//     number of hash functions with the standard formula. Allocate
//     the bits, all clear.
//  2. Add: compute the base hash of the item, derive the k variant
//     hashes from it, reduce each into range by taking the remainder
//     against the bit count, and set all k bits.
//  3. MayContain: derive the same k positions. If even one bit is
//     clear, return false. If all are set, return true.
//
// Notes:
//   - The test adds ten thousand items, then checks that none of
//     them is ever reported absent (a false negative is a hard bug),
//     and that the fraction of false positives on absent items lands
//     near the 1% target.
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
