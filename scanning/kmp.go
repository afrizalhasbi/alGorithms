package scanning

// [hard] KMP returns the index of the first occurrence of pattern in text, or -1.
//
// Background. Searching naively means: compare the pattern against
// the text at each starting position, and after any mismatch, slide
// the pattern one position and start comparing from the beginning of
// the pattern again. The weakness is that a mismatch throws away
// everything the comparisons already learned. KMP never does that.
//
// The failure table is the memory of what was learned. For each
// position in the pattern it records: if a mismatch happens just
// after this position, how much of the matched stretch is still
// usable? Concretely it is the length of the longest proper prefix
// of the pattern that is also a suffix of the stretch matched so
// far. Intuition: the end of a matched stretch looks like the start
// of the pattern, so when the match breaks, the pattern can be
// shifted only as far as that resemblance allows and part of it is
// already lined up. Drawing the table by hand for a small pattern
// makes this much clearer.
//
// The search then reads the text once, front to back, never moving
// backwards. A text character is compared, and whether it matches or
// not, the text position never goes back; only the amount of pattern
// currently matched jumps backwards through the failure table.
//
// Steps:
//  1. Build the failure table for the pattern before searching. The
//     computation is the same matching discipline as the search
//     itself, with the pattern compared against itself: on a mismatch
//     mid-comparison, fall back through the table rather than
//     restarting.
//  2. Search: walk the text one character at a time while tracking
//     how much of the pattern is currently matched. A matching
//     character extends the match. When the whole pattern is matched,
//     an occurrence has been found at the position where the match
//     began; for a first occurrence, return immediately.
//  3. On a mismatch with part of the pattern matched, look up how
//     much of the match survives in the failure table and re-compare
//     at the same text position. On a mismatch with nothing matched,
//     the text position moves on.
//  4. If the text is exhausted without a full match, report "not
//     found".
//
// Notes:
//   - Each text character is read at most twice, which is why the
//     running time is proportional to text plus pattern length.
//
// Expected: O(n+m)
// Input: not mutated
func KMP(text, pattern string) int {
	// TODO: implement
	return -1
}
