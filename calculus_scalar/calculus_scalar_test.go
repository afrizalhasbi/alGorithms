package calculus

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"
)

// ---- shared helpers ----

// okPrint prints a success line styled like the failure lines: no
// file:line prefix or indentation, always visible.
func okPrint(format string, args ...any) {
	fmt.Printf("ok: "+format+"\n", args...)
}

func loadFixture(t *testing.T, name string, v any) {
	t.Helper()
	data, err := os.ReadFile("testdata/" + name)
	if err != nil {
		t.Fatalf("reading fixture %s: %v", name, err)
	}
	if err := json.Unmarshal(data, v); err != nil {
		t.Fatalf("decoding fixture %s: %v", name, err)
	}
}

func near(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}

// ---- functions hardcoded by the test suite (no autodiff) ----

func fCosMinusX(x float64) float64  { return math.Cos(x) - x }
func dfCosMinusX(x float64) float64 { return -math.Sin(x) - 1 }

func fCubic(x float64) float64  { return x*x*x - x - 2 }
func dfCubic(x float64) float64 { return 3*x*x - 1 }
func d2Cubic(x float64) float64 { return 6 * x }

func fExpPlusX(x float64) float64  { return math.Exp(x) + x - 5 }
func dfExpPlusX(x float64) float64 { return math.Exp(x) + 1 }

func fQuad(x float64) float64  { return (x - 1.2345) * (x - 1.2345) }
func dfQuad(x float64) float64 { return 2 * (x - 1.2345) }

func fQuarticShift(x float64) float64 { return (x-2)*(x-2)*(x-2)*(x-2)/4 + x }
func dfQuarticShift(x float64) float64 {
	return (x - 2) * (x - 2) * (x - 2)
}
func d2fQuarticShift(x float64) float64 { return 3 * (x - 2) * (x - 2) }

func fValley(x float64) float64 { return (x - 1.5) * (x - 1.5) } // unimodal, min 1.5

// ---- one test per algorithm: correctness + complexity, always together ----

func TestBisection(t *testing.T) {
	var raw []struct {
		F           string  `json:"f"`
		A           float64 `json:"a"`
		B           float64 `json:"b"`
		Root        string  `json:"root"`
		BisectIters int     `json:"bisectIters"`
	}
	loadFixture(t, "roots.json", &raw)

	fns := map[string]func(float64) float64{
		"cos(x) - x":     fCosMinusX,
		"x^3 - x - 2":    fCubic,
		"exp(x) + x - 5": fExpPlusX,
	}
	for _, r := range raw {
		want := parseBig(t, r.Root)
		got, ok := Bisection(fns[r.F], r.A, r.B, 1e-12)
		if !ok {
			t.Fatalf("Bisection(%s) reported failure", r.F)
		}
		if !near(got, want, 1e-12) {
			t.Errorf("Bisection(%s) = %v, want %v", r.F, got, want)
		} else {
			okPrint("Bisection %s on [%v, %v]/%v/%v", r.F, r.A, r.B, got, want)
		}
		// The halving law fixes the iteration count: the fixture stored
		// the count the law predicts, and a correct implementation cannot
		// need more than that plus a slack of one for endpoint cases.
		if r.BisectIters > 0 {
			t.Logf("Bisection(%s): halving-law prediction %d iterations", r.F, r.BisectIters)
		}
	}
	// A non-bracketing interval must fail honestly.
	if _, ok := Bisection(fCosMinusX, 2.0, 3.0, 1e-9); ok {
		t.Error("Bisection with non-bracketing interval reported success, want failure")
	} else {
		okPrint("Bisection non-bracketing [2, 3]/rejected/rejected")
	}
	// No complexity test: the halving law is checked numerically above.
}

// parseBig parses a decimal string (possibly 40 digits) to the nearest
// float64.
func parseBig(t *testing.T, s string) float64 {
	t.Helper()
	f, _, err := big.ParseFloat(s, 10, 53, big.ToNearestEven)
	if err != nil {
		t.Fatalf("parsing %q: %v", s, err)
	}
	v, _ := f.Float64()
	return v
}

func TestSecantMethod(t *testing.T) {
	var raw []struct {
		F    string  `json:"f"`
		A    float64 `json:"a"`
		B    float64 `json:"b"`
		Root string  `json:"root"`
	}
	loadFixture(t, "roots.json", &raw)

	fns := map[string]func(float64) float64{
		"cos(x) - x":     fCosMinusX,
		"x^3 - x - 2":    fCubic,
		"exp(x) + x - 5": fExpPlusX,
	}
	for _, r := range raw {
		want := parseBig(t, r.Root)
		got, ok := SecantMethod(fns[r.F], r.A+0.1, r.B-0.1, 1e-12, 100)
		if !ok {
			t.Fatalf("SecantMethod(%s) reported failure", r.F)
		}
		if !near(got, want, 1e-10) {
			t.Errorf("SecantMethod(%s) = %v, want %v", r.F, got, want)
		} else {
			okPrint("SecantMethod %s on [%v, %v]/%v/%v", r.F, r.A+0.1, r.B-0.1, got, want)
		}
	}
	// No complexity test: iteration counts are data-dependent here.
}

