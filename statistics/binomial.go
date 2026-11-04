package statistics

// [easy] BinomialPMF returns the probability of exactly k successes in n
// independent trials with success probability p, using the recurrence for
// the point probabilities rather than factorials.
//
// Background. The binomial point probability is the number of orderings
// of the successes, times the probability of any one ordering. Written
// with factorials it overflows long before the probabilities do: two
// hundred factorial is already beyond float64's range even though the
// probability it should appear in is a perfectly ordinary number. The
// cure is to never compute the factorials at all.
//
// The point probabilities form a sequence in k with a simple relation
// between neighbors. Divide the formula for k by the formula for k-1 and
// almost everything cancels: the ratio involves only n, k and p, one
// turn of the success and failure probabilities each. So start from the
// zero-successes probability — every trial fails, one clean product — and
// climb to k by multiplying that ratio repeatedly. Every intermediate
// value is itself a legitimate point probability, no wider than the
// answers themselves.
//
// Symmetry is a free bonus: the distribution of k successes with success
// probability p is the mirror image of n-k successes with probability
// one minus p, so the climb can always start from whichever end of the
// distribution is closer to the target k and take the shorter walk.
//
// Definitions:
//  - point probability: the probability of exactly k successes. The PMF
//     is the list of these over all k.
//  - the zero-successes point: every trial fails; the clean product of
//     n copies of (one minus p). The all-successes point is the mirror
//     with p in place of (one minus p).
//  - ratio: the exact quotient of the point probability at k over the
//     one at k-1, obtained by dividing the formulas and cancelling. It
//     involves n, k and p only.
//  - climb: starting at one end's point probability and applying the
//     ratio repeatedly until the target k is reached. The direction is
//     whichever end is nearer to k.
//
// Steps:
//  1. Validate: n at least zero, p in the zero-one range, k in the
//     zero-to-n range. Out-of-range k is a probability of zero (no
//     successes in zero trials), not an error; a bad p is an error.
//  2. Choose the nearer end: either start at the all-failures point
//     probability (one times the failure probability raised to n) and
//     climb up with the ratio above, or start at the all-successes point
//     (with the roles of p and one-minus-p swapped) and climb down with
//     the same relation mirrored.
//  3. Return the reached point probability.
//
// Convention (matches the tests): the fixtures span n up to one hundred
// with extreme p values, where the factorial formula would overflow and
// the recurrence stays exact; the reference values were computed with
// exact integer arithmetic on the same formula.
//
// Expected: O(min(k, n-k))
func BinomialPMF(n, k int, p float64) (float64, bool) {
	// TODO: implement
	return 0, false
}
