package scanning

// [medium] BoyerMooreMajority returns the majority element of xs (appearing more than
// n/2 times) and true, or (0, false) if no majority exists.
//
// Background. A majority element is a value that occupies more than
// half the positions of the slice. Strictly more than half: a value
// appearing exactly half the time is not a majority (a slice like
// one, one, two, two has none). The obvious approach counts every
// distinct value, which needs extra memory for the counts.
//
// The Boyer-Moore vote gets the same answer in constant memory, in
// two passes. The first pass maintains one candidate and a counter,
// and works like pairing off items against each other: each item
// either agrees with the candidate (the counter goes up) or cancels
// one vote of it (the counter goes down). When the counter is
// exhausted, the next item starts a new candidacy. If a majority
// exists, it survives this process, because every one of its items
// that gets cancelled took a non-majority item down with it, and
// there are more of them than of everything else combined.
//
// The catch: the first pass only proves who would win if a majority
// existed. Without a majority, the survivor is arbitrary, a leftover
// of the pairing. So a second pass is required, simply counting how
// often the survivor actually occurs. That check is what makes the
// algorithm honest; skipping it gives wrong answers on inputs like
// A, A, B.
//
// Steps:
//  1. An empty slice has no candidate: report "no majority".
//  2. First pass, over the whole slice: no candidate yet means the
//     current item becomes the candidate with count one; otherwise
//     the count moves up on a match and down on a mismatch, and a
//     count reaching zero hands candidacy to the next item.
//  3. Second pass, over the whole slice: count occurrences of the
//     candidate. The scan never stops early.
//  4. More than half: return the candidate and true. Otherwise return
//     false. An empty slice also returns false.
//
// Notes:
//   - O(n) time, O(1) space, two passes over the input.
//
// Expected: O(n)
// Input: not mutated
func BoyerMooreMajority(xs []int) (int, bool) {
	// TODO: implement
	return 0, false
}