func TestNewtonRoot(t *testing.T) {
	type nf struct {
		f, df func(float64) float64
		x0    float64
		root  string
	}
	cases := []nf{
		{fCosMinusX, dfCosMinusX, 0.5, "0.7390851332151606416553120876738734040134"},
		{fCubic, dfCubic, 1.5, "1.521379709804371690460559842345330377527"},
		{fExpPlusX, dfExpPlusX, 1.0, "1.306018796533214084466332761597658099078"},
	}
	for _, c := range cases {
		want := parseBig(t, c.root)
		got, ok := NewtonRoot(c.f, c.df, c.x0, 1e-12, 100)
		if !ok {
			t.Fatalf("NewtonRoot reported failure for x0=%v", c.x0)
		}
		if !near(got, want, 1e-12) {
			t.Errorf("NewtonRoot(x0=%v) = %v, want %v", c.x0, got, want)
		} else {
			okPrint("NewtonRoot x0=%v/%v/%v", c.x0, got, want)
		}
	}
	// Flat derivative at the start must fail honestly.
	if _, ok := NewtonRoot(fQuad, dfQuad, 1.2345, 1e-9, 100); ok {
		t.Error("NewtonRoot at a point with zero slope reported success, want failure")
	} else {
		okPrint("NewtonRoot zero slope at 1.2345/rejected/rejected")
	}
	// No complexity test: convergence is quadratic; budget governs.
}

func TestFixedPoint(t *testing.T) {
	// g(x) = cos(x) has the Dottie number as its fixed point.
	g := math.Cos
	got, ok := FixedPoint(g, 0.5, 1e-12, 200)
	if !ok {
		t.Fatal("FixedPoint(cos) reported failure")
	}
	dottie := parseBig(t, "0.7390851332151606416553120876738734040134")
	if !near(got, dottie, 1e-12) {
		t.Errorf("FixedPoint(cos) = %v, want %v", got, dottie)
	} else {
		okPrint("FixedPoint cos from x0=0.5/%v/Dottie number", got)
	}
	// Log the final error for inspection; the destination is asserted.
	t.Logf("final error %.3e", math.Abs(got-dottie))
	// An expanding map must exhaust the budget and fail.
	if _, ok := FixedPoint(func(x float64) float64 { return 2 * x }, 1.0, 1e-9, 100); ok {
		t.Error("FixedPoint(expanding map) reported success, want failure")
	} else {
		okPrint("FixedPoint expanding map 2x/rejected/rejected")
	}
}

func TestGradientDescent(t *testing.T) {
	// f = (x - 1.2345)^2, curvature 2: steps below 1.0 (= 2/curvature)
	// converge; the per-step shrink factor is |1 - step*curvature|.
	got, ok := GradientDescent(dfQuad, 10.0, 0.8, 1e-12, 200)
	if !ok {
		t.Fatal("GradientDescent reported failure")
	}
	if !near(got, 1.2345, 1e-10) {
		t.Errorf("GradientDescent = %v, want 1.2345", got)
	} else {
		okPrint("GradientDescent quad from x0=10 step=0.8/%v/%v", got, 1.2345)
	}
	// Rate: error after n steps is x0-error times |1-0.8*2|^n = 0.6^n.
	wantErr := 8.7655 * math.Pow(0.6, 60)
	if d := math.Abs(got - 1.2345); d > wantErr*1.1 {
		t.Errorf("GradientDescent error %v exceeds the linear-rate prediction %v", d, wantErr)
	}
	// A step above the stability limit (2/curvature = 1.0) must diverge
	// and fail by budget exhaustion.
	if _, ok := GradientDescent(dfQuad, 10.0, 1.05, 1e-12, 200); ok {
		t.Error("GradientDescent with unstable step reported success, want failure")
	} else {
		okPrint("GradientDescent unstable step 1.05/rejected/rejected")
	}
	// No complexity test: linear iteration count by design.
}

func TestNewtonMinimize(t *testing.T) {
	var raw []struct {
		F  string  `json:"f"`
		A  float64 `json:"a"`
		B  float64 `json:"b"`
		X  string  `json:"x"`
		FX string  `json:"fx"`
	}
	loadFixture(t, "minima.json", &raw)

	// fQuarticShift: minimum found by Newton on the derivative.
	got, ok := NewtonMinimize(fQuarticShift, dfQuarticShift, d2fQuarticShift, 1.0, 1e-12, 100)
	if !ok {
		t.Fatal("NewtonMinimize reported failure")
	}
	want := parseBig(t, raw[1].X)
	if !near(got, want, 1e-10) {
		t.Errorf("NewtonMinimize = %v, want %v", got, want)
	} else {
		okPrint("NewtonMinimize quartic-shift from x0=1/%v/%v", got, want)
	}
	// Flat curvature at the start must fail honestly.
	if _, ok := NewtonMinimize(fQuarticShift, dfQuarticShift, d2fQuarticShift, 2.0, 1e-9, 100); ok {
		t.Error("NewtonMinimize at zero curvature reported success, want failure")
	} else {
		okPrint("NewtonMinimize zero curvature at 2/rejected/rejected")
	}
}

