package scanning

// [easy] BinarySearch returns the index of target in the sorted slice xs, or -1.
//
// Halve the search interval each step by comparing against the midpoint.
// Steps:
//  1. Keep a search interval — the part of the slice where the target
//     could still live. Initially that is the whole slice. The
//     invariant: if the target is present at all, it is inside the
//     interval.
//  2. Each round, compare the target with the item at the interval's
//     midpoint. Equal: you found it, return that position. The target
//     is larger: the midpoint and everything before it are ruled out,
//     so keep only the back half. Smaller: keep only the front half.
//     Either way the midpoint itself is excluded, so the interval
//     strictly shrinks every round — the search cannot stall.
//  3. The loop ends when the interval becomes empty (front passes the
//     back). By the invariant the target cannot then be present:
//     return "not found". Empty the interval only via this condition —
//     stopping one step earlier misses a target sitting on the last
//     surviving item.
//  4. With duplicate values, any position holding the target is an
//     acceptable answer; the tests only check that the reported
//     position really holds it.
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
