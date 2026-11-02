package sorting

// [medium] QuickSort returns xs sorted in ascending order.
//
// Partition around a pivot, then recurse on the two sides.
// Steps:
//  1. Sort a range of the slice as follows. The exit: a range holding
//     zero or one items is already sorted and is returned untouched.
//  2. Choose the pivot: pick a random item of the range, using one
//     random generator seeded once and shared by the whole recursion.
//  3. Partition the range around the pivot's VALUE. Walk two cursors,
//     one from each end, moving inward: the left cursor skips over
//     items smaller than the pivot, the right cursor skips over items
//     larger than it. When both are stuck on an item that belongs on
//     the other side, swap them and continue. When the cursors cross,
//     the partition is done, and everything up to the right cursor's
//     position is at most the pivot, everything after it is at least
//     the pivot. Note the pivot itself is not necessarily at that
//     split point — partitioning around the value, not the item.
//  4. Recurse on the two sides of the split. Each side is strictly
//     smaller than the original range, so the recursion always ends.
//
// Notes:
//   - Sorted or all-equal input is the classic worst case for *naive* pivot
//     choice (first/last element). Random pivots make that input harmless.
//   - Partition in place with two indices moving inward (Hoare-style) to
//     avoid extra allocations.
//
// Expected: best/average O(n log n), worst O(n^2) (rare with random pivots)
// Input: mutated in place
func QuickSort(xs []int) []int {
	// TODO: implement
	return nil
}
