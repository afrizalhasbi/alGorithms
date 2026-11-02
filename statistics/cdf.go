package statistics

// [medium] CumulativeDistribution returns the running sums of a
// probability list, i.e. the cumulative distribution of a discrete
// variable with those point probabilities.
//
// Background. The cumulative distribution is the prefix-sum of the point
// probabilities: entry k answers the probability of an outcome at most k.
// It is the bridge between a probability list and its samples, because
// sampling from a discrete distribution (see the sampling package) is
// exactly inversion of this running list.
//
// Two things make it more than a loop. First, the same cancellation trap
// as everywhere else in floating point: entries past the last one must
// come out exactly one, but summing many probabilities each slightly off
// can drift to one plus a rounding error, and a downstream sampler that
// compares against a uniform draw can then walk off the end of the list.
// The last entry is therefore forced to exactly one, and the difference
// absorbed into it, rather than trusting the accumulated sum. Second, the
// result must be non-decreasing by construction: any implementation that
// returns a running list that dips has a bug the test can catch
// structurally, without knowing the right numbers.
//
// Steps:
//  1. Validate: every probability is at least zero.
//  2. Accumulate the running sum into the output slice.
//  3. Set the final entry to exactly one (given at least one input
//     probability; the caller supplies a distribution that sums to one).
//
// Convention (matches the tests): the returned slice has the same length
// as the input; the input is not mutated; the output is non-decreasing
// and ends at exactly one.
//
// Expected: O(n)
// Input: not mutated
func CumulativeDistribution(probs []float64) ([]float64, bool) {
	// TODO: implement
	return nil, false
}
