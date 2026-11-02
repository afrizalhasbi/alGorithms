package numerics

// [medium] Welford returns the mean and the sample variance of xs in a
// single pass, using Welford's recurrence.
//
// Background. The textbook formulas need two passes: one to get the mean,
// one to accumulate squared deviations from it. Worse, the obvious
// one-pass shortcut — accumulate the sum of squares and the sum, then
// subtract — is numerically disastrous: it computes the variance of two
// huge nearly-equal quantities, and catastrophic cancellation throws away
// almost every significant digit when the data straddles a large offset
// (variance of [1e9, 1e9 + 1, 1e9 + 2] should be 1; the sum-of-squares
// route returns garbage).
//
// Welford's recurrence processes items one at a time and never forms large
// differences of large squares. After each item it keeps exactly two
// values: the mean of everything seen so far, and an accumulator of
// squared deviations from the current mean (the sum of squared distances
// of the seen items to their running mean). When a new item arrives, the
// mean shifts by the deviation of the item from the old mean, divided by
// the new count; the accumulator grows by the product of that deviation
// and the distance of the item from the new mean — the cross term that
// makes the algebra work out exactly.
//
// The variance is then the accumulated sum divided by (count minus one)
// for the sample variance, or by the count for the population variance.
// The tests check both, plus the classic cancellation case above, where
// the naive two-sum formula provably fails.
//
// Steps:
//  1. Carry three values through the pass: the count, the running mean,
//     and the running accumulator of squared deviations (the last two
//     start at zero).
//  2. For each item: increment the count; compute how far the item sits
//     from the current mean; update the mean by shifting it partway toward
//     the item (the shift shrinks as the count grows); then add the
//     product of the two deviations — item versus new mean, item versus
//     old mean — to the accumulator.
//  3. Return the mean, and the accumulator scaled for sample and for
//     population variance.
//
// Convention (matches the tests): the empty input has no mean or variance;
// a single item has mean but zero sample variance only in the population
// sense (the sample variance of one observation is undefined; the tests
// only exercise inputs of at least two items, and expect the sample
// variance call to handle it as zero).
//
// Expected: O(n), one pass
// Input: not mutated
func Welford(xs []float64) (mean, varSample, varPopulation float64) {
	// TODO: implement
	return 0, 0, 0
}
