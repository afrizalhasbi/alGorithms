package sorting

// [easy] SelectionSort returns xs sorted in ascending order.
//
// Repeatedly select the smallest remaining item and place it next.
// Steps:
//  1. Treat the front of the slice as a sorted prefix and the rest as
//     the unsorted remainder; the prefix starts empty.
//  2. Each round, find the smallest item in the unsorted remainder by
//     scanning the whole remainder once. If several items tie for the
//     smallest, keep the first one seen.
//  3. Swap that smallest item with the first slot of the remainder,
//     which extends the sorted prefix by one item.
//  4. Repeat while the remainder holds more than one item. Once only
//     one item remains, it is necessarily the largest of what is left
//     and needs no scan — stop there.
//
// Notes:
//   - Exactly O(n^2) comparisons on every input shape: there is no best case.
//   - Minimizes swaps (at most n-1), useful when writes are expensive.
//   - Not stable in this swap-based form.
//
// Expected: best/average/worst O(n^2)
// Input: mutated in place
func SelectionSort(xs []int) []int {
	// TODO: implement
	return nil
}
