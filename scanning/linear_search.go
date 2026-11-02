package scanning

// [easy] LinearSearch returns the index of target in xs, or -1 if absent.
//
// Scan left to right and compare every item. This is the fully implemented
// reference for the exercise pattern: read the doc, make the complexity test
// pass, keep the implementation minimal.
//
// Steps:
// 1. For i from 0 to len(xs)-1, compare xs[i] with target.
// 2. On the first match, return i immediately.
// 3. If the scan finishes without a match, return -1.
//
// Notes:
//   - The worst case is target absent (or last): it forces a full scan, and
//     that is exactly what the complexity test measures.
//
// Expected: O(n)
// Input: not mutated
func LinearSearch(xs []int, target int) int {
	for i, x := range xs {
		if x == target {
			return i
		}
	}
	return -1
}
