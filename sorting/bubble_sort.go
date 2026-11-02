package sorting

import "slices"

// [easy] BubbleSort returns xs sorted in ascending order.
//
// Repeatedly sweep the slice, swapping adjacent out-of-order pairs.
// Steps:
//  1. Each pass walks the unsorted part of the slice from front to
//     back, comparing each pair of neighbors and swapping them when
//     they are out of order.
//  2. After a full pass, the largest item of the unsorted part has
//     been carried to its end and is in its final place; the next pass
//     can ignore it. (If you track this, each pass is one item shorter
//     than the last; if you do not, the sort is still correct, just
//     does redundant work.)
//  3. Stop as soon as a whole pass performs no swaps — nothing was out
//     of order, so the slice is sorted. This early exit is what gives
//     the linear best case on already-sorted input; without it every
//     input pays for the full quadratic number of passes.
//
// Notes:
//   - The early exit (step 3) is what gives the O(n) best case on sorted input.
//   - Stable: equal items keep their relative order.
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
