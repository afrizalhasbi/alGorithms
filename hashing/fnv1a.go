package hashing

// [easy] FNV1a returns the 32-bit FNV-1a hash of s.
//
// Background. A hash function turns a string into a fixed-size
// number, with the property that similar inputs produce unrelated
// outputs. It is used to scatter keys across a hash table.
// FNV-1a is one of the simplest usable hashes: a running value
// starts at a fixed constant and is transformed once per input byte.
// Each byte is mixed in with an exclusive-or and then the whole
// value is multiplied by a fixed odd constant. The multiplication is
// what spreads the change across all bits of the value, so after a
// few bytes every output bit depends on many input bytes.
//
// The order matters and is the only real trap here: the exclusive-or
// comes before the multiply. Doing it the other way round gives the
// older FNV-1, which has weaker scattering and fails the test
// vectors.
//
// Definitions:
//  - offset basis / prime: the two FNV constants, 2166136261 and
//     16777619. Exact values, checked by test vectors.
//  - running value: the 32-bit accumulator. Arithmetic wraps
//     naturally at 32 bits (uint32); no masking is needed.
//  - the mix: exclusive-or the byte into the running value, THEN
//     multiply by the prime. The order is the entire difference
//     between FNV-1a and the older FNV-1.
//
// Steps:
//  1. Start the running value at the offset basis constant.
//  2. For each byte of the input, in order: exclusive-or the byte
//     into the running value, then multiply by the prime constant,
//     letting the arithmetic wrap around at 32 bits.
//  3. After the last byte, the running value is the hash. An empty
//     input skips the loop and returns the basis unchanged.
//
// Notes:
//   - The test checks known FNV-1a vectors, so constants must be
//     exact: offset basis 2166136261, prime 16777619.
//   - No complexity test: O(len(s)) per key, no shape dependence.
//
// Expected: O(len(s)) per key
// Input: not mutated
func FNV1a(s string) uint32 {
	sb := []byte(s)
	out := uint32(2166136261)
	for _, b := range sb {
		out = (uint32(b) ^ out) * 16777619
	}
	// TODO: implement
	return out
}
