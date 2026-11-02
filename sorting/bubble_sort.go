package sorting

import "slices"

// [easy] BubbleSort returns xs sorted in ascending order.
//
// Repeatedly sweep the slice, swapping adjacent out-of-order pairs.
// Steps:
// 1. Sort xs in place, swapping adjacent out-of-order pairs.
// 2. After each full pass, the largest remaining item is in its final place.
// 3. If a pass performs no swaps, the slice is sorted: stop early.
//
// Notes:
//   - The early exit (step 4) is what gives the O(n) best case on sorted input.
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
	// TODO: implement
	return xs
}
