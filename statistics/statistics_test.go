package statistics

import (
	"encoding/json"
	"fmt"
	"math"
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

func near(a, b, rtol float64) bool {
	return math.Abs(a-b) <= rtol*math.Max(math.Abs(a), math.Abs(b))
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestDiscreteExpectation(t *testing.T) {
	// A fair die: expectation 3.5.
	fair := []float64{1, 2, 3, 4, 5, 6}
	if v, ok := DiscreteExpectation(fair, []float64{1. / 6, 1. / 6, 1. / 6, 1. / 6, 1. / 6, 1. / 6}); !ok || !near(v, 3.5, 1e-12) {
		t.Errorf("DiscreteExpectation(die) = (%v, %v), want (3.5, true)", v, ok)
	} else {
		okPrint("DiscreteExpectation fair die/%v/%v", v, 3.5)
	}
	// Invalid distributions must be refused, not summed.
	bad := [][]float64{
		{0.5, 0.6},       // does not sum to one
		{0.5, -0.1, 0.6}, // negative probability
		{1.5, 0.5},       // probability above one
	}
	for _, p := range bad {
		if _, ok := DiscreteExpectation(fair[:len(p)], p); ok {
			t.Errorf("DiscreteExpectation(probs=%v) reported success, want failure", p)
		} else {
			okPrint("DiscreteExpectation probs %v (invalid)/rejected/rejected", p)
		}
	}
	if _, ok := DiscreteExpectation(nil, nil); ok {
		t.Error("DiscreteExpectation(empty) reported success, want failure")
	} else {
		okPrint("DiscreteExpectation empty input/rejected/rejected")
	}
	if _, ok := DiscreteExpectation(fair[:3], []float64{1. / 3, 1. / 3, 1. / 3, 0}); ok {
		t.Error("DiscreteExpectation(length mismatch) reported success, want failure")
	} else {
		okPrint("DiscreteExpectation length mismatch/rejected/rejected")
	}
	// No complexity test: O(n), no shape dependence.
}

func TestBayesPosterior(t *testing.T) {
	// Two hypotheses: fair coin (p=1/2) and two-headed (p=1). Observing
	// heads should favor the two-headed coin with weight 2/3.
	prior := []float64{0.5, 0.5}
	like := [][]float64{{0.5, 0.5}, {1, 0}} // rows: hypotheses; cols: outcomes heads, tails
	post, ok := BayesPosterior(prior, like, 0)
	if !ok {
		t.Fatal("BayesPosterior(heads) reported failure")
	}
	if !near(post[1], 2.0/3.0, 1e-12) || !near(post[0], 1.0/3.0, 1e-12) {
		t.Errorf("BayesPosterior(heads) = %v, want [1/3, 2/3]", post)
	} else {
		okPrint("BayesPosterior heads/[1/3 2/3] as %v/[1/3 2/3]", post)
	}
	// Observing tails rules out the two-headed coin entirely: the zero
	// likelihood must zero out its posterior weight without breaking
	// normalization.
	post, ok = BayesPosterior(prior, like, 1)
	if !ok {
		t.Fatal("BayesPosterior(tails) reported failure")
	}
	if !near(post[0], 1.0, 1e-12) {
		t.Errorf("BayesPosterior(tails) = %v, want [1, 0]", post)
	} else {
		okPrint("BayesPosterior tails/%v/[1 0]", post)
	}
	// An observation impossible under every hypothesis must fail cleanly.
	if _, ok := BayesPosterior(prior, [][]float64{{0, 1}, {0, 1}}, 0); ok {
		t.Error("BayesPosterior(impossible obs) reported success, want failure")
	} else {
		okPrint("BayesPosterior impossible observation/rejected/rejected")
	}
	// Inputs must not be mutated.
	like[0][0] = 0.5
	post2, _ := BayesPosterior(prior, like, 0)
	_ = post2
	if like[0][0] != 0.5 || prior[0] != 0.5 {
		t.Error("BayesPosterior mutated its inputs")
	} else {
		okPrint("BayesPosterior input mutation check/got=unmutated/want=unmutated")
	}
	// No complexity test: O(hypotheses x outcomes), no shape dependence.
}

func TestBinomialPMF(t *testing.T) {
	var fx struct {
		Cases []struct {
			N   int       `json:"n"`
			P   float64   `json:"p"`
			PMF []float64 `json:"pmf"`
		} `json:"cases"`
	}
	loadFixture(t, "binom.json", &fx)

	for _, c := range fx.Cases {
		// Whole distribution: every point probability matches the exact
		// integer-arithmetic reference.
		for k, want := range c.PMF {
			got, ok := BinomialPMF(c.N, k, c.P)
			if !ok {
				t.Fatalf("BinomialPMF(%d, %d, %v) reported failure", c.N, k, c.P)
			}
			if !near(got, want, 1e-12) {
				t.Errorf("BinomialPMF(%d, %d, %v) = %v, want %v", c.N, k, c.P, got, want)
			}
		}
		// Point probabilities sum to one.
		var total float64
		for k := range c.PMF {
			v, _ := BinomialPMF(c.N, k, c.P)
			total += v
		}
		if !near(total, 1.0, 1e-12) {
			t.Errorf("BinomialPMF(%d, %v) distribution sums to %v, want 1", c.N, c.P, total)
		} else {
			okPrint("BinomialPMF n=%d p=%v %d points/got=dist matches ref, sums to %v/want=matches, sums to 1",
				c.N, c.P, len(c.PMF), total)
		}
	}
	// Out-of-range k is zero, not an error.
	if v, ok := BinomialPMF(10, 11, 0.5); !ok || v != 0 {
		t.Errorf("BinomialPMF(10, 11, 0.5) = (%v, %v), want (0, true)", v, ok)
	} else {
		okPrint("BinomialPMF out-of-range k/(%v, %v)/(0, true)", v, ok)
	}
	if _, ok := BinomialPMF(10, 5, 1.5); ok {
		t.Error("BinomialPMF with p=1.5 reported success, want failure")
	} else {
		okPrint("BinomialPMF p=1.5 (invalid)/rejected/rejected")
	}
	// No complexity test: O(min(k, n-k)), no shape dependence.
}

func TestCumulativeDistribution(t *testing.T) {
	post, ok := CumulativeDistribution([]float64{0.2, 0.3, 0.1, 0.4})
	if !ok {
		t.Fatal("CumulativeDistribution reported failure")
	}
	want := []float64{0.2, 0.5, 0.6, 1.0}
	for i := range want {
		if !near(post[i], want[i], 1e-12) {
			t.Errorf("CumulativeDistribution[%d] = %v, want %v", i, post[i], want[i])
		} else {
			okPrint("CumulativeDistribution [0.2 0.3 0.1 0.4][%d]/%v/%v", i, post[i], want[i])
		}
	}
	// Structural invariants: non-decreasing, ends at exactly one.
	for i := 1; i < len(post); i++ {
		if post[i] < post[i-1] {
			t.Errorf("CumulativeDistribution decreases at %d: %v < %v", i, post[i], post[i-1])
		}
	}
	if post[len(post)-1] != 1.0 {
		t.Errorf("CumulativeDistribution last entry = %v, want exactly 1", post[len(post)-1])
	}
	if _, ok := CumulativeDistribution([]float64{0.5, -0.1, 0.6}); ok {
		t.Error("CumulativeDistribution with negative probability reported success, want failure")
	} else {
		okPrint("CumulativeDistribution negative probability/rejected/rejected")
	}
	// No complexity test: O(n), no shape dependence.
}

func TestKDE1D(t *testing.T) {
	var fx struct {
		Xs   []float64 `json:"xs"`
		H    float64   `json:"h"`
		Grid []float64 `json:"grid"`
		Dens []float64 `json:"dens"`
	}
	loadFixture(t, "kde.json", &fx)

	got := KDE1D(fx.Xs, fx.H, fx.Grid)
	if len(got) != len(fx.Grid) {
		t.Fatalf("KDE1D returned %d values, want %d", len(got), len(fx.Grid))
	}
	maxRel := 0.0
	for i := range got {
		if !near(got[i], fx.Dens[i], 1e-9) {
			t.Errorf("KDE1D(grid[%d] = %v) = %v, want %v", i, fx.Grid[i], got[i], fx.Dens[i])
		}
		if fx.Dens[i] != 0 {
			maxRel = math.Max(maxRel, math.Abs(got[i]-fx.Dens[i])/fx.Dens[i])
		}
	}
	okPrint("KDE1D %d grid points/got max rel err %.2e/want max rel err <=1e-9", len(got), maxRel)
	// The estimate must integrate (roughly, by the trapezoid rule over the
	// dense grid) to one — a check an independent failure mode, e.g. a
	// missing normalization, cannot pass.
	var area float64
	for i := 1; i < len(got); i++ {
		area += (got[i] + got[i-1]) / 2 * (fx.Grid[i] - fx.Grid[i-1])
	}
	if !near(area, 1.0, 0.01) {
		t.Errorf("KDE1D integrates to %v over the fixture grid, want ~1", area)
	} else {
		okPrint("KDE1D trapezoid integral over grid/%.4f/~1", area)
	}
	// Complexity: cost grows like (points times queries).
	p := complexity.FitExponent(40, 3, func(n int) {
		xs := make([]float64, n)
		for i := range xs {
			xs[i] = float64(i%17) - 8
		}
		KDE1D(xs, 0.5, xs)
	})
	if p > 2.4 {
		t.Errorf("KDE1D measured exponent %.2f, expected ~2", p)
	}
}

func TestEMBinomial(t *testing.T) {
	var raw struct {
		Series [][2]int     `json:"series"`
		Init   []float64    `json:"init"`
		Theta1 float64      `json:"theta1"`
		Theta2 float64      `json:"theta2"`
		Trace  [][3]float64 `json:"trace"`
	}
	loadFixture(t, "em_coin.json", &raw)
	var series []BinomialSeries
	for _, s := range raw.Series {
		series = append(series, BinomialSeries{Heads: s[0], Tails: s[1]})
	}

	t1, t2, ll, ok := EMBinomial(series, raw.Init[0], raw.Init[1], 200)
	if !ok {
		t.Fatal("EMBinomial reported failure")
	}
	if !near(t1, raw.Theta1, 1e-4) || !near(t2, raw.Theta2, 1e-4) {
		t.Errorf("EMBinomial converged to (%v, %v), want (%v, %v)",
			t1, t2, raw.Theta1, raw.Theta2)
	} else {
		okPrint("EMBinomial coin fixture/(%v, %v)/(%v, %v)", t1, t2, raw.Theta1, raw.Theta2)
	}
	// Monotonicity: the log-likelihood never decreases across iterations.
	for i := 1; i < len(ll); i++ {
		if ll[i] < ll[i-1]-1e-9 {
			t.Fatalf("EMBinomial log-likelihood decreased at iteration %d: %v < %v",
				i, ll[i], ll[i-1])
		}
	}
	// Trace head matches the independent implementation.
	for i, row := range raw.Trace {
		if i >= len(ll) {
			break
		}
		if !near(ll[i], row[2], 1e-8) {
			t.Errorf("EMBinomial log-likelihood[%d] = %v, want %v", i, ll[i], row[2])
		}
	}
	// Invalid initial probabilities are refused.
	if _, _, _, ok := EMBinomial(series, 1.5, 0.5, 10); ok {
		t.Error("EMBinomial with theta1=1.5 reported success, want failure")
	} else {
		okPrint("EMBinomial theta1=1.5 (invalid)/rejected/rejected")
	}
	// No complexity test: O(series) per iteration, no shape dependence.
}

func TestEMMultinomial(t *testing.T) {
	var fx struct {
		Obs       [][]float64  `json:"obs"`
		InitW     []float64    `json:"initW"`
		InitTheta [][]float64  `json:"initTheta"`
		W         []float64    `json:"w"`
		Theta     [][]float64  `json:"theta"`
		TraceHead [][5]float64 `json:"traceHead"`
	}
	loadFixture(t, "em_multinom.json", &fx)

	w, theta, ll, ok := EMMultinomial(fx.Obs, fx.InitW, fx.InitTheta, 300)
	if !ok {
		t.Fatal("EMMultinomial reported failure")
	}
	for i := range fx.W {
		if !near(w[i], fx.W[i], 1e-4) {
			t.Errorf("EMMultinomial w[%d] = %v, want %v", i, w[i], fx.W[i])
		}
		for j := range fx.W {
			if !near(theta[i][j], fx.Theta[i][j], 1e-4) {
				t.Errorf("EMMultinomial theta[%d][%d] = %v, want %v",
					i, j, theta[i][j], fx.Theta[i][j])
			}
		}
	}
	okPrint("EMMultinomial fixture %d components/got=(w, theta) match fixture/want=match fixture", len(fx.W))
	// Component probabilities stay normalized.
	for i := range theta {
		var s float64
		for j := range theta[i] {
			s += theta[i][j]
		}
		if !near(s, 1.0, 1e-9) {
			t.Errorf("EMMultinomial theta[%d] sums to %v, want 1", i, s)
		}
	}
	// Monotonic log-likelihood.
	for i := 1; i < len(ll); i++ {
		if ll[i] < ll[i-1]-1e-9 {
			t.Fatalf("EMMultinomial log-likelihood decreased at iteration %d: %v < %v",
				i, ll[i], ll[i-1])
		}
	}
	// Trace head matches the independent implementation.
	for i, row := range fx.TraceHead {
		if i >= len(ll) {
			break
		}
		if !near(ll[i], row[len(row)-1], 1e-8) {
			t.Errorf("EMMultinomial log-likelihood[%d] = %v, want %v", i, ll[i], row[len(row)-1])
		}
	}
	// No complexity test: O(observations x faces) per iteration.
}
