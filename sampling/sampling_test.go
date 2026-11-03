package sampling

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
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

// checkMoments asserts the sample mean and variance of xs against mu and
// sigma^2, with the tolerance set to k standard errors of the mean.
// Every sampler test uses this so the tolerances are principled, not
// magic numbers.
func checkMoments(t *testing.T, name string, xs []float64, mu, sigma, k float64) {
	t.Helper()
	n := float64(len(xs))
	var sum, sumSq float64
	for _, x := range xs {
		sum += x
		sumSq += x * x
	}
	mean := sum / n
	varP := sumSq/n - mean*mean
	se := sigma / math.Sqrt(n)
	if math.Abs(mean-mu) > k*se {
		t.Errorf("%s: sample mean %v is more than %v standard errors from %v",
			name, mean, k, mu)
	}
	wantVar := sigma * sigma
	varTol := wantVar * k * math.Sqrt(2.0/n)
	if math.Abs(varP-wantVar) > varTol {
		t.Errorf("%s: sample variance %v differs from %v by more than %v standard errors",
			name, varP, wantVar, k)
	}
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestInverseDiscrete(t *testing.T) {
	var fx struct {
		Pmf  []float64 `json:"pmf"`
		Us   []float64 `json:"us"`
		Outs []int     `json:"outs"`
	}
	loadFixture(t, "inverse_discrete.json", &fx)

	got := InverseDiscrete(fx.Pmf, fx.Us)
	if len(got) != len(fx.Outs) {
		t.Fatalf("InverseDiscrete returned %d outcomes, want %d", len(got), len(fx.Outs))
	}
	for i := range got {
		if got[i] != fx.Outs[i] {
			t.Errorf("InverseDiscrete(u=%v) = %d, want %d", fx.Us[i], got[i], fx.Outs[i])
		}
	}
	okPrint("InverseDiscrete %d fixture us/got=fixture outcomes/want=fixture outcomes", len(got))
	// Statistical: 100k samples from a seeded generator must land on each
	// outcome within a chi-square-style band. The band is hand-set wide
	// (5 sigma of the binomial proportion) so the test cannot flake.
	rng := rand.New(rand.NewSource(7))
	pmf := fx.Pmf
	n := 100000
	counts := make([]int, len(pmf))
	us := make([]float64, n)
	for i := range us {
		us[i] = rng.Float64()
	}
	for _, k := range InverseDiscrete(pmf, us) {
		counts[k]++
	}
	for k, c := range counts {
		want := pmf[k] * float64(n)
		se := math.Sqrt(pmf[k] * (1 - pmf[k]) * float64(n))
		if math.Abs(float64(c)-want) > 5*se {
			t.Errorf("outcome %d observed %d times, want ~%v (5 sigma)", k, c, want)
		} else {
			okPrint("InverseDiscrete outcome %d count %d/~%v (5 sigma)", k, c, want)
		}
	}
	// No complexity test: linear, no shape dependence.
}

func TestSampleExponential(t *testing.T) {
	var fx struct {
		Rate float64   `json:"rate"`
		Us   []float64 `json:"us"`
		ExpX []float64 `json:"expX"`
		TriX []float64 `json:"triX"`
	}
	loadFixture(t, "inverse_continuous.json", &fx)

	got := SampleExponential(fx.Rate, fx.Us)
	if len(got) != len(fx.ExpX) {
		t.Fatalf("SampleExponential returned %d samples, want %d", len(got), len(fx.ExpX))
	}
	for i := range got {
		if !near(got[i], fx.ExpX[i], 1e-12) {
			t.Errorf("SampleExponential(u=%v) = %v, want %v", fx.Us[i], got[i], fx.ExpX[i])
		}
	}
	okPrint("SampleExponential %d fixture us rate %v/got=fixture expX/want=fixture expX", len(got), fx.Rate)
	// Statistical: exponential(1.5) has mean 1/rate, variance 1/rate^2.
	rng := rand.New(rand.NewSource(11))
	n := 100000
	us := make([]float64, n)
	for i := range us {
		us[i] = rng.Float64()
	}
	checkMoments(t, "SampleExponential", SampleExponential(1.5, us),
		1.0/1.5, 1.0/1.5, 5)
	okPrint("SampleExponential %d seeded samples/got=mean+variance within 5 SE/want=within 5 SE", n)
	// Invalid inputs must be refused.
	if SampleExponential(-1, []float64{0.5}) != nil {
		t.Error("SampleExponential with negative rate returned samples, want nil")
	} else {
		okPrint("SampleExponential negative rate/rejected/rejected")
	}
	if SampleExponential(1, []float64{1.0}) != nil {
		t.Error("SampleExponential accepted u=1.0, want refusal")
	} else {
		okPrint("SampleExponential u=1.0 (out of domain)/rejected/rejected")
	}
}

func TestSampleTriangular(t *testing.T) {
	var fx struct {
		Rate float64   `json:"rate"`
		Us   []float64 `json:"us"`
		ExpX []float64 `json:"expX"`
		TriX []float64 `json:"triX"`
	}
	loadFixture(t, "inverse_continuous.json", &fx)

	got := SampleTriangular(fx.Us)
	if len(got) != len(fx.TriX) {
		t.Fatalf("SampleTriangular returned %d samples, want %d", len(got), len(fx.TriX))
	}
	for i := range got {
		if !near(got[i], fx.TriX[i], 1e-12) {
			t.Errorf("SampleTriangular(u=%v) = %v, want %v", fx.Us[i], got[i], fx.TriX[i])
		}
	}
	okPrint("SampleTriangular %d fixture us/got=fixture triX/want=fixture triX", len(got))
	// Statistical: triangular on [0,2] with peak at 1 has mean 1 and
	// variance 1/3.
	rng := rand.New(rand.NewSource(13))
	n := 100000
	us := make([]float64, n)
	for i := range us {
		us[i] = rng.Float64()
	}
	checkMoments(t, "SampleTriangular", SampleTriangular(us), 1.0, math.Sqrt(1.0/3.0), 5)
	okPrint("SampleTriangular %d seeded samples/got=mean+variance within 5 SE/want=within 5 SE", n)
}

func TestBoxMuller(t *testing.T) {
	// Exact pair: (0.5, 0.5) is the anchor case — radius 0, both normals 0.
	n1, n2, ok := BoxMuller(0.5, 0.5)
	if !ok {
		t.Fatal("BoxMuller(0.5, 0.5) reported failure")
	}
	if !near(n1, 0, 1e-12) || !near(n2, 0, 1e-12) {
		t.Errorf("BoxMuller(0.5, 0.5) = (%v, %v), want (0, 0)", n1, n2)
	} else {
		okPrint("BoxMuller (0.5, 0.5) anchor/(%v, %v)/(0, 0)", n1, n2)
	}
	// u2 = 0 is out of domain (log of zero).
	if _, _, ok := BoxMuller(0.5, 0.0); ok {
		t.Error("BoxMuller accepted u2=0, want refusal")
	} else {
		okPrint("BoxMuller u2=0 (out of domain)/rejected/rejected")
	}
	if _, _, ok := BoxMuller(-0.1, 0.5); ok {
		t.Error("BoxMuller accepted u1<0, want refusal")
	} else {
		okPrint("BoxMuller u1=-0.1 (out of domain)/rejected/rejected")
	}
	// Statistical: batch normals against N(0,1) moments.
	rng := rand.New(rand.NewSource(17))
	xs := BatchNormal(rng, 100000)
	if len(xs) != 100000 {
		t.Fatalf("BatchNormal returned %d samples, want 100000", len(xs))
	}
	checkMoments(t, "BatchNormal", xs, 0, 1, 5)
	okPrint("BatchNormal 100k seeded samples/got=mean+variance within 5 SE/want=within 5 SE")
	// Determinism: same seed, same stream.
	rng2 := rand.New(rand.NewSource(17))
	xs2 := BatchNormal(rng2, 100000)
	for i := range xs {
		if xs[i] != xs2[i] {
			t.Fatalf("BatchNormal not deterministic at %d", i)
		}
	}
	// No complexity test: O(n), no shape dependence.
}

func TestRejectionSample(t *testing.T) {
	var fx struct {
		TargetMax  float64 `json:"targetMax"`
		Envelope   float64 `json:"envelope"`
		M          float64 `json:"M"`
		AcceptRate float64 `json:"acceptRateTheoretical"`
	}
	loadFixture(t, "rejection.json", &fx)

	// Target: 6x(1-x) on [0,1] (density, max 1.5). Envelope: uniform(0,1).
	target := func(x float64) float64 { return 6 * x * (1 - x) }
	envelope := func(x float64) float64 { return 1 }

	rng := rand.New(rand.NewSource(23))
	n := 50000
	samples, candidates := RejectionSample(target, envelope, fx.M, n, rng)
	if len(samples) != n {
		t.Fatalf("RejectionSample returned %d samples, want %d", len(samples), n)
	}
	// Acceptance rate band: the theoretical rate is 1/M = 2/3; a wrong
	// accept ratio shows up as a rate near 0 or 1.
	rate := float64(n) / float64(candidates)
	se := math.Sqrt(fx.AcceptRate * (1 - fx.AcceptRate) / float64(candidates))
	if math.Abs(rate-fx.AcceptRate) > 6*se {
		t.Errorf("acceptance rate %v is more than 6 standard errors from the theoretical %v",
			rate, fx.AcceptRate)
	} else {
		okPrint("RejectionSample acceptance rate %d/%d/%.4f within 6 SE of %.4f",
			n, candidates, rate, fx.AcceptRate)
	}
	// Moments of the target: mean 1/2, variance 1/20.
	checkMoments(t, "RejectionSample", samples, 0.5, math.Sqrt(1.0/20.0), 5)
	okPrint("RejectionSample %d samples/got=mean+variance within 5 SE/want=within 5 SE", n)
	// All samples in the envelope's support.
	for _, x := range samples {
		if x < 0 || x > 1 {
			t.Fatalf("RejectionSample produced out-of-support sample %v", x)
		}
	}
	okPrint("RejectionSample %d accepted xs/got=all within support [0, 1]/want=same", n)
	// No complexity test: O(M*n) by the acceptance-rate check above.
}

func TestMetropolisHastings1D(t *testing.T) {
	var fx struct {
		Target    string  `json:"target"`
		Delta     float64 `json:"proposalHalfWidth"`
		ExpectAcc float64 `json:"expectedAccept"`
	}
	loadFixture(t, "metropolis.json", &fx)

	// Target: standard normal density, unnormalized (drop 1/sqrt(2*pi) —
	// the implementation must not need it).
	target := func(x float64) float64 { return math.Exp(-x * x / 2) }

	rng := rand.New(rand.NewSource(29))
	iters := 200000
	chain := MetropolisHastings1D(target, 0.0, 1.0, iters, rng)
	if len(chain) != iters {
		t.Fatalf("MetropolisHastings1D returned %d states, want %d", len(chain), iters)
	}
	// Burn-in: drop the first half.
	post := chain[iters/2:]

	// Acceptance-rate band: count moves.
	moves := 0
	for i := 1; i < len(chain); i++ {
		if chain[i] != chain[i-1] {
			moves++
		}
	}
	accRate := float64(moves) / float64(len(chain)-1)
	if accRate < 0.2 || accRate > 0.8 {
		t.Errorf("acceptance rate %v outside the plausible band [0.2, 0.8] — the accept ratio is suspect", accRate)
	} else {
		okPrint("MetropolisHastings1D acceptance rate/%.3f/within [0.2, 0.8]", accRate)
	}
	// Moments with autocorrelation-inflated tolerances: the effective
	// sample size of a random-walk chain is roughly n/10, so allow 10x
	// the naive standard error.
	checkMoments(t, "MetropolisHastings1D", post, 0, 1, 15)
	okPrint("MetropolisHastings1D %d post-burn-in states/got=mean+variance within 15 SE/want=within 15 SE", len(post))
	// Determinism: same seed reproduces the chain exactly.
	rng2 := rand.New(rand.NewSource(29))
	chain2 := MetropolisHastings1D(target, 0.0, 1.0, iters, rng2)
	for i := range chain {
		if chain[i] != chain2[i] {
			t.Fatalf("MetropolisHastings1D not deterministic at %d", i)
		}
	}
	// No complexity test: O(iters), no shape dependence.
}
