package scanning

// [hard] RabinKarp returns the index of the first occurrence of pattern in text,
// or -1.
//
// Background. The naive search compares the pattern against every
// window (a stretch of the text as long as the pattern) one
// character at a time. RabinKarp instead compares a single number
// per window: the hash of the window against the hash of the
// pattern, computed once. Hashing a whole window is far more
// expensive than comparing it, so the algorithm would be pointless
// if each window had to be hashed from scratch. The rolling hash
// fixes that: when the window slides one character, the old hash is
// updated in constant time by removing the contribution of the
// character that left and adding the contribution of the one that
// entered, with a multiplication in between. So the entire text is
// hashed in about one operation per character.
//
// A hash is not a fingerprint. Two different stretches can produce
// the same hash (a collision), so a hash match means only "worth
// checking". On every hash hit the window and the pattern are
// compared character by character. Skipping that check gives wrong
// answers, not slower ones.
//
// The performance is then a bet: if hash collisions are rare, the
// expected work is one hash update per text position plus a few
// verifications. Adversarial input can force a collision at every
// position and drag the search back to quadratic, which is the
// algorithm's known weakness.
//
// Steps:
//  1. Trivial shapes first: an empty pattern matches at the very
//     start of any text; a pattern longer than the text cannot match.
//  2. Hash the pattern once, and hash the first window of the text
//     with the same hash function.
//  3. Slide the window one character at a time, updating the window
//     hash by the remove-and-add rule above. Never re-hash a window
//     from scratch.
//  4. Whenever the window hash equals the pattern hash, compare the
//     window and the pattern character by character. On a real match,
//     return the window's position.
//  5. When the window would slide past the end of the text, the
//     pattern is absent: report "not found".
//
// Notes:
//   - Roll the hash manually with a small base/modulus; do not reach
//     for a cryptographic hash. If you have already done
//     hashing/fnv1a, notice the family resemblance.
//
// Expected: O(n+m) average
// Input: not mutated
func RabinKarp(text, pattern string) int {
	// TODO: implement
	return -1
}
