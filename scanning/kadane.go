package scanning

// [medium] Kadane returns the maximum sum of any contiguous subarray of xs.
//
// One pass, tracking the best subarray ending at the current position.
// Steps:
//  1. Fix the convention first, because the tests depend on it: the
//     empty subarray counts as a valid answer and sums to zero. So the
//     result is never below zero, and an empty or all-negative input
//     gives zero.
//  2. Sweep the slice once, carrying two running values: the best sum
//     of any subarray that ends exactly at the previous position, and
//     the best answer seen so far. Both start at zero.
//  3. At each item, decide: does the best subarray ending here extend
//     the one from before, or does it start fresh at this item? Extend
//     unless the running sum before this item is negative — a negative
//     prefix only drags the item down, so dropping it is always safe.
//     Then fold the result into the best-so-far.
//  4. The sweep runs to the end of the slice with no early exit; the
//     best-so-far is then the answer.
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
