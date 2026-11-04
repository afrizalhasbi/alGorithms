// Package bitmap provides a minimal RGB raster buffer with PPM
// output. It is the support layer for the images/ exercises: tests
// verify algorithms by inspecting the buffer, and -out artifacts are
// written as viewable .ppm files (P6, viewable with most image
// viewers, e.g. `feh`, or converted with ImageMagick).
package bitmap

import (
	"fmt"
	"io"
)

// RGB is a dense WxH raster, 3 bytes per pixel, row-major, no alpha.
type RGB struct {
	W, H int
	Pix  []uint8 // len == W*H*3, pixel (x,y) at (y*W+x)*3
}

// NewRGB returns a black WxH buffer.
func NewRGB(w, h int) *RGB {
	return &RGB{W: w, H: h, Pix: make([]uint8, w*h*3)}
}

// Set writes the pixel at (x,y). Out-of-range writes are ignored so
// rasterizer exercises never panic on overshoot.
func (b *RGB) Set(x, y int, r, g, bb uint8) {
	if x < 0 || y < 0 || x >= b.W || y >= b.H {
		return
	}
	i := (y*b.W + x) * 3
	b.Pix[i] = r
	b.Pix[i+1] = g
	b.Pix[i+2] = bb
}

// At reads the pixel at (x,y). Out-of-range reads are black.
func (b *RGB) At(x, y int) (r, g, bb uint8) {
	if x < 0 || y < 0 || x >= b.W || y >= b.H {
		return 0, 0, 0
	}
	i := (y*b.W + x) * 3
	return b.Pix[i], b.Pix[i+1], b.Pix[i+2]
}

// Fill sets every pixel to one color.
func (b *RGB) Fill(r, g, bb uint8) {
	for i := 0; i < len(b.Pix); i += 3 {
		b.Pix[i] = r
		b.Pix[i+1] = g
		b.Pix[i+2] = bb
	}
}

// WritePPM writes the buffer as binary PPM (P6).
func (b *RGB) WritePPM(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "P6\n%d %d\n255\n", b.W, b.H); err != nil {
		return err
	}
	_, err := w.Write(b.Pix)
	return err
}

// SideBySide returns a new buffer with a placed left of b, separated
// by a 2px black divider column. Buffers must have equal heights.
func SideBySide(a, b *RGB) *RGB {
	if a.H != b.H {
		panic(fmt.Sprintf("SideBySide: heights differ: %d vs %d", a.H, b.H))
	}
	const gap = 2
	out := NewRGB(a.W+gap+b.W, a.H)
	for y := 0; y < a.H; y++ {
		for x := 0; x < a.W; x++ {
			r, g, bb := a.At(x, y)
			out.Set(x, y, r, g, bb)
		}
		for x := 0; x < b.W; x++ {
			r, g, bb := b.At(x, y)
			out.Set(a.W+gap+x, y, r, g, bb)
		}
	}
	return out
}
