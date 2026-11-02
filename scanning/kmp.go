package scanning

// [hard] KMP returns the index of the first occurrence of pattern in text, or -1.
//
// Knuth-Morris-Pratt: never re-compare characters, build a failure table.
// Steps:
//  1. Build the failure table for the pattern: fail[i] = length of the longest
//     proper prefix of pattern that is also a suffix of pattern[0:i].
//     Compute it with the same matching loop you use for the search — the
//     pattern is matched against itself.
//  2. Scan text with pointers i (text) and j (pattern).
//     Match: advance both. When j reaches len(pattern), return i-j.
//     Mismatch with j > 0: j = fail[j-1]; do NOT move i, do NOT restart j.
//     Mismatch with j == 0: advance i.
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