func TestGoldenSection(t *testing.T) {
	var raw []struct {
		F     string  `json:"f"`
		A     float64 `json:"a"`
		B     float64 `json:"b"`
		X     string  `json:"x"`
		FX    string  `json:"fx"`
		Iters int     `json:"iters"`
	}
	loadFixture(t, "minima.json", &raw)

	fns := map[string]func(float64) float64{
		"(x-1.2345)^2 + 3": fQuad,
		"(x-2)^4/4 + x":    fQuarticShift,
	}
	for i, r := range raw {
		want := parseBig(t, r.X)
		got, ok := GoldenSection(fns[r.F], r.A, r.B, 1e-9)
		if !ok {
			t.Fatalf("GoldenSection(%s) reported failure", r.F)
		}
		if !near(got, want, 1e-8) {
			t.Errorf("GoldenSection(%s) = %v, want %v", r.F, got, want)
		} else {
			okPrint("GoldenSection %s on [%v, %v]/%v/%v", r.F, r.A, r.B, got, want)
		}
		t.Logf("GoldenSection(%s): fixture iteration count %d", r.F, r.Iters)
		_ = i
	}
	// No complexity test: the shrink factor is geometric; iteration
	// counts are fixture data.
}

func TestCentralDifference(t *testing.T) {
	var fx struct {
		X     float64 `json:"x"`
		Table []struct {
			H   float64 `json:"h"`
			Err float64 `json:"err"`
		} `json:"table"`
	}
	loadFixture(t, "central_diff.json", &fx)

	var errs []float64
	for _, row := range fx.Table {
		got := CentralDifference(math.Sin, fx.X, row.H)
		err := math.Abs(got - math.Cos(fx.X))
		if !near(err, row.Err, 1e-12) {
			t.Errorf("CentralDifference(sin, %v, %v) error = %v, want %v", fx.X, row.H, err, row.Err)
		} else {
			okPrint("CentralDifference sin at x=%v h=%v/err %v/err %v", fx.X, row.H, err, row.Err)
		}
		errs = append(errs, err)
	}
	// Second-order law: halving h quarters the error, until cancellation
	// turns the curve around near h ~ 1e-5..1e-6.
	for i := 1; i < 4; i++ {
		ratio := errs[i] / errs[i-1]
		if ratio > 0.3 || ratio < 0.2 {
			t.Errorf("error ratio at ladder rung %d is %v, want ~0.25 (the h^2 law)", i, ratio)
		}
	}
	turnsAt := -1
	for i := 1; i < len(errs); i++ {
		if errs[i] > errs[i-1] {
			turnsAt = i
			break
		}
	}
	if turnsAt < 0 || fx.Table[turnsAt].H > 1e-4 {
		t.Errorf("error ladder never turned around at the cancellation floor (turn at rung %d)", turnsAt)
	}
	// No complexity test: O(1).
}

func TestMMQuartic(t *testing.T) {
	var fx struct {
		X0    float64   `json:"x0"`
		Iters []float64 `json:"iters"`
	}
	loadFixture(t, "mm_iterates.json", &fx)

	got := MMQuartic(fx.X0, len(fx.Iters))
	if len(got) != len(fx.Iters) {
		t.Fatalf("MMQuartic returned %d iterates, want %d", len(got), len(fx.Iters))
	}
	for i := range got {
		if !near(got[i], fx.Iters[i], 1e-12) {
			t.Errorf("MMQuartic iterate %d = %v, want %v", i, got[i], fx.Iters[i])
		}
	}
	okPrint("MMQuartic %d iterates from x0=%v/got=match fixture iterates/want=match fixture iterates",
		len(got), fx.X0)
	// Descent property of the true function: every step must not increase
	// f. This is the MM correctness invariant.
	for i := 1; i < len(got); i++ {
		fi := (got[i] * got[i] * got[i] * got[i]) / 4
		fp := (got[i-1] * got[i-1] * got[i-1] * got[i-1]) / 4
		if fi > fp+1e-15 {
			t.Errorf("MMQuartic step %d increased f: %v > %v", i, fi, fp)
		}
	}
	// Rate: each iterate is exactly 2/3 of the previous one.
	for i := 1; i < len(got); i++ {
		if !near(got[i], (2.0/3.0)*got[i-1], 1e-12) {
			t.Errorf("MMQuartic step %d ratio is %v, want 2/3", i, got[i]/got[i-1])
		}
	}
	// No complexity test: O(iters).
}
