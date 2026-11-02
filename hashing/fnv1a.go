package hashing

// [easy] FNV1a returns the 32-bit FNV-1a hash of s.
//
// A tiny non-cryptographic hash: multiply-then-xor over each byte.
//
// Steps:
//  1. Begin from a fixed starting value (the offset basis).
//  2. Take the input one byte at a time, in order. For each byte: first
//     mix the byte into the current value with an exclusive-or, then
//     multiply the result by a fixed odd constant, letting the
//     arithmetic wrap around at the word size. The xor comes before
//     the multiply — that ordering is the "1a" in FNV-1a, and swapping
//     it gives the older, weaker FNV-1.
//  3. When every byte has been mixed in, the current value is the hash.
//     An empty input skips the loop entirely and returns the basis
//     unchanged.
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
