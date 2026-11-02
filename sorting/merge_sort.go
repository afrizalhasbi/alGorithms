package sorting

// [medium] MergeSort returns xs sorted in ascending order.
//
// Divide and conquer: split, sort halves, merge.
// Steps:
// 1. If the slice has 0 or 1 items, it is already sorted.
// 2. Split at the midpoint; sort each half recursively.
// 3. Merge the two sorted halves by repeatedly taking the smaller front item.
//
// Notes:
//   - O(n log n) guaranteed: input shape does not matter, so only test the
//     average case.
//   - Merge needs O(n) auxiliary space.
//   - Stable if you prefer the left half on ties.
//
// Expected: best/average/worst O(n log n)
// Input: not mutated
func MergeSort(xs []int) []int {
	// TODO: implement
	return nil
}
