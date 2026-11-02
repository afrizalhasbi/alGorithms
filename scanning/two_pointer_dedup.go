package scanning

// [medium] TwoPointerDedup removes duplicates from the sorted slice xs in place and
// returns the length k of the deduplicated prefix.
//
// One slow pointer marking the end of the unique prefix, one fast pointer
// scanning ahead.
// Steps:
//  1. If xs is empty, return 0.
//  2. Slow pointer k = 1: everything before index k is unique and sorted.
//  3. Fast pointer i scans from 1 to end. When xs[i] != xs[k-1], copy
//     xs[i] into xs[k] and increment k.
//  4. Return k. Items after index k hold leftover garbage — that is fine.
//
// Notes:
//   - This is the canonical "write index" pattern; it shows up everywhere
//     (partitioning, compaction, stable partition variants).
//   - Input must be sorted, otherwise equal items are not adjacent.
//
// Expected: O(n)
// Input: mutated in place (only up to the returned length is meaningful)
func TwoPointerDedup(xs []int) int {
	// TODO: implement
	return 0
}
