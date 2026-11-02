// Package statistics holds the statistical algorithm exercises: discrete
// probability, density estimation, and expectation-maximization on
// discrete data.
package statistics

// [easy] DiscreteExpectation returns the expectation of a discrete random
// variable given its support and its probabilities.
//
// Background. The expectation of a discrete variable is the sum of every
// value weighted by its probability. The only discipline in the exercise
// is the contract: values and probabilities are parallel lists, the
// probabilities must sum to one, and a probability outside the zero-one
// range means the input is not a distribution at all and should be
// reported rather than quietly summed.
//
// Steps:
//  1. Validate: the two lists have the same length, every probability is
//     at least zero and at most one, and they sum to one within a small
//     tolerance (floating point is allowed a little slack).
//  2. Accumulate value times probability over all entries.
//
// Convention (matches the tests): the empty distribution is invalid. The
// tolerance on the sum-to-one check is one part in a million, relative.
//
// Expected: O(n)
// Input: not mutated
func DiscreteExpectation(values, probs []float64) (float64, bool) {
	// TODO: implement
	return 0, false
}
