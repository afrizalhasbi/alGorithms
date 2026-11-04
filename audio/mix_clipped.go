package audio

import "alGorithms/internal/audio"

// [easy] MixClipped returns the sample-wise sum of two buffers,
// clamped to the int16 range.
//
// Background. Mixing two tracks is adding their samples: the air
// pressure contributions of two sounds add physically, so their
// digital representations add. The sum can exceed int16 range — two
// full-scale tones sum to twice full scale — and something must
// happen at the ceiling. Clipping is the simplest policy: every sum
// above 32767 becomes 32767, every sum below -32768 becomes -32768.
// The region below the ceiling is exact linear addition; the region
// at the ceiling is flat-topped distortion, audible as harshness but
// predictable and testable.
//
// Buffers of different lengths mix over the shorter length and the
// result has the longer length, with the tail of the longer buffer
// passed through unchanged — a mixer cannot drop audio that has no
// partner to clash with.
//
// Definitions:
//  - sample-wise: output[i] depends only on a[i] and b[i].
//  - clipping: clamping the sum to [MinInt16, MaxInt16].
//  - pass-through tail: samples of the longer buffer beyond the
//     shorter one appear unmodified in the output.
//
// Steps:
//  1. Allocate the result at the longer length.
//  2. Over the overlap, add the two samples as ints and clamp.
//  3. Beyond the overlap, copy the longer buffer's samples.
//
// Notes:
//   - The test checks the linear zone sums exactly, the ceiling
//     clamps both directions, the tail passes through for both
//     length orders, and mixing with silence is the identity.
//
// Expected: O(n) time, O(n) space
// Input: not mutated
func MixClipped(a, b *audio.PCM) *audio.PCM {
	// TODO: implement
	return nil
}
