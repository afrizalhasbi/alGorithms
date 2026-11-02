package hashing

// [easy] FNV1a returns the 32-bit FNV-1a hash of s.
//
// A tiny non-cryptographic hash: multiply-then-xor over each byte.
//
// Steps:
// 1. Start with offset basis 2166136261.
// 2. For each byte b of s: hash = (hash ^ b) * 16777619  (uint32 wraparound).
// 3. Return the final hash.
//
// Notes:
//   - The xor comes *before* the multiply (that is the "1a" in FNV-1a);
//     swapping the two gives the older, weaker FNV-1.
//   - Every input byte changes the result: avalanche by construction.
//
// Expected: O(len(s)) per key
// Input: not mutated
func FNV1a(s string) uint32 {
	// TODO: implement
	return 0
}
