package sorting

// [hard] RadixSort returns xs (non-negative ints) sorted in ascending order using
// LSD radix sort in base 16.
//
// Sort digit by digit from least significant to most significant, using a
// stable counting sort per digit.
// Steps:
//  1. Copy xs so the caller's slice is not mutated.
//  2. Find the maximum value; the number of hex digits k follows from it.
//  3. For each digit position d = 0..k-1:
//     a. Count how many items have each hex digit value (16 buckets).
//     b. Turn counts into prefix sums (bucket boundaries).
//     c. Scatter items into buckets in order; this must preserve current order.
//  4. After the most significant digit, the slice is sorted.
//
// Notes:
//   - Stability of each pass is not an optimization detail, it is required
//     for correctness.
//   - Runtime is O(n*k) where k = digits of the max value. With k constant
//     the empirical exponent is ~1, same shape as O(n).
//   - Non-negative ints only; values are in [0, n) from the generators.
//
// Expected: O(n*k), empirically ~O(n) for fixed k
// Input: not mutated
func RadixSort(xs []int) []int {
	// TODO: implement
	return nil
}
