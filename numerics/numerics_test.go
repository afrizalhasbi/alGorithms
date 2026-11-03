package numerics

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"os"
	"testing"

	"alGorithms/internal/complexity"
)

// ---- shared helpers ----

// okPrint prints a success line styled like the failure lines: no
// file:line prefix or indentation, always visible.
func okPrint(format string, args ...any) {
	fmt.Printf("ok: "+format+"\n", args...)
}

// loadFixture decodes a testdata file. go test runs with the package
// directory as working directory, so "testdata/x.json" resolves correctly.
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

// exactF64 parses a decimal string holding an exact rational sum into the
// nearest float64 (via big.Float, so no reference rounding is introduced
// here either).
func exactF64(t *testing.T, s string) float64 {
	t.Helper()
	f, ok := new(big.Float).SetPrec(200).SetString(s)
	if !ok {
		t.Fatalf("bad exact decimal %q", s)
	}
	v, _ := f.Float64()
	return v
}

func relErr(got, want float64) float64 {
	return math.Abs(got-want) / math.Abs(want)
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestKahanSum(t *testing.T) {
	var fx struct {
		Cases []struct {
			Name     string    `json:"name"`
			Xs       []float64 `json:"xs"`
			Exact    string    `json:"exact"`
			ExactF64 float64   `json:"exactF64"`
			NaiveF64 float64   `json:"naiveF64"`
		} `json:"cases"`
		Repeat struct {
			X     float64 `json:"x"`
			N     int     `json:"n"`
			Exact string  `json:"exact"`
		} `json:"repeat"`
	}
	loadFixture(t, "sums.json", &fx)

	for _, c := range fx.Cases {
		want := exactF64(t, c.Exact)
		got := KahanSum(c.Xs)
		if !NearEqual(got, want, 1e-15, 1e-300) {
			t.Errorf("KahanSum(%s) = %v, want %v", c.Name, got, want)
		} else {
			okPrint("KahanSum %s/%v/%v", c.Name, got, want)
		}
		// The test documents why the algorithm exists: on the adversarial
		// fixture the naive sum must be strictly wrong.
		if c.Name == "adversarial" && c.NaiveF64 == want {
			t.Errorf("fixture drift: naive summation is no longer wrong on %s", c.Name)
		}
	}

	// Repeat case: n copies of a small value vs the exact rational total.
	rep := make([]float64, fx.Repeat.N)
	for i := range rep {
		rep[i] = fx.Repeat.X
	}
	want := exactF64(t, fx.Repeat.Exact)
	if got := KahanSum(rep); !NearEqual(got, want, 1e-13, 1e-300) {
		t.Errorf("KahanSum(%d copies of %v) = %v, want %v", fx.Repeat.N, fx.Repeat.X, got, want)
	} else {
		okPrint("KahanSum %d copies of %v/%v/%v", fx.Repeat.N, fx.Repeat.X, got, want)
	}

	// Complexity: O(n). Sizes are modest so the test stays quick.
	p := complexity.FitExponent(20000, 3, func(n int) {
		xs := make([]float64, n)
		for i := range xs {
			xs[i] = 0.1
		}
		KahanSum(xs)
	})
	if p > 1.3 {
		t.Errorf("KahanSum measured exponent %.2f, expected ~1", p)
	}
}

func TestNeumaierSum(t *testing.T) {
	var fx struct {
		Xs       []float64 `json:"xs"`
		ExactF64 float64   `json:"exactF64"`
	}
	loadFixture(t, "sums_desc.json", &fx)

	got := NeumaierSum(fx.Xs)
	if !NearEqual(got, fx.ExactF64, 1e-15, 1e-300) {
		t.Errorf("NeumaierSum(descending) = %v, want %v", got, fx.ExactF64)
	} else {
		okPrint("NeumaierSum descending fixture/%v/%v", got, fx.ExactF64)
	}
	// Same fixture through Kahan: it may or may not match; log, don't fail.
	// (The descending order is exactly where the two algorithms differ.)
	t.Logf("KahanSum on same fixture = %v (exact %v)", KahanSum(fx.Xs), fx.ExactF64)

	// Adversarial case shared with Kahan.
	adv := []float64{1e16, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, -1e16}
	if got := NeumaierSum(adv); got != 20 {
		t.Errorf("NeumaierSum(adversarial) = %v, want 20", got)
	} else {
		okPrint("NeumaierSum adversarial [1e16 ... -1e16]/%v/%d", got, 20)
	}
}

func TestNearEqual(t *testing.T) {
	cases := []struct {
		a, b, rtol, floor float64
		want              bool
	}{
		{1.0, 1.0 + 1e-12, 1e-9, 1e-300, true},
		{1.0, 1.0 + 1e-6, 1e-9, 1e-300, false},
		// Absolute tolerance fails here, relative works: 1e6 apart by 1e-6.
		{1e6, 1e6 + 1e-6, 1e-9, 1e-300, false},
		{1e6, 1e6 + 1, 1e-9, 1e-300, true},
		// Near zero the relative test is meaningless; the floor catches it.
		{1e-12, 2e-12, 1e-9, 1e-9, true},
		{1e-12, 2e-12, 1e-9, 1e-300, false},
		{math.Inf(1), math.Inf(1), 1e-9, 0, true},
		{math.Inf(1), math.Inf(-1), 1e-9, 0, false},
		{math.NaN(), math.NaN(), 1e-9, 0, false},
		{math.NaN(), 1.0, 1e-9, 0, false},
	}
	for _, c := range cases {
		if got := NearEqual(c.a, c.b, c.rtol, c.floor); got != c.want {
			t.Errorf("NearEqual(%v, %v, %v, %v) = %v, want %v",
				c.a, c.b, c.rtol, c.floor, got, c.want)
		} else {
			okPrint("NearEqual %v %v rtol %v floor %v/%v/%v", c.a, c.b, c.rtol, c.floor, got, c.want)
		}
	}
	// No complexity test: O(1).
}

func TestWelford(t *testing.T) {
	var fx struct {
		Xs   []float64 `json:"xs"`
		Mean float64   `json:"mean"`
		VarP float64   `json:"varP"`
		VarS float64   `json:"varS"`
	}
	loadFixture(t, "welford.json", &fx)

	mean, varS, varP := Welford(fx.Xs)
	if !NearEqual(mean, fx.Mean, 1e-12, 1e-300) {
		t.Errorf("Welford mean = %v, want %v", mean, fx.Mean)
	} else {
		okPrint("Welford fixture mean/%v/%v", mean, fx.Mean)
	}
	if !NearEqual(varS, fx.VarS, 1e-12, 1e-300) {
		t.Errorf("Welford sample variance = %v, want %v", varS, fx.VarS)
	} else {
		okPrint("Welford fixture sample variance/%v/%v", varS, fx.VarS)
	}
	if !NearEqual(varP, fx.VarP, 1e-12, 1e-300) {
		t.Errorf("Welford population variance = %v, want %v", varP, fx.VarP)
	} else {
		okPrint("Welford fixture population variance/%v/%v", varP, fx.VarP)
	}

	// Catastrophic cancellation case: the naive sum-of-squares formula
	// returns garbage; Welford must return exactly the small variance.
	canc := []float64{1e9, 1e9 + 1, 1e9 + 2}
	m2, vs2, vp2 := Welford(canc)
	if !NearEqual(m2, 1e9+1, 1e-15, 1e-300) {
		t.Errorf("Welford(cancellation) mean = %v, want %v", m2, 1e9+1)
	} else {
		okPrint("Welford cancellation mean/%v/%v", m2, 1e9+1)
	}
	if !NearEqual(vs2, 1.0, 1e-12, 1e-300) || !NearEqual(vp2, 2.0/3.0, 1e-12, 1e-300) {
		t.Errorf("Welford(cancellation) variances = (%v, %v), want (1, 2/3)", vs2, vp2)
	} else {
		okPrint("Welford cancellation variances/(%v, %v)/(1, 2/3)", vs2, vp2)
	}

	// Complexity: O(n), single pass.
	p := complexity.FitExponent(20000, 3, func(n int) {
		xs := make([]float64, n)
		for i := range xs {
			xs[i] = float64(i)
		}
		Welford(xs)
	})
	if p > 1.3 {
		t.Errorf("Welford measured exponent %.2f, expected ~1", p)
	}
}

func TestHorner(t *testing.T) {
	var fx struct {
		Polys []struct {
			Coeffs []float64 `json:"coeffs"`
			Xs     []float64 `json:"xs"`
			Refs   []float64 `json:"refs"`
		} `json:"polys"`
	}
	loadFixture(t, "horner.json", &fx)

	for _, poly := range fx.Polys {
		for i, x := range poly.Xs {
			want := poly.Refs[i]
			if got := Horner(poly.Coeffs, x); !NearEqual(got, want, 1e-12, 1e-300) {
				t.Errorf("Horner(%v, %v) = %v, want %v", poly.Coeffs, x, got, want)
			} else {
				okPrint("Horner poly with %d coeffs at x=%v/%v/%v", len(poly.Coeffs), x, got, want)
			}
		}
	}
	if Horner(nil, 3.0) != 0 {
		t.Error("Horner(empty) = nonzero, want 0")
	} else {
		okPrint("Horner empty coeffs at x=3/%v/%v", 0.0, 0.0)
	}
	if got := Horner([]float64{7.5}, 123.0); got != 7.5 {
		t.Errorf("Horner([7.5], 123) = %v, want 7.5", got)
	} else {
		okPrint("Horner [7.5] at x=123/%v/%v", got, 7.5)
	}
	// No complexity test: O(n), no shape dependence.
}

func TestLogSumExp(t *testing.T) {
	var fx struct {
		Cases []struct {
			Xs  []float64 `json:"xs"`
			Ref float64   `json:"ref"`
		} `json:"cases"`
	}
	loadFixture(t, "logsumexp.json", &fx)

	for i, c := range fx.Cases {
		// First demonstrate the failure the function exists to prevent:
		// the naive sum of exponentials must not be finite-and-correct.
		var naiveSum float64
		for _, x := range c.Xs {
			naiveSum += math.Exp(x)
		}
		t.Logf("case %d: naive sum of exps = %v (0 or +Inf is expected)", i, naiveSum)

		got := LogSumExp(c.Xs)
		if !NearEqual(got, c.Ref, 1e-12, 1e-300) {
			t.Errorf("LogSumExp(case %d) = %v, want %v", i, got, c.Ref)
		} else {
			okPrint("LogSumExp case %d (%d values)/%v/%v", i, len(c.Xs), got, c.Ref)
		}
	}
	// No complexity test: O(n), no shape dependence.
}

func TestTaylorExp(t *testing.T) {
	for _, x := range []float64{0, 0.5, 1, 2, -1, -2, 5, -5} {
		want := math.Exp(x)
		got := TaylorExp(x, 60)
		if !NearEqual(got, want, 1e-12, 1e-300) {
			t.Errorf("TaylorExp(%v) = %v, want %v", x, got, want)
		} else {
			okPrint("TaylorExp %v (60 terms)/%v/%v", x, got, want)
		}
	}
	// Alternating-series case: negative x must be just as accurate as
	// positive x — this fails under naive left-to-right summation.
	x := -8.0
	if relErr(TaylorExp(x, 60), math.Exp(x)) > 1e-12 {
		t.Errorf("TaylorExp(%v) relative error too large for an alternating series", x)
	} else {
		okPrint("TaylorExp %v alternating series/got rel err %.2e/want rel err <=1e-12", x, relErr(TaylorExp(x, 60), math.Exp(x)))
	}
	// No complexity test: loop count is bounded by the term budget.
}

func TestSafeLogSafeExp(t *testing.T) {
	if v, ok := SafeLog(2.718281828459045); !ok || !NearEqual(v, 1.0, 1e-15, 1e-300) {
		t.Errorf("SafeLog(e) = (%v, %v), want (1, true)", v, ok)
	} else {
		okPrint("SafeLog e/%v/%v", v, 1.0)
	}
	for _, x := range []float64{0, -1, math.NaN(), math.Inf(-1)} {
		if _, ok := SafeLog(x); ok {
			t.Errorf("SafeLog(%v) reported success, want failure", x)
		} else {
			okPrint("SafeLog %v (invalid)/rejected/rejected", x)
		}
	}
	if v, ok := SafeExp(0); !ok || v != 1 {
		t.Errorf("SafeExp(0) = (%v, %v), want (1, true)", v, ok)
	} else {
		okPrint("SafeExp 0/%v/%v", v, 1.0)
	}
	if _, ok := SafeExp(1000); ok {
		t.Error("SafeExp(1000) reported success, want overflow failure")
	} else {
		okPrint("SafeExp 1000 (overflow)/rejected/rejected")
	}
	if v, ok := SafeExp(-1); !ok || !NearEqual(v, math.Exp(-1), 1e-15, 1e-300) {
		t.Errorf("SafeExp(-1) = (%v, %v), want (e^-1, true)", v, ok)
	} else {
		okPrint("SafeExp -1/%v/%v", v, math.Exp(-1))
	}
	// No complexity test: O(1).
}
