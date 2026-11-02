package scanning

// [hard] KMP returns the index of the first occurrence of pattern in text, or -1.
//
// Knuth-Morris-Pratt: never re-compare characters, build a failure table.
// Steps:
//  1. Build the failure table first. For each position in the pattern,
//     record the length of the longest proper prefix of the pattern
//     that is also a suffix of the pattern up to that position — the
//     part of the pattern's start that is already known to match there.
//     Compute it by matching the pattern against itself with the same
//     discipline as the search below: when the self-comparison fails,
//     fall back through the table itself rather than starting over.
//  2. Search. Sweep the text with one cursor, never moving it
//     backwards, while tracking how much of the pattern currently
//     matches. On a match, both advance. When the whole pattern
//     matches, you have found an occurrence — return the text position
//     where that match began.
//  3. On a mismatch, do not restart. If some of the pattern was
//     matched, use the failure table to keep the longest still-valid
//     prefix matched and re-compare at the same text position. Only
//     when nothing at all is matched does the text cursor move on.
//  4. The sweep ends when the text is exhausted without a full match:
//     report "not found".
//
// Notes:
//   - The failure table is the heart. Draw "abacab" and fill it by hand
//     before coding; the table means "on failure, this much of my prefix
//     is already known to match".
//   - Each text character is read at most twice: i never moves backwards.
//
// Expected: O(n+m)
// Input: not mutated
func KMP(text, pattern string) int {
	// TODO: implement
	return -1
}
