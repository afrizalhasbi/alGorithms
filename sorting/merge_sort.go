package sorting

// [medium] MergeSort returns xs sorted in ascending order.
//
// Divide and conquer: split, sort halves, merge.
// Steps:
//  1. Base case: a slice with at most one item is sorted; return it
//     unchanged. This is what stops the recursion.
//  2. Otherwise split the slice into two halves, sort each half by the
//     same method, and merge the two sorted halves back together.
//  3. Merging: the right half is left where it is; the left half is
//     copied to a scratch area first, because the merge writes over the
//     positions the left half still needs. Then repeatedly compare the
//     front item of the left half with the front item of the right half
//     and move the smaller one into the next output position. On a tie,
//     take from the left half — this is what makes the sort stable.
//  4. When one half runs out: if the right half ran out, copy the rest
//     of the left half over; if the left half ran out, do nothing, since
//     the remaining right-half items already sit in their final places.
//
// Notes:
//   - O(n log n) guaranteed: input shape does not matter, so only test the
//     average case.
//   - Merge needs O(n) auxiliary space.
//   - Stable if you prefer the left half on ties.
//
// Expected: best/average/worst O(n log n)
// Input: mutated in place
func MergeSort(xs []int) []int {
	// TODO: implement
	return nil
}
