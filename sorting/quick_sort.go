package sorting

// [medium] QuickSort returns xs sorted in ascending order.
//
// Partition around a pivot, then recurse on the two sides.
// Steps:
//  1. Choose a pivot using a deterministic pseudo-random index
//     (seed a rand.Rand once per call chain).
//  2. Partition: items < pivot to the left, items >= pivot to the right;
//     the pivot lands in its final position.
//  3. Recurse on both sides.
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
