package sorting

// [medium] HeapSort returns xs sorted in ascending order.
//
// Build a max-heap, then repeatedly extract the maximum.
// Steps:
//  1. First understand the repair operation, since both later phases
//     depend on it. A section of the slice is treated as a heap: each
//     item has two children beneath it. To repair a node, compare it
//     with its children; if the node is smaller than its largest child,
//     swap the two and repair the child's position in turn. Stop when
//     the node has no children, or when it is already at least as large
//     as both of them.
//  2. Build the heap: repair every parent node, starting from the last
//     parent in the slice and working backwards to the root. Working
//     backwards matters: by the time a node is repaired, everything
//     below it is already a valid heap.
//  3. Extract repeatedly. The maximum is always at the root. Swap the
//     root with the last item of the heap, shrink the heap by that one
//     item (it is now in its final place, sorted), and repair the new
//     root. Repeat, shrinking one item at a time, until the heap is
//     empty and the whole slice is sorted.
//
// Notes:
//   - Children of node i live at 2i+1 and 2i+2; no pointers needed.
//   - O(n log n) on every input shape, O(1) auxiliary space.
//   - Not stable.
//
// Expected: best/average/worst O(n log n)
// Input: mutated in place
func HeapSort(xs []int) []int {
	// TODO: implement
	return nil
}
