package sorting

// [easy] InsertionSort returns xs sorted in ascending order.
//
// Grow a sorted prefix one item at a time, like sorting a hand of cards.
// Steps:
// 1. Copy xs so the caller's slice is not mutated.
// 2. For each index i, take item xs[i] as the "key".
// 3. Shift every larger item in the sorted prefix xs[0:i] one slot right.
// 4. Drop the key into the gap.
//
// Notes:
//   - Shifting instead of swapping halves the number of writes vs bubble sort.
//   - Very fast on nearly-sorted input; still O(n^2) on random input.
//   - Stable.
//
// Expected: best O(n), average/worst O(n^2)
// Input: not mutated
func InsertionSort(xs []int) []int {
	// TODO: implement
	return nil
}
