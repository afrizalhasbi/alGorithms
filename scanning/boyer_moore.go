package scanning

// [medium] BoyerMooreMajority returns the majority element of xs (appearing more than
// n/2 times) and true, or (0, false) if no majority exists.
//
// Two passes: pair off candidates, then verify.
// Steps:
//  1. Pass 1 — voting: keep a candidate and a counter. Same value as the
//     candidate: increment. Different: decrement. Counter hits 0: adopt the
//     current item as the new candidate with counter 1.
//  2. Pass 2 — verification: count actual occurrences of the candidate.
//     More than n/2: return it and true; otherwise return false.
//     The second pass is what makes the algorithm honest — without it,
//     "A A B" style inputs produce a bogus winner.
//
// Notes:
//   - Easy to code, subtle to prove: the majority element survives being
//     paired off against everything else, because it outnumbers them all.
//   - O(n) time, O(1) space, single pass over the input twice.
//
// Expected: O(n)
// Input: not mutated
func BoyerMooreMajority(xs []int) (int, bool) {
	// TODO: implement
	return 0, false
}
