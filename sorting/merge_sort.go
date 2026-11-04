package sorting

// [medium] MergeSort returns xs sorted in ascending order.
//
// Background. This is a divide and conquer algorithm: instead of
// sorting the whole slice directly, sort two smaller slices and
// combine the results. The combining step is the merge, and it is
// the part that does the real work. Merging takes two slices that
// are each already sorted and produces one sorted slice holding all
// of their items, by repeatedly comparing the front items of the two
// and moving the smaller one into the output. Because each
// comparison retires one item, merging two slices takes time
// proportional to their combined length.
//
// A sort is stable when items that compare equal keep the order they
// had in the input. Merge sort can easily be made stable: when the
// two front items are equal, take the one from the left half, since
// everything in the left half started out before everything in the
// right half.
//
// Definitions:
//  - merge: combining two already-sorted slices into one sorted slice.
//     It is the only place items move; the sort itself is just
//     recursion plus merges.
//  - left half / right half: the two pieces of the split. "Left" means
//     earlier positions in the original slice, which is what stability
//     depends on.
//  - front item: the first not-yet-consumed item of a half during the
//     merge. Each comparison consumes exactly one front item.
//  - buffer: the auxiliary slice in step 3. Only the left half is
//     copied into it; the right half stays in place.
//
// Steps:
//  1. Base case: a slice with one or zero items is sorted, so return
//     it unchanged. This is what stops the recursion.
//  2. Otherwise split the slice into a left and a right half, sort
//     each half by the same method, then merge the two sorted halves
//     back into the original slice.
//  3. Merging in place needs one extra buffer holding the left half,
//     because the merge writes into positions the left half has not
//     been read yet. The right half needs no copy.
//  4. Merge until one of the halves runs out. If the right half runs
//     out first, copy the rest of the left half over; if the left
//     half runs out first, the remaining right-half items are already
//     in their final positions and nothing needs to be done.
//
// Notes:
//   - O(n log n) guaranteed: input shape does not matter, so only
//     test the average case.
//   - Merge needs O(n) auxiliary space.
//
// Expected: best/average/worst O(n log n)
// Input: mutated in place
func MergeSort(xs []int) []int {
	// TODO: implement
	return nil
}
