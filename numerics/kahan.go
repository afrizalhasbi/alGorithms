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
// Definitions:
//   - temporary: the item minus the compensation. This is what actually
//     gets added to the total this iteration.
//   - rounded sum: the float64 value that adding the total and the
//     temporary actually returns. Every float64 addition rounds its
//     exact result to the nearest representable value; there is no
//     separate rounding function.
//   - compensation: the part of the addition that the rounding threw
//     away, carried to the next iteration with the sign that makes
//     subtracting it from the next item restore the lost bits. It is a
//     single per-iteration quantity, not a running sum: each iteration's
//     value replaces the previous one.
//
// Steps:
//  1. Carry two values through the pass: the running total and the
//     compensation from the previous step (initially zero).
//  2. For each item: subtract the compensation from the item to form
//     the temporary, form the rounded sum of the total and the
//     temporary, then recover the rounding loss as the rounded sum
//     minus the total minus the temporary, evaluated in that order.
//     Store that loss as the new compensation, and set the total to the
//     rounded sum.
//  3. When the pass ends, the total is the answer. The final
//     compensation is not added again. There is no early exit.
//
// Convention (matches the tests): the empty list sums to zero. The exact
// reference sums in the test fixtures were computed with rational
// arithmetic on the exact binary values of the inputs, so they are the
// true sums, not another rounded estimate.
//
// Expected: O(n)
// Input: not mutated
func KahanSum(xs []float64) float64 {
	var tot = 0.
	var comp = 0.
	var tmp = 0.
	// had to have major hints/spelled it out for me
	// because the docs werent fucking clear when i
	// tried to implement this, fuck this
	for _, x := range xs {
		tmp = x - comp
		pred := tot + tmp
		comp = pred - tot - tmp
		tot += comp + tmp
	}
	return tot
}
