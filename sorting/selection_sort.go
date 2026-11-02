package sorting

// [easy] SelectionSort returns xs sorted in ascending order.
//
// Repeatedly select the smallest remaining item and place it next.
// Steps:
// 1. Copy xs so the caller's slice is not mutated.
// 2. Scan the unsorted suffix for its minimum.
// 3. Swap that minimum with the first slot of the suffix.
// 4. The sorted prefix grows by one; repeat until the suffix is empty.
//
// Notes:
//   - Exactly O(n^2) comparisons on every input shape: there is no best case.
//   - Minimizes swaps (at most n-1), useful when writes are expensive.
//   - Not stable in this swap-based form.
//
// Expected: best/average/worst O(n^2)
// Input: not mutated
func SelectionSort(xs []int) []int {
	// TODO: implement
	return nil
}
