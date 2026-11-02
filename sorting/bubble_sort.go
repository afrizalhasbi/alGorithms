package sorting

import "slices"

// [easy] BubbleSort returns xs sorted in ascending order.
//
// Background. A pass walks the slice from front to back, comparing
// each neighboring pair and swapping the two when they are out of
// order. One pass does not sort the slice, but it guarantees that
// the largest item has been moved at least one step further back
// per comparison, so by the end of a pass the largest item of the
// whole slice has reached the back. That item is then in its final
// position for good.
//
// Sorting is just repeating passes, each one able to ignore one more
// item at the back because it is already final. The slice can be
// sorted before the passes run out, and there is an easy way to
// notice: if a whole pass performs no swaps, nothing was out of
// order, so the slice is sorted and the algorithm can stop. That
// early exit is the difference between a best case of O(n) on
// already-sorted input and paying the full quadratic number of
// passes every time.
//
// Steps:
//  1. Each pass: walk the unsorted part of the slice, compare each
//     pair of neighbors, and swap them when they are out of order.
//  2. After a pass, the largest item of the unsorted part sits at its
//     end, final. Shrink the unsorted part by that one item.
//  3. If a pass performed no swaps, stop. Otherwise repeat while the
//     unsorted part has items in it.
//
// Notes:
//   - Stable: equal neighbors are never swapped, so equal items keep
//     their relative order.
//   - The test measures both the quadratic average case and the
//     linear best case, so the early exit is required, not optional.
//
// Expected: best O(n), average/worst O(n^2)
// Input: mutated in place
func BubbleSort(xs []int) []int {
	for !slices.IsSorted(xs) {
		for i := range len(xs) - 1 {
			if xs[i] > xs[i+1] {
				tmp := xs[i+1]
				xs[i+1] = xs[i]
				xs[i] = tmp
			}
		}
	}
	return xs
}
