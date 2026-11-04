package images

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"testing"

	"alGorithms/internal/bitmap"
)

var outFlag = flag.Bool("out", false, "write before/after artifacts to internal/output")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

// okPrint prints a success line styled like the failure lines: no
// file:line prefix or indentation, always visible.
func okPrint(format string, args ...any) {
	fmt.Printf("ok: "+format+"\n", args...)
}

// artifactDir walks up from the working directory to the module root
// (the dir containing go.mod) and returns <root>/internal/output.
// go test runs the binary with CWD set to the package dir, so a bare
// relative path would land inside images/ or audio/.
func artifactDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(dir + "/go.mod"); err == nil {
			return dir + "/internal/output", nil
		}
		parent := dir[:strings.LastIndexByte(dir, '/')]
		if parent == "" || parent == dir {
			return "", fmt.Errorf("go.mod not found above %s", dir)
		}
		dir = parent
	}
}

// writeArtifact writes a before/after PPM into internal/output and
// prints the path. It is a no-op unless the test runs with -out.
func writeArtifact(t *testing.T, img *bitmap.RGB, name string) {
	t.Helper()
	if !*outFlag {
		return
	}
	outDir, err := artifactDir()
	if err != nil {
		t.Fatalf("locating internal/output: %v", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		t.Fatalf("creating internal/output: %v", err)
	}
	f, err := os.Create(outDir + "/" + name + ".ppm")
	if err != nil {
		t.Fatalf("creating artifact: %v", err)
	}
	defer f.Close()
	if err := img.WritePPM(f); err != nil {
		t.Fatalf("writing artifact: %v", err)
	}
	okPrint("wrote internal/output/%s.ppm", name)
}

// floodFillDemo builds the standard demo scene: a red disk and a
// separate red ring on white, so the fill visibly recolors one shape
// and provably leaves the other untouched.
func floodFillDemo() *bitmap.RGB {
	img := bitmap.NewRGB(64, 48)
	img.Fill(255, 255, 255)
	for y := range img.H {
		for x := range img.W {
			dx, dy := x-16, y-24
			if dx*dx+dy*dy <= 100 {
				img.Set(x, y, 220, 40, 40) // solid disk, seed target
			}
			rx, ry := x-46, y-24
			d := rx*rx + ry*ry
			ang := math.Atan2(float64(ry), float64(rx))
			if d >= 64 && d <= 144 && !(math.Abs(ang) < 0.5) {
				img.Set(x, y, 220, 40, 40) // open arc, same color but disconnected
			}
		}
	}
	return img
}

func TestFloodFill(t *testing.T) {
	// case 1: the disk component is fully recolored, the ring keeps
	// its color because it is not 4-connected to the seed.
	img := floodFillDemo()
	before := floodFillDemo()
	FloodFill(img, 16, 24, 40, 160, 90)

	count := 0
	for y := range img.H {
		for x := range img.W {
			r, g, b := img.At(x, y)
			br, bg, bb := before.At(x, y)
			dx, dy := x-16, y-24
			inDisk := dx*dx+dy*dy <= 100
			wasRed := br == 220 && bg == 40 && bb == 40
			switch {
			case inDisk && (r != 40 || g != 160 || b != 90):
				t.Fatalf("disk pixel (%d,%d) not recolored", x, y)
			case wasRed && !inDisk && (r != 220 || g != 40 || b != 40):
				t.Fatalf("arc pixel (%d,%d) was recolored but is disconnected", x, y)
			case !wasRed && !inDisk && (r != br || g != bg || b != bb):
				t.Fatalf("background pixel (%d,%d) changed", x, y)
			}
			if inDisk {
				count++
			}
		}
	}
	okPrint("FloodFill recolored %d disk pixels, arc untouched", count)

	// case 2: target color equals seed color is a no-op.
	img = floodFillDemo()
	FloodFill(img, 16, 24, 220, 40, 40)
	for y := range img.H {
		for x := range img.W {
			g1, g2, g3 := img.At(x, y)
			b1, b2, b3 := before.At(x, y)
			if g1 != b1 || g2 != b2 || g3 != b3 {
				t.Fatalf("no-op fill changed pixel (%d,%d)", x, y)
			}
		}
	}
	okPrint("FloodFill same-color no-op leaves image unchanged")

	// case 3: fill from white recolors the entire background and both
	// shapes stop it — the whole non-shape area becomes one new color.
	img = floodFillDemo()
	FloodFill(img, 1, 1, 25, 25, 60)
	changed := 0
	for y := range img.H {
		for x := range img.W {
			r, g, b := img.At(x, y)
			if r == 25 && g == 25 && b == 60 {
				changed++
			}
		}
	}
	// Count the exact number of non-white pixels in the original: the
	// fill must recolor every white pixel and no shape pixel.
	white := img.W*img.H - 0
	for i := 0; i < len(before.Pix); i += 3 {
		if before.Pix[i] != 255 || before.Pix[i+1] != 255 || before.Pix[i+2] != 255 {
			white--
		}
	}
	if changed != white {
		t.Fatalf("background fill recolored %d pixels, want %d (all white pixels)", changed, white)
	}
	okPrint("FloodFill from background recolored %d pixels (all white)", changed)

	writeArtifact(t, bitmap.SideBySide(before, img), "floodfill_before_after")
}

// gradient builds a WxH horizontal gradient with a light vertical
// wave, the standard input for filters and dithering.
func gradient(w, h int) *bitmap.RGB {
	img := bitmap.NewRGB(w, h)
	for y := range h {
		for x := range w {
			v := uint8(127 + 60*math.Sin(2*math.Pi*float64(x)/float64(w)) + 40*math.Sin(2*math.Pi*3*float64(y)/float64(h)))
			img.Set(x, y, v, v, v)
		}
	}
	return img
}

// checkEqual fails if the two buffers differ anywhere.
func checkEqual(t *testing.T, name string, a, b *bitmap.RGB) {
	t.Helper()
	if a.W != b.W || a.H != b.H {
		t.Fatalf("%s: dimensions differ: %dx%d vs %dx%d", name, a.W, a.H, b.W, b.H)
	}
	for y := range a.H {
		for x := range a.W {
			ar, ag, ab := a.At(x, y)
			br, bg, bb := b.At(x, y)
			if ar != br || ag != bg || ab != bb {
				t.Fatalf("%s: pixel (%d,%d) differs", name, x, y)
			}
		}
	}
}

// isColor reports whether pixel (x,y) equals the given color.
func isColor(img *bitmap.RGB, x, y int, r, g, b uint8) bool {
	pr, pg, pb := img.At(x, y)
	return pr == r && pg == g && pb == b
}

// meanVar computes the global mean and variance of all channels.
func meanVar(img *bitmap.RGB) (float64, float64) {
	var sum, sum2 float64
	n := float64(len(img.Pix))
	for _, v := range img.Pix {
		sum += float64(v)
		sum2 += float64(v) * float64(v)
	}
	mean := sum / n
	return mean, sum2/n - mean*mean
}

func TestRotate90Flip(t *testing.T) {
	src := floodFillDemo()

	// Rotate90: dimensions swap, corners map exactly.
	rot := Rotate90(src)
	if rot.W != src.H || rot.H != src.W {
		t.Fatalf("Rotate90 dimensions %dx%d, want %dx%d", rot.W, rot.H, src.H, src.W)
	}
	r, g, b := src.At(0, 0) // top-left source
	if gr, gg, gb := rot.At(src.H-1, 0); gr != r || gg != g || gb != b {
		t.Fatalf("Rotate90 moved top-left to the wrong corner")
	}
	r, g, b = src.At(src.W-1, 0) // top-right source
	if gr, gg, gb := rot.At(src.H-1, src.W-1); gr != r || gg != g || gb != b {
		t.Fatalf("Rotate90 moved top-right to the wrong corner")
	}
	okPrint("Rotate90 dims %dx%d→%dx%d, corners map clockwise", src.W, src.H, rot.W, rot.H)

	// Four rotations are the identity.
	back := Rotate90(Rotate90(Rotate90(rot)))
	checkEqual(t, "Rotate90^4", back, src)
	okPrint("Rotate90 four times is the identity")

	// FlipH: dimensions keep, two flips are the identity, corners swap.
	flip := FlipH(src)
	if flip.W != src.W || flip.H != src.H {
		t.Fatalf("FlipH changed dimensions")
	}
	r, g, b = src.At(0, 24)
	if gr, gg, gb := flip.At(src.W-1, 24); gr != r || gg != g || gb != b {
		t.Fatalf("FlipH did not mirror row 24")
	}
	checkEqual(t, "FlipH^2", FlipH(flip), src)
	okPrint("FlipH mirrors and is its own inverse")

	writeArtifact(t, bitmap.SideBySide(src, rot), "rotate90_before_after")
}

func TestBoxBlur(t *testing.T) {
	src := gradient(96, 64)

	blurred := BoxBlur(src, 3)
	if blurred == nil {
		t.Fatal("BoxBlur returned nil")
	}

	// A flat image stays exactly flat.
	flat := bitmap.NewRGB(32, 32)
	flat.Fill(90, 120, 200)
	flatBlurred := BoxBlur(flat, 2)
	checkEqual(t, "flat blur", flatBlurred, flat)
	okPrint("BoxBlur flat image stays flat")

	// Mean preserved (up to edge clipping), variance drops.
	m0, v0 := meanVar(src)
	m1, v1 := meanVar(blurred)
	if math.Abs(m0-m1) > 2.0 {
		t.Fatalf("BoxBlur changed mean: %.2f → %.2f", m0, m1)
	}
	if v1 >= v0 {
		t.Fatalf("BoxBlur did not reduce variance: %.1f → %.1f", v0, v1)
	}
	okPrint("BoxBlur mean %.1f→%.1f, variance %.0f→%.0f", m0, m1, v0, v1)

	// Source untouched.
	fresh := gradient(96, 64)
	checkEqual(t, "blur input", src, fresh)

	writeArtifact(t, bitmap.SideBySide(src, blurred), "boxblur_before_after")
}

func TestBresenhamLine(t *testing.T) {
	cases := [][4]int{
		{2, 2, 60, 40},  // shallow descending
		{2, 40, 60, 4},  // steep ascending
		{30, 2, 30, 44}, // vertical
		{2, 24, 60, 24}, // horizontal
		{60, 40, 2, 2},  // reversed: same pixel set
	}

	img := bitmap.NewRGB(64, 48)
	for _, c := range cases {
		BresenhamLine(img, c[0], c[1], c[2], c[3], 30, 30, 30)
	}

	// Verify each line: endpoints drawn, pixels near the ideal
	// segment, 8-connected. Colored pixels are collected per line by
	// drawing into a fresh buffer.
	for i, c := range cases {
		li := bitmap.NewRGB(64, 48)
		BresenhamLine(li, c[0], c[1], c[2], c[3], 30, 30, 30)
		lr, lg, lb := li.At(c[0], c[1])
		if lr != 30 || lg != 30 || lb != 30 {
			t.Fatalf("line %d: start point not drawn", i)
		}
		lr, lg, lb = li.At(c[2], c[3])
		if lr != 30 || lg != 30 || lb != 30 {
			t.Fatalf("line %d: end point not drawn", i)
		}
		dx, dy := float64(c[2]-c[0]), float64(c[3]-c[1])
		length := math.Max(math.Abs(dx), math.Abs(dy))
		prev := [2]int{-100, -100}
		drawn := 0
		for y := range 48 {
			for x := range 64 {
				if isColor(li, x, y, 30, 30, 30) {
					continue
				}
				drawn++
				// Distance from the ideal line (normalized projection).
				tt := (float64(x-c[0])*dx + float64(y-c[1])*dy) / (dx*dx + dy*dy)
				if tt < -0.01 || tt > 1.01 {
					t.Fatalf("line %d: pixel (%d,%d) outside segment", i, x, y)
				}
				px, py := float64(c[0])+tt*dx, float64(c[1])+tt*dy
				d := math.Hypot(float64(x)-px, float64(y)-py)
				if d > 0.71 {
					t.Fatalf("line %d: pixel (%d,%d) is %.2f off the ideal line", i, x, y, d)
				}
				if math.Abs(float64(x-prev[0])) > 1 || math.Abs(float64(y-prev[1])) > 1 {
					t.Fatalf("line %d: gap between (%d,%d) and (%d,%d)", i, prev[0], prev[1], x, y)
				}
				prev = [2]int{x, y}
			}
		}
		if i < 4 && drawn < int(length) {
			t.Fatalf("line %d: only %d pixels for length %.0f", i, drawn, length)
		}
	}
	// Reversed line colors the same set (symmetry case 5 vs case 1).
	a := bitmap.NewRGB(64, 48)
	bb := bitmap.NewRGB(64, 48)
	BresenhamLine(a, 2, 2, 60, 40, 30, 30, 30)
	BresenhamLine(bb, 60, 40, 2, 2, 30, 30, 30)
	checkEqual(t, "line symmetry", a, bb)
	okPrint("BresenhamLine endpoints, accuracy ≤ 0.71px, connectivity, symmetry ok")

	writeArtifact(t, bitmap.SideBySide(floodFillDemo(), img), "bresenham_before_after")
}

func TestMidpointCircle(t *testing.T) {
	const rad = 15
	img := bitmap.NewRGB(64, 48)
	const cx, cy = 32, 24
	MidpointCircle(img, cx, cy, rad, 40, 90, 200)

	drawn := 0
	for y := range 48 {
		for x := range 64 {
			if !isColor(img, x, y, 40, 90, 200) {
				continue
			}
			drawn++
			d := math.Hypot(float64(x-cx), float64(y-cy))
			if math.Abs(d-rad) > 1 {
				t.Fatalf("pixel (%d,%d) at distance %.2f, radius %d", x, y, d, rad)
			}
		}
	}
	// 8-fold symmetry: the drawn set mirrors onto itself.
	symErr := 0
	for y := range 48 {
		for x := range 64 {
			c := isColor(img, x, y, 40, 90, 200)
			mc := isColor(img, 2*cx-x, 2*cy-y, 40, 90, 200)
			if c != mc {
				symErr++
			}
		}
	}
	if symErr != 0 {
		t.Fatalf("8-fold symmetry violated at %d pixel pairs", symErr)
	}
	// Pixel count near the circumference 2πr.
	circ := 2 * math.Pi * rad
	if drawn < int(circ*0.85) || drawn > int(circ*1.35) {
		t.Fatalf("drew %d pixels, want ≈ %.0f (2πr)", drawn, circ)
	}

	okPrint("MidpointCircle %d pixels (2πr ≈ %.0f), accuracy and symmetry ok", drawn, 2*math.Pi*rad)

	writeArtifact(t, bitmap.SideBySide(floodFillDemo(), img), "circle_before_after")
}

func TestFloydSteinbergDither(t *testing.T) {
	src := gradient(96, 64)
	dit := FloydSteinbergDither(src)
	if dit == nil {
		t.Fatal("FloydSteinbergDither returned nil")
	}

	// Only {0, 255} per channel.
	for y := range dit.H {
		for x := range dit.W {
			r, g, b := dit.At(x, y)
			for _, c := range [3]uint8{r, g, b} {
				if c != 0 && c != 255 {
					t.Fatalf("pixel (%d,%d) has channel value %d, want 0 or 255", x, y, c)
				}
			}
		}
	}
	okPrint("FloydSteinberg output uses only black and white")

	// 8x8 block means track the source within a tolerance — the
	// defining property of dithering.
	for by := 0; by < dit.H; by += 8 {
		for bx := 0; bx < dit.W; bx += 8 {
			var sSum, dSum float64
			for y := by; y < by+8 && y < dit.H; y++ {
				for x := bx; x < bx+8 && x < dit.W; x++ {
					r, g, b := src.At(x, y)
					sSum += (float64(r) + float64(g) + float64(b)) / 3
					r, g, b = dit.At(x, y)
					dSum += (float64(r) + float64(g) + float64(b)) / 3
				}
			}
			n := 64.0
			if math.Abs(sSum-dSum)/n > 12 {
				t.Fatalf("block (%d,%d) mean drift %.1f too large", bx, by, (sSum-dSum)/n)
			}
		}
	}
	okPrint("FloydSteinberg 8x8 block means preserved within ±12")

	// Solid images stay solid.
	black := bitmap.NewRGB(16, 16)
	white := bitmap.NewRGB(16, 16)
	white.Fill(255, 255, 255)
	checkEqual(t, "dither black", FloydSteinbergDither(black), black)
	checkEqual(t, "dither white", FloydSteinbergDither(white), white)
	okPrint("FloydSteinberg solid images stay solid")

	writeArtifact(t, bitmap.SideBySide(src, dit), "dither_before_after")
}
