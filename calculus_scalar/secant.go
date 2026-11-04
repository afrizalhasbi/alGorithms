package calculus

// [medium] SecantMethod returns a root of f from two starting points x0
// and x1, without a derivative: the slope of the secant line through the
// two most recent iterates stands in for it.
//
// Background. Newton's method needs the derivative at every step; the
// secant method gets a similar effect from function values alone. The
// line through the two last evaluated points crosses zero somewhere
// between or near them; that crossing is the next iterate, the older
// point is dropped, and the process repeats. The slope used is thus a
// finite-difference approximation of the derivative — but one assembled
// for free from points the iteration visits anyway.
//
// The convergence story is the interesting part. Newton doubles the
// correct digits per step; the secant method manages about one point six
// — slower per step, but each step costs only one new function
// evaluation instead of two, so per unit of work the two are close, and
// the secant method wins whenever the derivative is expensive or
// unavailable. The same failure modes as Newton apply — horizontal
// secant when the two iterates return the same value, divergence,
// cycling — and the same bounded budget with explicit failure applies.
// The tests hold it to the iteration counts, which sit strictly between
// bisection's logarithmic rate and Newton's quadratic rate; matching
// Newton's count or bisection's count is a bug, not a rounding difference.
//
// Definitions:
//  - iterates: the two most recent points, the window. The older is
//     dropped each step; the new iterate becomes the newer point.
//  - secant line: the line through the two iterates. Its zero crossing
//     is the next iterate; its slope is the stand-in for the
//     derivative, assembled from points the iteration visits anyway.
//  - flat secant: the two most recent iterates have equal function
//     values — the slope is zero. Reported as failure.
//  - slide the window: the update in step 3. Pure re-labeling; no
//     values are recomputed for it.
//  - stopping criteria: the step size or the function value within
//     tolerance; failure on a flat secant or exhausted budget.
//
// Steps:
//  1. Validate: the two starting points are distinct and give distinct
//     function values (otherwise the secant is vertical or the slope is
//     zero).
//  2. Step: the new iterate is the newer point minus its function value
//     times the gap between the points, divided by the difference of the
//     two function values.
//  3. Slide the window: the old newer point becomes the old point, the
//     new iterate becomes the newer point.
//  4. Stop when the step or the function value is within tolerance; stop
//     with failure on a flat secant or an exhausted budget.
//
// Convention (matches the tests): the fixture stores the expected root
// and the expected iteration count; the tolerance and budget are fixed
// by the test so the counts are comparable across runs.
//
// Expected: sub-quadratic, better than linear convergence, at one
// evaluation per iteration
func SecantMethod(f func(float64) float64, x0, x1, tol float64, maxIter int) (float64, bool) {
	// TODO: implement
	return 0, false
}
