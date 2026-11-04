package sampling

// [medium] SampleExponential samples from the exponential distribution
// with the given rate, by inverting its cumulative distribution function
// in closed form.
//
// Background. The exponential distribution is the continuous case where
// the master recipe can be carried all the way: its CDF is one minus the
// exponential of minus rate times x, and that function is analytically
// invertible. Solving u equals the CDF for x gives: minus the logarithm
// of one minus u, divided by the rate. That closed form — one logarithm
// per sample — is the entire sampler.
//
// Two details are the exercise. First the one-minus-u: it is not
// decoration. The logarithm's argument must avoid zero, because the
// logarithm of zero is negative infinity; with u uniform on the
// half-open zero-one interval, u never reaches one, but one minus u can
// still round to zero when u is within one ulp of one. Some libraries
// use minus the logarithm of u instead, which is legal — u and one minus
// u have the same distribution — but gives different samples for the
// same uniform, so the convention must be stated and fixed; the tests
// pin it. Second, the domain guard: a negative uniform (bad input) or
// a non-positive rate must be refused rather than produce a NaN or a
// sign-flipped sample.
//
// Definitions:
//  - CDF: the cumulative distribution function of the target — the
//     probability of a draw at most x. Inverting it turns uniforms
//     into samples.
//  - rate: the exponential's parameter. Must be strictly positive.
//  - the uniform convention: u in [0, 1), and the inverse uses one
//     minus u (not u) in the logarithm. Legal alternates exist that
//     give the same distribution but different samples; the fixture
//     pins this one.
//  - branch split: where a piecewise inverse changes formula — at the
//     CDF's value at the peak (one half here). A draw exactly on the
//     split belongs to the lower branch.
//
// Steps:
//  1. Validate: the rate is positive; every u lies in [0, 1).
//  2. For each u: return minus the natural logarithm of (one minus u),
//     divided by the rate.
//
// Convention (matches the tests): the fixture stores uniforms and the
// expected samples computed with the convention "minus log of one
// minus u", to full float64 precision; agreement is exact up to a few
// ulps.
//
// Expected: O(draws)
func SampleExponential(rate float64, us []float64) []float64 {
	// TODO: implement
	return nil
}

// SampleTriangular samples from a triangular distribution on the
// interval from zero to two, with its peak at one, by inverting the CDF
// piecewise. The CDF is a quadratic on each side of the peak, so the
// inverse is a square root on each side: below the median, the square
// root of two u; above it, two minus the square root of two times
// (one minus u). This is the shape of the exercise: not every inverse
// needs to exist in closed form, but when a piecewise one does, the
// branches split at the CDF's value at the peak — one half here — and
// the fixture's uniforms straddle that split exactly.
//
// Convention (matches the tests): same uniform convention as
// SampleExponential; the u equal to exactly one half belongs to the
// lower branch, where both branches agree.
//
// Expected: O(draws)
func SampleTriangular(us []float64) []float64 {
	// TODO: implement
	return nil
}
