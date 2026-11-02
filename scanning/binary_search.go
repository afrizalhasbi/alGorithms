package scanning

// [easy] BinarySearch returns the index of target in the sorted slice xs, or -1.
//
// Halve the search interval each step by comparing against the midpoint.
// Steps:
//  1. Set lo = 0, hi = len(xs)-1.
//  2. While lo <= hi: let mid = lo + (hi-lo)/2.
//  3. If xs[mid] == target, return mid.
//  4. If xs[mid] < target, search the right half (lo = mid+1); else the left
//     half (hi = mid-1).
//  5. If the interval empties, the target is absent: return -1.
//
// Notes:
//   - Use lo + (hi-lo)/2, not (lo+hi)/2, to avoid integer overflow.
//   - Complexity tests for O(log n) show an exponent near 0 — the time barely
//     grows. The test amplifies the signal by doing many searches per call.
//
// Expected: O(log n) — empirically exponent ~0
// Input: not mutated
func BinarySearch(xs []int, target int) int {
	// TODO: implement
	return -1
}
