package sorting

// [hard] RadixSort returns xs (non-negative ints) sorted in ascending order using
// LSD radix sort in base 16.
//
// Sort digit by digit from least significant to most significant, using a
// stable counting sort per digit.
// Steps:
//  1. A slice with at most one item is already sorted; return it.
//  2. Look at the largest value to decide how many hex-digit passes
//     are needed: keep making passes while some item still has an
//     unprocessed digit left, starting from the least significant digit
//     and moving toward the most significant one.
//  3. Each pass sorts the whole slice by that single digit, using a
//     counting sort: first tally how many items carry each digit value;
//     then turn the tallies into starting positions, so you know where
//     the first item of each digit group belongs; then walk the slice
//     from front to back, dropping every item into its group's next
//     open slot in a scratch copy. Walking front to back is what keeps
//     items with equal digits in their current relative order — the
//     pass must be stable or the whole sort breaks. Finally copy the
//     scratch back over the original.
//  4. After the pass for the most significant digit, no passes remain
//     and the slice is sorted.
//
// Notes:
//   - Stability of each pass is not an optimization detail, it is required
//     for correctness.
//   - Runtime is O(n*k) where k = digits of the max value. With k constant
//     the empirical exponent is ~1, same shape as O(n).
//   - Non-negative ints only; values are in [0, n) from the generators.
//
// Expected: O(n*k), empirically ~O(n) for fixed k
// Input: mutated in place
func RadixSort(xs []int) []int {
	// TODO: implement
	return nil
}
