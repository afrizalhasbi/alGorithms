// Package calculus_scalar holds the scalar calculus exercises: root
// finding, minimization, and iteration families that need at most a
// derivative (supplied by the caller, never computed automatically).
package calculus

import "math"

// [easy] Bisection returns a root of f on the bracket [a, b], to within
// tol, or reports failure.
//
// Background. The one algorithm that always works, if it works at all.
// If f is continuous and f(a) and f(b) have opposite signs, some root
// lies between them. Evaluate the midpoint: the half on which the sign
// change survives still brackets a root, and its sign pattern is the
// same guarantee, half the width. Every iteration halves the bracket, so
// the width after m steps is the initial width times one half to the m —
// the number of steps needed is therefore known in advance from the
// tolerance alone, before a single evaluation. That is also its
// weakness: the bracket shrinks, but nothing about the shape of f is
// ever exploited, so it walks to the answer at a fixed one bit per
// iteration.
//
// The failure mode to handle honestly is the degenerate bracket: if f(a)
// and f(b) do not bracket a sign change, bisection has no contract, and
// must report that rather than converge to whatever point it happens to
// visit. The tolerance has a floor too: asking for a bracket narrower
// than one ulp of the midpoint is asking for the impossible, so the
// iteration also stops when the midpoint stops moving, and the test
// counts the iterations to confirm the halving law held.
//
// Definitions:
//   - bracket: the interval [a, b] whose endpoints have f(a) and f(b) of
//     opposite signs. The bracket, not f, is the state the iteration
//     carries.
//   - sign change survives: the half of the bracket on which the
//     endpoints still have opposite signs (or hit zero exactly).
//   - midpoint: (a+b)/2 computed in floating point. The iteration ends
//     when the midpoint equals an endpoint in float — that is the
//     ulp-level floor, not the tolerance.
//   - iteration count: the number of halvings. It is predicted in
//     advance by the tolerance and verified by the test.
//
// Steps:
//  1. Validate: a is less than b, and f(a) times f(b) is negative (or
//     one of them is zero — an endpoint is already a root).
//  2. Compute the midpoint. If f(midpoint) is zero, or the bracket is
//     already within the tolerance, or the midpoint equals an endpoint in
//     floating point, stop and return the midpoint.
//  3. Otherwise keep the half that still brackets the sign change: the
//     f(endpoint) whose sign matches f(midpoint) is replaced by the
//     midpoint. Repeat from step 2.
//
// Convention (matches the tests): the fixture stores the bracket, the
// root computed to forty digits with an independent library, and the
// number of iterations the halving law predicts; the returned point must
// be within the tolerance of the true root, and the iteration count must
// match the prediction.
//
// Expected: O(log of the initial width over the tolerance) evaluations
func Bisection(f func(float64) float64, a, b, tol float64) (float64, bool) {
	if a >= b || math.Signbit(f(a)) == math.Signbit(f(b)) {
		return 0, false
	}
	if f(a) == 0 {
		return a, true
	}
	if f(b) == 0 {
		return b, true
	}

	var lo = a
	var hi = b

	for true {
		mid := (lo + hi) / 2.
		midf := f(mid)

		if midf == 0 || math.Abs(hi-lo) < tol || mid == hi || mid == lo {
			return mid, true
		}

		lof := f(lo)
		hif := f(hi)

		if math.Signbit(lof) == math.Signbit(midf) {
			lo = mid
		}
		if math.Signbit(hif) == math.Signbit(midf) {
			hi = mid
		}

	}
	return 0, false
}
