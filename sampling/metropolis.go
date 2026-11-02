package sampling

import "math/rand"

// [hard] MetropolisHastings1D draws from a one-dimensional target density
// known only up to proportionality, by the Metropolis random walk with a
// symmetric uniform proposal.
//
// Background. The method that made sampling from unnormalized densities
// possible. A Markov chain wanders the real line: from the current state
// it proposes a step drawn uniformly from minus delta to delta; the step
// is accepted with probability the ratio of the target density at the
// proposed point to the target density at the current point — not the
// difference, the ratio, which is why the target may be known only up to
// a multiplicative constant. An accepted step moves the chain; a rejected
// step leaves it where it is, and the rejected-step stationary state is
// part of the chain, not a failure. When the proposal is symmetric, this
// accept rule satisfies detailed balance against the target, so the
// target is the chain's stationary distribution: run it long enough and
// the states it visits are samples from the target.
//
// Three practical facts the tests rely on, and which are the actual
// lessons. First, acceptance behavior is diagnostic: a wrong accept
// ratio (comparing densities instead of their ratio, or inverting the
// ratio) collapses the observed acceptance rate toward zero or one; the
// test asserts a mid-range band. Second, consecutive samples are
// correlated — this is the method's defining weakness against rejection
// sampling — so only every k-th state is treated as a sample, and the
// chain's early states (before it forgets the starting point) are
// discarded as burn-in; the test drops the first half. Third, the
// uniformity of the accept probability must use a draw even when
// rejecting, so the random stream's consumption is deterministic given
// the seed — otherwise reproducing a run twice becomes impossible.
//
// Steps:
//  1. Validate: delta positive, at least one iteration, initial state
//     finite.
//  2. For each iteration: draw a uniform step from minus delta to delta
//     and add it to the current state; draw a uniform acceptance
//     threshold on the zero-one interval; accept the proposed state when
//     the threshold is below the ratio of the target at the proposal to
//     the target at the current state (and always accept when the
//     current state has zero density); otherwise stay.
//  3. Record every state, accepted or not, in order.
//
// Convention (matches the tests): the fixture's target is the standard
// normal (up to its normalizing constant, which the implementation must
// not require); the test fixes the seed, drops the first half of the
// chain as burn-in, and checks the remaining states' mean, variance, and
// an acceptance-rate band. Tolerances are stated in standard errors of
// the sample size.
//
// Expected: O(iters)
// Input: not mutated
func MetropolisHastings1D(target func(float64) float64, x0, delta float64, iters int, rng *rand.Rand) []float64 {
	// TODO: implement
	return nil
}
