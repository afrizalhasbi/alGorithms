package numerics

// [easy] Horner evaluates a polynomial at a point using Horner's method.
//
// Background. A polynomial of degree n naively costs one exponentiation
// and one multiplication per coefficient — about n squared by one count,
// n by a better one — and evaluates a power of x from scratch every time.
// Horner's method rewrites the polynomial by repeatedly factoring x out
// of the remaining terms: the value of a degree-three polynomial at x is
// the coefficient of the highest term, times x, plus the next coefficient,
// times x, plus the next, times x, plus the constant. One pass, n
// multiplications, n additions, no exponentiation, and it is also the
// most accurate order to evaluate in: the additions group the way the
// polynomial's own algebra groups.
//
// The polynomial is given with coefficients in descending order: the first
// coefficient belongs to the highest power and the last is the constant.
// That is the same convention as the test fixture, which stores reference
// values produced by an independent polynomial evaluator.
//
// Definitions:
//  - running value: the single accumulator that carries the partially
//     evaluated polynomial through the pass. It starts as the leading
//     coefficient and ends as the result.
//  - remaining coefficients: everything after the first, visited in the
//     order they appear — highest power to constant. The loop does not
//     re-visit the leading coefficient; it is the seed of the running
//     value, not a loop iteration.
//
// Steps:
//  1. Start the running value at the leading coefficient.
//  2. For each remaining coefficient, in order: multiply the running value
//     by x, then add the coefficient.
//  3. The running value is the answer after the last coefficient.
//
// Convention (matches the tests): the empty coefficient list evaluates to
// zero. A one-element list is that constant for every x.
//
// Expected: O(n)
// Input: not mutated
func Horner(coeffs []float64, x float64) float64 {
	// TODO: implement
	return 0
}
