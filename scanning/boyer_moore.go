package scanning

// [medium] BoyerMooreMajority returns the majority element of xs (appearing more than
// n/2 times) and true, or (0, false) if no majority exists.
//
// Two passes: pair off candidates, then verify.
// Steps:
//  1. Pass one — voting. Sweep the slice once, carrying a current
//     candidate and a counter. Before the first item there is no
//     candidate: the first item becomes the candidate with a count of
//     one. After that, each item either matches the candidate (count
//     goes up), differs from it (count goes down), or arrives when the
//     count has fallen to zero — in which case it displaces the old
//     candidate and the count restarts at one.
//  2. When the sweep finishes, the surviving candidate is the only
//     value that could possibly be a majority — but that is only a
//     claim at this point; nothing has proven it.
//  3. Pass two — verification. Sweep the whole slice again, simply
//     counting how often the candidate actually occurs. The scan is a
//     plain count: it never stops early. If the candidate occurs more
//     than half the time, return it as the majority; if it occurs
//     exactly half the time or less, there is no majority.
//  4. An empty slice has no candidate to vote on: report "no majority"
//     without any voting at all.
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
