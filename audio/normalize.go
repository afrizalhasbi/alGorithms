package audio

import "alGorithms/internal/audio"

// [easy] Normalize returns the buffer rescaled so its peak sample
// hits the target amplitude exactly, preserving the waveform shape.
//
// Background. Normalizing is uniform scaling: every sample is
// multiplied by the same factor target/peak, where peak is the
// largest absolute sample currently in the buffer. Because the factor
// is a scalar, differences between samples scale by it too — zero
// crossings stay zero crossings, and the ratio of any two samples is
// unchanged. That is the shape-preservation the test verifies: output
// samples must equal input samples times the factor (up to rounding),
// not any other transformation.
//
// Two edge cases matter. Silence (peak 0) has no defined scale factor
// and must return the buffer unchanged rather than divide by zero.
// And scaling can clip: a factor above 1 on an already loud buffer
// would push samples past int16 range, so the operation must either
// refuse to boost (scale only down) or clamp — here the contract is
// scale-only-down: normalize never amplifies, target above the
// current peak is a no-op.
//
// Definitions:
//  - peak: max |sample| over the buffer, 0 for silence.
//  - scale factor: target/peak as a float, computed once.
//  - shape preservation: out[i]/out[j] == in[i]/in[j] for all pairs,
//     up to int16 rounding.
//
// Steps:
//  1. Find the peak. If it is 0, or target >= peak, return an
//     unchanged copy.
//  2. Multiply each sample by target/peak and round to int16.
//
// Notes:
//   - The test checks the output peak equals the target exactly, the
//     sample-by-sample ratio against a float reference stays within
//     rounding error, silence is handled, and boosting is refused.
//
// Expected: O(n) time, O(n) space
// Input: not mutated
func Normalize(src *audio.PCM, target int16) *audio.PCM {
	// TODO: implement
	return nil
}
