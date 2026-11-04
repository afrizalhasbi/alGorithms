package audio

import "alGorithms/internal/audio"

// [easy] SineGen synthesizes a sine wave of the given frequency and
// duration at a fixed sample rate.
//
// Background. Digital audio stores a wave as one amplitude sample per
// tick of a fixed clock: the sample rate, typically 44100 Hz. A sine
// of frequency f Hz is then just sin(2π·f·t) evaluated at each sample
// time t = i / sampleRate, scaled to the wanted amplitude. One full
// cycle spans 1/f seconds, so a 440 Hz tone crossing zero twice per
// cycle crosses the zero axis 880 times per second — which is the
// cheap way to verify the frequency without a spectral transform:
// count sign changes and compare against f·2·dur.
//
// Definitions:
//  - sample: one int16 amplitude value; 0 is silence, ±32767 the
//     full-scale limits.
//  - sample rate: samples per second; sample i happens at time
//     i / sampleRate.
//  - zero crossing: two consecutive samples with opposite signs (or a
//     zero followed by a nonzero of either sign).
//  - amplitude: the peak |sample| the synthesized wave should reach.
//
// Steps:
//  1. Allocate n = round(durSec · sampleRate) samples.
//  2. For each i, write round(amp · sin(2π · freq · i / sampleRate)).
//
// Notes:
//   - The test checks the zero-crossing count against 2·f·dur and
//     the RMS against amp/√2, the theoretical RMS of a sine. Both
//     checks fail for a wrong frequency or a wrong amplitude scale,
//     so both are required.
//
// Expected: O(n) time, O(n) space
func SineGen(sampleRate int, freq, durSec float64, amp int16) *audio.PCM {
	// TODO: implement
	return nil
}
