package trees

// [medium] Fenwick is a Fenwick tree, also called a binary indexed tree.
// It supports point additions and prefix-sum queries over a fixed-length
// slice of integers.
//
// Background. A Fenwick tree stores partial sums in an array indexed
// from one. The entry at index j covers a block of elements whose length
// is the lowest set bit of j. Adding to an element updates one entry per
// block that contains it, and the next block index is found by adding
// that lowest set bit. A prefix sum reads one entry per block that ends at
// or before the requested position, and the previous block index is found
// by subtracting the lowest set bit. Both walks visit O(log n) entries.
// A range sum is the difference of two prefix sums.
//
// Definitions:
//   - external index: an index in 0..n-1, as used by the caller.
//   - internal index: an external index plus one, so internal indices run
//     1..n.
//   - lowest set bit: for a positive internal index j, the value j & -j.
//     It is the length of the block that entry j summarizes.
//   - prefix sum: the sum of the first i elements, at external indices
//     0..i-1. PrefixSum(0) is zero and PrefixSum(n) is the total.
//   - range sum: the sum over the half-open external range [l, r).
//
// Steps:
//  1. Construction: allocate an internal array of n+1 zeros.
//  2. Add(i, delta): start at internal index j = i+1. While j is at most
//     n, add delta to the internal entry at j, then set j to j plus its
//     lowest set bit.
//  3. PrefixSum(i): start at internal index j = i. While j is greater
//     than zero, add the internal entry at j to the running sum, then set
//     j to j minus its lowest set bit. Return the sum.
//  4. RangeSum(l, r): return PrefixSum(r) minus PrefixSum(l).
//
// Notes:
//   - All values start at zero.
//   - The caller's indices are zero-based. Add requires 0 <= i < n.
//     PrefixSum requires 0 <= i <= n. RangeSum requires
//     0 <= l <= r <= n.
//   - Unlike a segment tree, the Fenwick tree supports adding to an
//     element, not assigning a new value. An assignment is expressed as
//     adding the difference between the new and old values.
//
// Expected: O(log n) per operation, O(n) space
// Input: mutated in place
type Fenwick struct {
	// TODO: fields
}

// NewFenwick returns a tree of n zeros.
func NewFenwick(n int) *Fenwick {
	// TODO: implement
	return nil
}

// Add adds delta to the element at index i.
func (f *Fenwick) Add(i, delta int) {
	// TODO: implement
}

// PrefixSum returns the sum of the first i elements, indices 0..i-1.
func (f *Fenwick) PrefixSum(i int) int {
	// TODO: implement
	return 0
}

// RangeSum returns the sum of the half-open range [l, r).
func (f *Fenwick) RangeSum(l, r int) int {
	// TODO: implement
	return 0
}
