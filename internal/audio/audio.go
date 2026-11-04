// Package audio provides a minimal mono PCM buffer with WAV output.
// It is the support layer for the audio/ exercises: tests verify
// algorithms by inspecting the samples, and -out artifacts are written
// as playable .wav files.
package audio

import (
	"encoding/binary"
	"io"
	"math"
)

// PCM is a mono int16 sample buffer.
type PCM struct {
	SampleRate int
	Samples    []int16
}

// NewPCM returns an empty buffer at the given sample rate.
func NewPCM(sampleRate int) *PCM {
	return &PCM{SampleRate: sampleRate}
}

// Silence returns a buffer of n silent samples.
func Silence(sampleRate, n int) *PCM {
	return &PCM{SampleRate: sampleRate, Samples: make([]int16, n)}
}

// Sine returns one period-accurate sine of the given frequency and
// duration, scaled to the given peak amplitude (0..32767).
func Sine(sampleRate int, freq, durSec float64, amp int16) *PCM {
	n := int(durSec * float64(sampleRate))
	s := &PCM{SampleRate: sampleRate, Samples: make([]int16, n)}
	for i := range s.Samples {
		t := float64(i) / float64(sampleRate)
		s.Samples[i] = int16(math.Round(float64(amp) * math.Sin(2*math.Pi*freq*t)))
	}
	return s
}

// Peak returns the largest absolute sample.
func (p *PCM) Peak() int16 {
	var m int16
	for _, v := range p.Samples {
		if a := int16(abs(int32(v))); a > m {
			m = a
		}
	}
	return m
}

// RMS returns the root-mean-square of the samples.
func (p *PCM) RMS() float64 {
	if len(p.Samples) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range p.Samples {
		sum += float64(v) * float64(v)
	}
	return math.Sqrt(sum / float64(len(p.Samples)))
}

// ConcatWithGap returns in followed by gapMs of silence followed by
// out, for before/after demonstration artifacts.
func ConcatWithGap(in, out *PCM, gapMs int) *PCM {
	gap := Silence(in.SampleRate, gapMs*in.SampleRate/1000)
	r := NewPCM(in.SampleRate)
	r.Samples = append(r.Samples, in.Samples...)
	r.Samples = append(r.Samples, gap.Samples...)
	r.Samples = append(r.Samples, out.Samples...)
	return r
}

// WriteWAV writes the buffer as an uncompressed 16-bit mono WAV.
func (p *PCM) WriteWAV(w io.Writer) error {
	const bitsPerSample = 16
	dataLen := len(p.Samples) * 2
	hdr := make([]uint8, 44)
	copy(hdr[0:4], "RIFF")
	binary.LittleEndian.PutUint32(hdr[4:8], uint32(36+dataLen))
	copy(hdr[8:12], "WAVE")
	copy(hdr[12:16], "fmt ")
	binary.LittleEndian.PutUint32(hdr[16:20], 16)                            // fmt chunk size
	binary.LittleEndian.PutUint16(hdr[20:22], 1)                             // PCM
	binary.LittleEndian.PutUint16(hdr[22:24], 1)                             // mono
	binary.LittleEndian.PutUint32(hdr[24:28], uint32(p.SampleRate))          // sample rate
	binary.LittleEndian.PutUint32(hdr[28:32], uint32(p.SampleRate*2))        // byte rate
	binary.LittleEndian.PutUint16(hdr[32:34], 2)                             // block align
	binary.LittleEndian.PutUint16(hdr[34:36], bitsPerSample)                 // bits
	copy(hdr[36:40], "data")
	binary.LittleEndian.PutUint32(hdr[40:44], uint32(dataLen))
	if _, err := w.Write(hdr); err != nil {
		return err
	}
	for _, v := range p.Samples {
		var b [2]uint8
		binary.LittleEndian.PutUint16(b[:], uint16(v))
		if _, err := w.Write(b[:]); err != nil {
			return err
		}
	}
	return nil
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

func clamp16(x int32) int16 {
	if x > math.MaxInt16 {
		return math.MaxInt16
	}
	if x < math.MinInt16 {
		return math.MinInt16
	}
	return int16(x)
}

