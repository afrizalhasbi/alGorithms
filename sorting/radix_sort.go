package sorting

// [hard] RadixSort returns xs (non-negative ints) sorted in ascending order using
// LSD radix sort in base 16.
//
// Background. Radix sort does not compare items with each other at
// all. Instead it sorts by the digits of the items, one digit
// position at a time, starting from the least significant digit
// (that is what LSD stands for) and ending at the most significant.
// In base 16 each digit is a value from 0 to 15, so one hex digit of
// an integer is just four of its bits.
//
// Each pass sorts the whole slice by one digit. The pass itself is a
// counting sort, which works because a digit has only a small number
// of possible values: count how many items carry each digit, use the
// counts to work out where each digit group starts in the output,
// then move every item into its group's region of a scratch copy.
//
// A sort is stable when items that compare equal keep the order they
// had before the sort. For radix sort stability is not a nicety, it
// is the correctness argument. After the pass on the last digit, all
// items with equal last digits are grouped in the order they came
// in. The pass on the next digit regroups everything by that digit,
// and stability is what preserves the earlier grouping among items
// that share the new digit. If any pass is unstable, the earlier
// work is thrown away and the result is wrong.
//
// Definitions:
//  - digit: the base-16 digit of an item at the current position.
//     Digit 0 is the least significant (the lowest four bits); each
//     digit is four bits of the integer.
//  - pass: one counting sort over one digit position, from digit 0
//     upward. The number of passes is set by the digit count of the
//     largest value.
//  - counts: the 16 counters, one per possible digit value, filled in
//     step 3a.
//  - starting positions: the prefix sums of the counts. Counts "how
//     many items carry this digit"; starting positions answer "where
//     in the output this digit's region begins".
//  - scratch copy: the auxiliary buffer each pass writes into, before
//     being copied back. Each pass reads the slice and writes the
//     scratch; the roles are NOT swapped between passes — the copy
//     back in step 3d keeps the slice the source of truth.
//
// Steps:
//  1. A slice with one or zero items is already sorted.
//  2. Find the largest value. It decides how many passes are needed:
//     keep passing while some item still has an unprocessed digit,
//     starting at the least significant one.
//  3. One pass, for the current digit:
//     a. Count the items carrying each of the 16 digit values.
//     b. Convert counts to starting positions, so each digit value
//     has a known region of the output.
//     c. Walk the slice from front to back, moving each item into the
//     next free slot of its digit's region in the scratch copy.
//     Walking in order is what makes the pass stable.
//     d. Copy the scratch back over the slice.
//  4. After the pass on the most significant digit, the slice is
//     sorted.
//
// Notes:
//   - Runtime is O(n*k) where k is the number of digits of the
//     largest value. With k constant the empirical exponent is ~1,
//     same shape as O(n).
//   - Non-negative ints only; values are in [0, n) from the
//     generators.
//
// Expected: O(n*k), empirically ~O(n) for fixed k
// Input: mutated in place
func RadixSort(xs []int) []int {
	// TODO: implement
	return nil
}
