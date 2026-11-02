// Package calculus_scalar holds the scalar calculus exercises: root
// finding, minimization, and iteration families that need at most a
// derivative (supplied by the caller, never computed automatically).
package calculus

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
// Steps:
//  1. Validate: a is less than b, and f(a) times f(b) is negative (or
//     one of them is zero — an endpoint is already a root).
//  2. Compute the midpoint. If f(midpoint) is zero, or the bracket is
//     already within the tolerance, or the midpoint equals an endpoint in
//     floating point, stop and return the midpoint.
//  3. Otherwise keep the half that still brackets the sign change: the
//     endpoint whose sign matches f(midpoint) is replaced by the
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
	// TODO: implement
	return 0, false
}
