package calculus

// [easy] CentralDifference approximates the derivative of f at x with the
// symmetric difference quotient, and the test verifies its error obeys
// the second-order law.
//
// Background. The forward difference — the function value a step h ahead
// minus the value here, over h — is the obvious approximation, but its
// error is proportional to h: halving the step halves the error. The
// central difference averages the slopes to the points h ahead and h
// behind, and by symmetry the first-order error term cancels: the error
// becomes proportional to h squared. Halving the step now quarters the
// error, a law the test verifies exactly across a ladder of step sizes.
//
// The other half of the lesson is where the law breaks. Shrink h too far
// and the difference f(x+h) minus f(x-h) subtracts two nearly equal
// numbers: catastrophic cancellation eats the gains faster than the
// theory predicts, and the error turns around and grows as h shrinks
// past the sweet spot (near the square root of the machine epsilon). The
// fixture's ladder deliberately walks into that regime; the test asserts
// the error decreases down the ladder until the cancellation floor, and
// turns around below it.
//
// This routine is a utility for the tests of every other algorithm in
// the package — a way to check a hardcoded derivative — and an exercise
// in its own right about what finite differences can and cannot promise.
//
// Steps:
//  1. Evaluate f at x plus h and x minus h.
//  2. Return their difference divided by two h.
//
// Convention (matches the tests): the reference derivatives are known in
// closed form; the fixture stores the error at each rung of the step
// ladder.
//
// Expected: O(1)
func CentralDifference(f func(float64) float64, x, h float64) float64 {
	// TODO: implement
	return 0
}
