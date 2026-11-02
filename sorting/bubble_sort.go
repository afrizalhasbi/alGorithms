package sorting

// [easy] BubbleSort returns xs sorted in ascending order.
//
// Repeatedly sweep the slice, swapping adjacent out-of-order pairs.
// Steps:
// 1. Copy xs so the caller's slice is not mutated.
// 2. Compare items i and i+1 across the slice; swap if out of order.
// 3. After each full pass, the largest remaining item is in its final place.
// 4. If a pass performs no swaps, the slice is sorted: stop early.
//
// Notes:
//   - The early exit (step 4) is what gives the O(n) best case on sorted input.
//   - Stable: equal items keep their relative order.
//
// Expected: best O(n), average/worst O(n^2)
// Input: not mutated
func BubbleSort(xs []int) []int {
	// TODO: implement
	return nil
}
