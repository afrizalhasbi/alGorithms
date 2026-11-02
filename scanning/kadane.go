package scanning

// [medium] Kadane returns the maximum sum of any contiguous subarray of xs.
//
// One pass, tracking the best subarray ending at the current position.
// Steps:
//  1. Track two values: best (answer so far) and cur (best subarray that
//     ends exactly at index i).
//  2. For each item: cur = max(item, cur+item). Either extend the previous
//     subarray or start fresh here — whichever is larger.
//  3. best = max(best, cur) after each item.
//  4. Return best. Decide upfront what an all-negative slice should give
//     (empty subarray -> 0, or the largest single item) and be consistent.
//
// Notes:
//   - The core insight: a negative running prefix can never help, so dropping
//     it is always safe. That is the whole proof of correctness.
//
// Expected: O(n)
// Input: not mutated
func Kadane(xs []int) int {
	// TODO: implement
	return 0
}
