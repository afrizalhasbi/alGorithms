package numerics

import "math"

// [easy] SafeLog and SafeExp are the domain-guarded wrappers around the
// standard logarithm and exponential.
//
// Background. Two failure modes come up constantly when composing numeric
// code. The logarithm of a non-positive number is not a number: the
// standard library reports it as NaN, and NaN is a poison pill — it
// silently propagates through every sum, product, and comparison that
// touches it, and the final answer is NaN with no clue where it entered.
// The exponential of a large argument overflows to positive infinity,
// which is quieter but just as fatal downstream.
//
// The guarded wrappers convert these silent failures into explicit ones.
// SafeLog refuses non-positive arguments by reporting failure instead of
// returning NaN. SafeExp refuses arguments beyond the overflow threshold
// (about seven hundred and nine, the log of the largest finite float) by
// reporting failure instead of returning infinity. Callers one level up
// can then make a real decision — clamp, reparameterize, or give up —
// instead of debugging a NaN five functions later.
//
// Steps:
//  1. SafeLog: if the argument is positive, return its logarithm and
//     success; otherwise return failure. NaN input is failure.
//  2. SafeExp: if the argument is within the representable range, return
//     its exponential and success; otherwise return failure. (Negative
//     arguments never fail; they underflow to a legal, harmless zero.)
//
// Convention (matches the tests): the overflow threshold is the argument
// beyond which the standard exponential returns infinity; a result of
// positive infinity is a failure, not a value.
//
// Expected: O(1)
func SafeLog(x float64) (float64, bool) {
	if x <= 0 || math.IsNaN(x) {
		return 0, false
	}
	return math.Log(x), true
}

// SafeExp returns the exponential of x, reporting failure when the result
// would overflow. See SafeLog for the rationale.
func SafeExp(x float64) (float64, bool) {
	if x > 709 {
		return 0, false
	}
	return math.Exp(x), true
}
