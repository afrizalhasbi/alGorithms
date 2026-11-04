package numerics

// [medium] TaylorExp approximates the exponential function by its Taylor
// series, summing terms until they stop changing the answer.
//
// Background. The exponential is the friendliest possible Taylor series:
// every derivative at zero is one, so the coefficient of the k-th term is
// exactly one over k factorial, and the terms start at one, then x, then
// x squared over two, and shrink monotonically for positive x. Two things
// make the exercise more than "write a loop over factorials".
//
// First, computing factorials separately overflows quickly and is the
// wrong shape: each term should be derived from the previous one by
// multiplying by x and dividing by the term index — a term recurrence.
// The intermediate values then stay modest for moderate x.
//
// Second, and this is the numerics lesson of the exercise: for negative x
// the series strictly alternates in sign, and summing an alternating
// series whose terms shrink through rounding is where cancellation bites.
// The terms near the end of the sum are tiny compared to the running
// total, so naive left-to-right summation rounds them away, and the
// partial sums for negative x stop improving even as more terms are added.
// The fix is to sum from the smallest term to the largest: accumulate the
// terms into a slice and fold them in reverse, or recurse — the last term
// is built first and the additions always add something small to something
// at most as large as it, keeping every bit of every term.
//
// Definitions:
//  - term: one summand of the series. The first term is 1, the next is x,
//     the next is x squared over two, and so on. Terms are numbered
//     starting at 1: term number k equals the previous term times x,
//     divided by k.
//  - term index: the number of the NEW term being computed, not of the
//     previous one. Dividing by the previous index is the off-by-one.
//  - stop condition: adding the next term to the running result leaves
//     the result unchanged — the term is smaller than half an ulp of the
//     total. The term budget is a hard cap independent of this test.
//
// Steps:
//  1. Generate terms by recurrence: start from one; each next term is the
//     previous term times x, divided by its term index; stop when adding a
//     new term to the running result would not change it (the term is
//     smaller than half an ulp of the total) or when the term budget runs
//     out.
//  2. Sum the terms smallest-first as described above.
//
// Convention (matches the tests): the term budget bounds the loop for
// large or negative x; the reference is the standard library's Exp within
// a small relative tolerance. The series is only accurate for moderate x
// (say within about ten of zero); the tests stay in that range.
//
// Expected: O(maxTerms)
func TaylorExp(x float64, maxTerms int) float64 {
	// TODO: implement
	return 0
}
