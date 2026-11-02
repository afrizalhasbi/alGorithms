package scanning

// [hard] RabinKarp returns the index of the first occurrence of pattern in text,
// or -1.
//
// Rolling hash: compare hash values instead of strings, verify on match.
// Steps:
//  1. Handle the trivial shapes up front: an empty pattern matches at
//     the very start of any text (including empty text); a pattern
//     longer than the text can never match — report "not found".
//  2. Hash the pattern once, and hash the first window of the text
//     (the opening stretch as long as the pattern) with the same hash.
//  3. Slide the window one character at a time. Each slide updates the
//     window hash in constant time: remove the contribution of the
//     character leaving at the front and add the character entering at
//     the back. Never re-hash a whole window.
//  4. After each slide, compare the window hash with the pattern hash.
//     Equal hashes do not prove a match — different strings can share
//     a hash — so on a hash hit, compare the actual window against the
//     pattern character by character, and only then declare a match and
//     return the window's position.
//  5. The slide continues while the window still fits inside the text.
//     Once the window would run off the end, the pattern is absent:
//     report "not found".
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
