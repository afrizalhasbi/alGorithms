package sorting

// [medium] HeapSort returns xs sorted in ascending order.
//
// Build a max-heap, then repeatedly extract the maximum.
// Steps:
// 1. Copy xs so the caller's slice is not mutated.
// 2. Build a max-heap in place: for i from n/2-1 down to 0, sift down node i.
// 3. Repeatedly swap the root (maximum) with the last unsorted item.
// 4. Shrink the heap by one and sift the new root down; repeat.
//
// Notes:
//   - Children of node i live at 2i+1 and 2i+2; no pointers needed.
//   - O(n log n) on every input shape, O(1) auxiliary space.
//   - Not stable.
//
// Expected: best/average/worst O(n log n)
// Input: not mutated
func HeapSort(xs []int) []int {
	// TODO: implement
	return nil
}
