package sorting

// [medium] QuickSort returns xs sorted in ascending order.
//
// Background. Pick one item of the slice and call it the pivot. A
// partition rearranges the slice so that every item smaller than the
// pivot ends up before it and every item larger ends up after it.
// After one partition the pivot is somewhere between the two groups,
// and everything on each side of it still needs sorting, but no item
// can ever cross the pivot again. Sorting is then just the same
// operation applied to each side, and each side is smaller than the
// whole, so the recursion stops.
//
// The pivot choice decides everything. If the pivot always lands
// near the middle, the halves halve each time and the total work is
// about n log n. If the pivot is always the smallest or largest
// item (which happens deterministically on sorted input when the
// pivot is always taken from the front), one side is empty and the
// work is quadratic. Choosing the pivot at random makes that input
// harmless.
//
// This partition is the Hoare scheme: two cursors start at the two
// ends and move inward. The left cursor stops at the first item that
// is at least the pivot; the right cursor stops at the first item
// that is at most the pivot. When both are stopped, those two items
// are on the wrong sides of each other, so they are swapped. The
// cursors resume. When they cross, the partition is done and the
// position where they crossed splits the slice into the two sides to
// recurse on.
//
// Definitions:
//  - pivot: the item chosen in step 2. Partitions and comparisons use
//     its value; the item itself moves during partitioning like any
//     other.
//  - partition: one application of the two-cursor scheme, rearranging
//     a range so that no item larger than the pivot value sits before
//     the split and none smaller sits after it.
//  - cursors: the two moving positions that scan the range inward.
//     The left one advances past items smaller than the pivot value;
//     the right one advances past items larger. Each stops at the
//     first item that belongs on the other side.
//  - split (where the cursors cross): the boundary returned by the
//     partition. It is NOT necessarily the pivot's final position —
//     unlike the Lomuto scheme, no item is guaranteed to be in its
//     final position after one partition. The two recursion ranges
//     are divided at the split.
//
// Steps:
//  1. Exit condition: a range of one or zero items is sorted.
//  2. Choose the pivot at random, using one random generator created
//     once and shared by the whole recursion, so results are
//     reproducible.
//  3. Partition the range with the two-cursor scheme described above,
//     comparing against the pivot's value.
//  4. Recurse on the two sides of the split.
//
// Notes:
//   - Sorted or all-equal input is the classic worst case for a
//     naive pivot (front or back item). Random pivots fix the sorted
//     case; all-equal input still stresses the partition, which is
//     why it is one of the tested shapes.
//   - Partitioning happens inside the slice, so no extra allocation
//     is needed.
//   - Not stable.
//
// Expected: best/average O(n log n), worst O(n^2) (rare with random pivots)
// Input: mutated in place
func QuickSort(xs []int) []int {
	// TODO: implement
	return nil
}
