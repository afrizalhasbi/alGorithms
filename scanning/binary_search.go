package scanning

// [easy] BinarySearch returns the index of target in the sorted slice xs, or -1.
//
// Background. The input is sorted, and that single fact allows the
// search to discard half of the slice on every comparison. Keep a
// range of positions, the interval, where the target could still be.
// At the start the interval is the whole slice. Compare the target
// with the item at the interval's midpoint. If they are equal, the
// search is done. If the target is larger, the midpoint and
// everything before it cannot hold the target (the slice is sorted),
// so the entire front half of the interval is discarded. If the
// target is smaller, the back half is discarded instead.
//
// Every comparison halves the interval, which is why the number of
// comparisons grows like the logarithm of the length: doubling the
// slice costs one more comparison.
//
// The invariant to keep in mind: if the target is present at all,
// it is inside the current interval. The search ends when the
// interval is empty, and by the invariant that means the target is
// not present.
//
// Steps:
//  1. Start with the interval covering the whole slice.
//  2. Each round, compare the target with the midpoint item. Equal:
//     return that position. Target larger: keep the back half of the
//     interval. Target smaller: keep the front half. The midpoint
//     itself is excluded either way, so the interval always shrinks
//     and the loop cannot stall.
//  3. When the interval becomes empty, report "not found" (-1).
//     Empty means front passed back; stopping any earlier would risk
//     missing a target sitting on the last surviving item.
//
// Notes:
//   - Use lo + (hi-lo)/2, not (lo+hi)/2, to avoid integer overflow.
//   - Complexity tests for O(log n) show an exponent near 0 — the
//     time barely grows. The test amplifies the signal by doing many
//     searches per call.
//   - With duplicate values, any position holding the target is an
//     acceptable answer; the tests only check that the reported
//     position really holds it.
//
// Expected: O(log n) — empirically exponent ~0
// Input: not mutated
func BinarySearch(xs []int, target int) int {
	// TODO: implement
	return -1
}
