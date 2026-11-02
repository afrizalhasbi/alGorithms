package calculus

// [easy] NewtonRoot returns a root of f near the starting point x0, given
// the derivative df, by Newton's method. Both f and df are supplied by
// the caller; nothing is differentiated automatically.
//
// Background. Where the root actually is, locally: follow the tangent.
// The tangent line at the current point is the first-order model of f,
// it crosses zero at the current point minus the function value over the
// slope, and that crossing point is the next iterate. When it converges,
// it doubles the number of correct digits per step — a root found to
// twelve digits in six evaluations is typical — which is the fastest
// behavior of any method in this package.
//
// The price is the failure modes, and the exercise is to notice them
// rather than paper over them. A slope of zero at the iterate means the
// tangent is horizontal and the step is undefined; report failure. A
// function with several roots near the start can converge to a different
// one than expected, or leave the region of interest entirely. And the
// iteration can cycle or diverge — the classic case is a function whose
// tangent throws the iterate far away, from which it never returns — so
// a bounded number of iterations is a correctness requirement, not an
// optimization: exceeding the budget without meeting the tolerance is
// reported as failure, exactly like the slope problem.
//
// Steps:
//  1. Validate: the slope at the starting point is not zero (or
//     denormal-close to it).
//  2. Step: subtract the function value over the slope from the current
//     point.
//  3. Stop when the absolute step, or the function value, is within the
//     tolerance; stop with failure if the slope is zero or the iteration
//     budget is exhausted first.
//
// Convention (matches the tests): the fixture stores the function's
// identity, the root computed to forty digits by an independent library,
// and the expected iteration count; the iteration counts are part of the
// contract, since they distinguish Newton from its slower cousins.
//
// Expected: O(1) evaluations for a simple root — the budget exists for
// the non-converging cases
func NewtonRoot(f, df func(float64) float64, x0, tol float64, maxIter int) (float64, bool) {
	// TODO: implement
	return 0, false
}
