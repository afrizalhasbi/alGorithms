package statistics

// [easy] BayesPosterior returns the posterior distribution over a set of
// hypotheses after one discrete observation.
//
// Background. Bayes' rule with finite hypotheses. The input is a prior —
// one probability per hypothesis — and a likelihood table: for each
// hypothesis, the probability of each possible observation. Seeing
// observation o, every hypothesis h gets weight prior(h) times
// likelihood(h, o); those weights are then normalized so they sum to one
// again. The normalization constant is the marginal probability of the
// observation — the denominator of Bayes' rule — and it falls out of the
// sum; it never needs to be known in advance.
//
// The exercise is small but the failure mode is classic: if some
// hypothesis assigned zero probability to the observed outcome, the
// unnormalized weights can all be zero for that observation, and
// normalization becomes division by zero. A single observation that every
// hypothesis considers impossible makes the posterior undefined; that is
// a reported error, not a crash.
//
// Steps:
//  1. Validate: the prior sums to one; every row of the likelihood table
//     sums to one; the observation index is in range.
//  2. Compute the unnormalized posterior weight of each hypothesis: prior
//     times the likelihood of the observation under that hypothesis.
//  3. Sum the weights. If the sum is zero, the observation is impossible
//     under every hypothesis; report failure.
//  4. Divide every weight by the sum and return the normalized list.
//
// Convention (matches the tests): the returned posterior is a new slice,
// ordered like the prior; the prior and likelihood are not mutated. The
// fixtures include a prior that makes one hypothesis a near-certainty and
// one observation that is impossible under one hypothesis, so the
// normalization must cope with zeros in the weights.
//
// Expected: O(hypotheses times outcomes)
// Input: not mutated
func BayesPosterior(prior []float64, likelihood [][]float64, obs int) ([]float64, bool) {
	// TODO: implement
	return nil, false
}
