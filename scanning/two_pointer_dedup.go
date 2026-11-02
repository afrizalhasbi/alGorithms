package scanning

// [medium] TwoPointerDedup removes duplicates from the sorted slice xs in place and
// returns the length k of the deduplicated prefix.
//
// One slow pointer marking the end of the unique prefix, one fast pointer
// scanning ahead.
// Steps:
//  1. An empty slice is trivially deduplicated: return zero.
//  2. Think of the slice as two regions: a front region holding the
//     unique items kept so far, and the rest being examined. At the
//     start the front region holds just the first item (a single item
//     is always unique).
//  3. Examine the remaining items one at a time, in order. An item
//     equal to the last kept item is a duplicate — skip it, leaving
//     whatever garbage sits in its place. A different item is kept:
//     copy it into the slot right after the front region and let the
//     front region grow by one.
//  4. When every item has been examined, the front region is the
//     deduplicated slice; return its length. Nothing past it is
//     meaningful.
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
