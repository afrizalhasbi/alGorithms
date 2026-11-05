package numerics

import "math"

// [easy] NeumaierSum returns the sum of xs using Neumaier's improvement of
// Kahan's algorithm.
//
// Background. Kahan summation has a blind spot: it only recovers the lost
// bits when the next item is smaller in magnitude than the running total.
// When an item is larger, the compensation itself gets rounded away.
// Neumaier's fix is to track the compensation as a full second-order sum:
// whenever an item is at least as large as the running total in magnitude,
// the roles swap — the item becomes the new total and the old total, which
// is now the small one, is what gets folded into the compensation. The
// compensation then accumulates the errors of all the additions that were
// "small relative to the total", regardless of the order they arrived in.
//
// The practical difference shows up on inputs whose magnitudes are not
// sorted, or that grow as you go: on the descending-magnitude fixture the
// plain Kahan loop can still lose most of the small terms, while Neumaier
// recovers them. Both agree with the exact rational answer up to a few
// ulps on friendly inputs; the fixtures in sums_desc.json are where the
// two implementations part ways.
//
// Definitions:
//   - ordinary sum: the float64 value that evaluating item + total returns.
//     As everywhere in this package, the exact mathematical sum exists
//     separately; the + operator rounds, so the ordinary sum is its rounded
//     shadow. There is no rounding function.
//   - loss: the bits that rounding threw away when forming the ordinary
//     sum — the difference between the exact mathematical sum of item and
//     total and the ordinary sum, in whichever of the two directions the
//     step below describes.
//   - compensation accumulator: unlike Kahan's, this one accumulates: each
//     iteration's measured loss is added into it, and it is folded into
//     the total only once, at the end of the pass.
//
// Steps:
//  1. Carry two values through the pass: the running total and the
//     compensation accumulator (initially zero).
//  2. For each item: form the ordinary sum of the item and the total. If
//     the item's magnitude is at least the total's magnitude, the loss is
//     measured as: item minus that ordinary sum, plus the total.
//     Otherwise it is: total minus that ordinary sum, plus the item.
//     Add the measured loss into the compensation accumulator, then set
//     the total to the ordinary sum.
//  3. When the pass ends, add the compensation accumulator into the total
//     and return that.
//
// Convention (matches the tests): the empty list sums to zero.
//
// Expected: O(n)
// Input: not mutated
func NeumaierSum(xs []float64) float64 {
	var tot = 0.
	var comp = 0.
	for _, x := range xs {
		orsum := tot + x
		var loss = 0.
		if math.Abs(x) >= math.Abs(tot) {
			loss = x - orsum + tot
		} else {
			loss = tot - orsum + x
		}
		comp += loss
		tot = orsum
	}
	return tot + comp
}
