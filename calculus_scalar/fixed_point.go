package calculus

// [easy] FixedPoint seeks a fixed point of g — a value x with g(x) = x —
// by iterating x from the starting point x0 until it stops moving.
//
// Background. Root finding and fixed points are the same problem in
// different clothes: a root of f is a fixed point of x minus f(x), and
// conversely. What the fixed-point form isolates is the single idea
// underlying every method in this package: replace the equation by a
// rearrangement whose right-hand side can be iterated, and hope the
// iteration contracts. Newton, gradient descent, and EM are all fixed
// point iterations under the hood — each step is "apply a map to the
// current guess".
//
// Whether the iteration converges is governed by the map's slope at the
// fixed point. If the magnitude of the local slope is below one, each
// pass shrinks the distance to the fixed point by that factor: a
// contraction. The linear convergence rate is exactly that factor, so
// the fixture checks the ratio between successive errors, not just the
// destination — an implementation that stumbles to the right answer at
// the wrong rate has not understood the method.
//
// The failure case is the mirror image: a map whose slope exceeds one in
// magnitude at the starting point repels iterates, and the honest
// response is to hit the iteration budget and report failure, not to
// mask the divergence with the budget.
//
// Definitions:
//  - fixed point: a value x with g(x) = x.
//  - iterate: the current guess. Each step applies g to it once.
//  - local slope: g's derivative at the fixed point. Its magnitude
//     below one means contraction (each step scales the distance to
//     the fixed point by that factor); above one means divergence.
//  - convergence test: the absolute difference of two successive
//     iterates within the tolerance.
//  - budget exhaustion: the iteration cap reached without meeting the
//     tolerance. Reported as failure — including on purpose, for the
//     deliberately expanding map in the fixture.
//
// Steps:
//  1. Iterate: the new value is g of the current value.
//  2. Stop when the absolute difference between successive values is
//     within the tolerance; stop with failure when the budget is
//     exhausted.
//
// Convention (matches the tests): the fixture stores contracting maps
// with known fixed points (computed in closed form) and their expected
// contraction rates, plus a deliberately expanding map that must fail by
// budget exhaustion.
//
// Expected: linear convergence at the map's local slope, one evaluation
// per iteration
func FixedPoint(g func(float64) float64, x0, tol float64, maxIter int) (float64, bool) {
	// TODO: implement
	return 0, false
}
