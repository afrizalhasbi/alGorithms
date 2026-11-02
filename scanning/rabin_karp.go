package scanning

// [hard] RabinKarp returns the index of the first occurrence of pattern in text,
// or -1.
//
// Rolling hash: compare hash values instead of strings, verify on match.
// Steps:
//  1. Hash the pattern and the first window of text.
//  2. Slide the window one character at a time: subtract the leaving
//     character's contribution, multiply by the base, add the entering
//     character. The window hash updates in O(1) — never re-hash the window.
//  3. When a window hash equals the pattern hash, compare the actual
//     substring to rule out hash collisions.
//  4. Return the index on a confirmed match; -1 if the scan ends.
//
// Notes:
//   - Roll the hash manually with a small base/modulus; do not reach for
//     a cryptographic hash. If you have already done hashing/fnv1a, notice
//     the family resemblance.
//   - The verification step on hash equality is not optional: different
//     strings can share a hash, and skipping the check gives wrong answers.
//
// Expected: O(n+m) average
// Input: not mutated
func RabinKarp(text, pattern string) int {
	// TODO: implement
	return -1
}
