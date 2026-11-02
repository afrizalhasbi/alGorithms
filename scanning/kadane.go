package scanning

// [medium] Kadane returns the maximum sum of any contiguous subarray of xs.
//
// Background. A subarray is a stretch of consecutive items of the
// slice. Contiguous means the items sit next to each other with no
// gaps; the answer can never skip an item in the middle. The task is
// to find the stretch with the largest total. There is an obvious
// O(n^2) way: try every starting point and extend it to every ending
// point. Kadane's algorithm gets the same answer in one pass by
// never looking back.
//
// The idea: for each position, what is the best subarray that ends
// exactly there? It either extends the best subarray ending at the
// previous position, or it starts fresh at this position, and the
// better of the two is the answer. Starting fresh wins exactly when
// the running sum carried from before is negative: a negative prefix
// only drags the current item down, so dropping it is always safe
// and never loses a better answer. That is the whole correctness
// argument.
//
// Convention (matches the tests): the empty subarray is allowed and
// sums to zero. The result is therefore never negative, and an empty
// or all-negative input gives zero.
//
// Steps:
//  1. Carry two values through one pass: the best subarray sum ending
//     at the previous position, and the best sum seen anywhere so
//     far. Both start at zero.
//  2. At each item: extend-or-restart as described above to get the
//     best sum ending here, then fold it into the best sum seen so
//     far.
//  3. When the pass ends, the best-so-far is the answer. There is no
//     early exit.
//
// Expected: O(n)
// Input: not mutated
func Kadane(xs []int) int {
	// best prev & best global
	// bp := 0
	// bg := 0
	// for i := range len(xs) {
	// is contiguous
	// }
	return 0
}
