package sampling

import "math/rand"

// [medium] RejectionSample draws n samples from a one-dimensional target
// density, given a simpler envelope density and the constant m with
// target(x) at most m times envelope(x) everywhere, by accept-reject.
//
// Background. The only fully general recipe for exact sampling from an
// arbitrary density. Draw a candidate from the envelope distribution and
// an independent uniform on the zero-one interval; accept the candidate
// when the uniform is below the ratio of the target density to m times
// the envelope density at the candidate, and reject otherwise; repeat
// until enough samples are banked. The correctness argument is that the
// accepted points are the target density reweighted by the envelope —
// the acceptance probability at each x is exactly proportional to the
// target there, so thinning the envelope's samples reproduces the
// target's distribution exactly. Nothing is approximate: this is exact
// sampling, not an asymptotic method.
//
// The constant m is the price of the method: the expected number of
// candidates per accepted sample is exactly m, so a loose envelope (large
// m) wastes proportionally more draws. The fixture's target is a
// parabola density on the unit interval with its maximum of one and a
// half, under a uniform envelope with m equal to one and a half, so the
// theoretical acceptance rate is two thirds — and the test asserts the
// observed acceptance count sits in a band around that value, because a
// wrong m or a reversed ratio shows up immediately as a wildly wrong
// rate.
//
// Definitions:
//  - target / envelope: the density wanted, and the simpler density
//     candidates are drawn from. Both non-negative; the envelope is
//     sampled by a caller-supplied sampler.
//  - m: the envelope constant with target(x) at most m times
//     envelope(x) everywhere. Determines the expected candidates per
//     accepted sample, exactly m.
//  - acceptance test: a uniform draw, multiplied by m times the
//     envelope density at the candidate, below the target density
//     there. Reversing the ratio or mis-sizing m shows up as a wildly
//     wrong acceptance rate.
//  - banked samples: the accepted candidates. Collected until there
//     are n; rejected candidates are consumed and gone.
//  - candidates: the total number of envelope draws consumed,
//     returned so the test can compare the observed acceptance rate
//     with the theoretical one.
//
// Steps:
//  1. Validate: m is at least the sup of the ratio of target to envelope
//     densities (the caller's responsibility; the tests use m from the
//     fixture), and the densities are non-negative.
//  2. Repeat until n samples are banked: draw a candidate from the
//     envelope sampler (supplied by the caller, so the test controls it
//     through the seeded generator), draw a uniform; accept the candidate
//     when the uniform, multiplied by m times the envelope density at
//     the candidate, is below the target density there.
//  3. Return the banked samples, and the total number of candidates
//     consumed, so the test can check the acceptance rate.
//
// Convention (matches the tests): the fixture stores the target's
// parameters, m, and the theoretical acceptance rate; the observed rate
// must sit within a few standard errors of the theoretical value, and
// the accepted samples' moments must match the target's moments.
//
// Expected: O(m times n) candidates for n samples
func RejectionSample(target, envelope func(float64) float64, m float64, n int, rng *rand.Rand) (samples []float64, candidates int) {
	// TODO: implement
	return nil, 0
}
