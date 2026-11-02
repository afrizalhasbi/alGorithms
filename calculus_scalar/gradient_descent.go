package calculus

// [easy] GradientDescent minimizes a function of one variable whose
// derivative df is supplied, starting at x0 with a fixed step size.
//
// Background. The derivative is the local slope; the function decreases
// in the direction the slope points up, so the steepest downhill step
// from any point is in the direction opposite the derivative. Gradient
// descent applies exactly that: subtract a multiple of the derivative
// from the current point, forever. It is the ancestor of every
// first-order method — nothing in modern large-scale optimization uses
// a different core loop, only better step sizes and extra momentum.
//
// The step size is the whole art. Too large and the iterate overshoots
// the valley and oscillates, or diverges outright; too small and the
// iteration crawls. For a quadratic function the stability limit is
// exact: the fixed step must be smaller than two over the curvature,
// or the iteration bounces outward instead of settling. The fixture
// includes a step just over that limit and requires the implementation
// to fail by budget exhaustion rather than return whatever it landed on.
//
// The convergence is linear — the error shrinks by a constant factor per
// step, set by how close the step size is to the stability limit — and
// the fixture checks the rate as well as the destination.
//
// Steps:
//  1. Iterate: the new point is the current point minus the step size
//     times the derivative at the current point.
//  2. Stop when the absolute change in the point is within the
//     tolerance; stop with failure when the budget is exhausted.
//
// Convention (matches the tests): derivatives are hardcoded in the test;
// the fixture stores minimizers computed in closed form and the expected
// per-step shrink factors.
//
// Expected: linear convergence, two evaluations per iteration (or one,
// if the derivative is reused from the same call)
func GradientDescent(df func(float64) float64, x0, step, tol float64, maxIter int) (float64, bool) {
	// TODO: implement
	return 0, false
}
