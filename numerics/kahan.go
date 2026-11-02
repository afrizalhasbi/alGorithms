// Package numerics collects the small, dry, floating-point workhorses that
// every other numeric algorithm silently depends on: stable summation,
// stable evaluation, and honest comparison of floats.
package numerics

// [easy] KahanSum returns the sum of xs using compensated (Kahan) summation.
//
// Background. Adding floats left to right is not associative: each addition
// rounds, and once the running total is large, adding a tiny number to it
// does nothing — the small value is rounded away entirely. The classic
// demonstration is the list 1e16, 1, 1, ..., 1, -1e16: the exact answer is
// the number of ones, but naive summation returns zero, because every 1
// added to 1e16 disappears into rounding.
//
// The idea: keep a second value, the running compensation, which remembers
// exactly what was lost to rounding in the previous addition. Instead of
// adding x to the total directly, first recover the lost low-order bits of
// the total (the difference between the true sum of total and x and the
// rounded total), park them in the compensation, and add the compensation
// back in on the next round. The error is not eliminated, but it stops
// accumulating: naive summation's error grows in proportion to the number
// of terms, while Kahan's grows much more slowly, independently of n.
//
// Steps:
//  1. Carry three values through the pass: the running total, the
//     compensation from the previous step (initially zero), and a temporary.
//  2. For each item: compute the temporary as the item minus the
//     compensation. Predict what the total will become (total plus the
//     temporary), then compute the compensation as the difference between
//     the predicted value and the actual rounded total — in the correct
//     direction (predicted minus new total, added back to the temporary).
//     Finally add the compensation to the total.
//  3. When the pass ends, the total is the answer. There is no early exit.
//
// Convention (matches the tests): the empty list sums to zero. The exact
// reference sums in the test fixtures were computed with rational
// arithmetic on the exact binary values of the inputs, so they are the
// true sums, not another rounded estimate.
//
// Expected: O(n)
// Input: not mutated
func KahanSum(xs []float64) float64 {
	// TODO: implement
	return 0
}
