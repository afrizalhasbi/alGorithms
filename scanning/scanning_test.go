package scanning

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"alGorithms/internal/complexity"
)

// ---- shared helpers ----

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	if p < lo || p > hi {
		t.Errorf("%s: measured exponent %.2f, expected in [%.2f, %.2f]", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestLinearSearch(t *testing.T) {
	// Correctness.
	xs := []int{5, 3, 8, 3, 1}
	if got := LinearSearch(xs, 8); got != 2 {
		t.Errorf("LinearSearch(xs, 8) = %d, want 2", got)
	}
	if got := LinearSearch(xs, 3); got != 1 {
		t.Errorf("LinearSearch: must return the FIRST match, got %d", got)
	}
	if got := LinearSearch(nil, 1); got != -1 {
		t.Errorf("LinearSearch(nil, 1) = %d, want -1", got)
	}
	for _, n := range []int{0, 1, 17, 1000} {
		in := complexity.RandInts(n, 11)
		for target := range 10 {
			if got, want := LinearSearch(in, target), slices.Index(in, target); got != want {
				t.Errorf("LinearSearch(%v, %d) = %d, want %d", in, target, got, want)
			}
		}
	}
	// Complexity: search for an absent value, forcing the full scan.
	assertExponent(t, "O(n)", complexity.FitExponent(20000, 5, func(n int) {
		LinearSearch(complexity.RandInts(n, 42), -1)
	}), 0.8, 1.3)
}

func TestBinarySearch(t *testing.T) {
	// Correctness.
	xs := []int{1, 3, 5, 7, 9, 11}
	for i, want := range []int{0, 1, 2, 3, 4, 5} {
		if got := BinarySearch(xs, xs[i]); got != want {
			t.Errorf("BinarySearch(xs, %d) = %d, want %d", xs[i], got, want)
		}
	}
	if got := BinarySearch(nil, 1); got != -1 {
		t.Errorf("BinarySearch(nil, 1) = %d, want -1", got)
	}
	for _, n := range []int{0, 1, 2, 17, 1000} {
		in := complexity.DuplicatedInts(n)
		for target := -1; target < n/2+2; target++ {
			got := BinarySearch(in, target)
			present := slices.Contains(in, target)
			if present != (got != -1) {
				t.Fatalf("BinarySearch(%d): returned %d but presence is %v", target, got, present)
			}
			if present && in[got] != target {
				t.Errorf("BinarySearch(%d) = %d, but xs[%d] = %d", target, got, got, in[got])
			}
		}
	}
	// Complexity: single searches are nanoseconds — amplify with many probes;
	// O(log n) still shows exponent ~0.
	const probes = 20_000
	assertExponent(t, "O(log n)", complexity.FitExponent(1000, 5, func(n int) {
		xs := complexity.SortedInts(n)
		for i := range probes {
			BinarySearch(xs, i%n)
		}
	}), -0.3, 0.3)
}

// kadaneBrute is the O(n^2) oracle: check every subarray sum.
func kadaneBrute(xs []int) int {
	best := 0 // empty subarray convention
	for i := range xs {
		sum := 0
		for j := i; j < len(xs); j++ {
			sum += xs[j]
			if sum > best {
				best = sum
			}
		}
	}
	return best
}

func TestKadane(t *testing.T) {
	// Correctness.
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6}, // classic: [4,-1,2,1]
		{[]int{-5, -3, -8}, 0},                    // all negative -> empty subarray
		{[]int{5}, 5},
		{[]int{}, 0},
		{[]int{1, 2, 3}, 6},
	}
	for _, c := range cases {
		if got := Kadane(c.in); got != c.want {
			t.Errorf("Kadane(%v) = %d, want %d", c.in, got, c.want)
		}
	}
	for _, n := range []int{1, 2, 50, 200} {
		in := complexity.RandInts(n, 5)
		if got, want := Kadane(in), kadaneBrute(in); got != want {
			t.Errorf("Kadane(%v) = %d, want %d", in, got, want)
		}
	}
	// Complexity.
	assertExponent(t, "O(n)", complexity.FitExponent(20000, 5, func(n int) {
		Kadane(complexity.RandInts(n, 42))
	}), 0.8, 1.3)
}

func TestTwoPointerDedup(t *testing.T) {
	// Correctness.
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 1}, 1},
		{[]int{1, 2, 3}, 3},
		{[]int{1, 1, 2, 2, 3}, 3},
	}
	for _, c := range cases {
		in := slices.Clone(c.in)
		if got := TwoPointerDedup(in); got != c.want {
			t.Errorf("TwoPointerDedup(%v) = %d, want %d", c.in, got, c.want)
		}
		if want := slices.Compact(slices.Clone(c.in)); !slices.Equal(in[:c.want], want) {
			t.Errorf("TwoPointerDedup(%v): deduped prefix = %v, want %v",
				c.in, previewInts(in[:c.want]), previewInts(want))
		}
	}
	for _, n := range []int{1, 17, 1000} {
		in := complexity.DuplicatedInts(n)
		orig := slices.Clone(in)
		k := TwoPointerDedup(in)
		want := slices.Compact(orig)
		if k != len(want) || !slices.Equal(in[:k], want) {
			t.Errorf("TwoPointerDedup: got k=%d prefix %s, want k=%d %s",
				k, previewInts(in[:k]), len(want), previewInts(want))
		}
	}
	// Complexity.
	assertExponent(t, "O(n)", complexity.FitExponent(20000, 5, func(n int) {
		TwoPointerDedup(complexity.DuplicatedInts(n))
	}), 0.8, 1.3)
}

func previewInts(xs []int) string {
	const max = 8
	suffix := ""
	if len(xs) > max {
		suffix = ", ..."
		xs = xs[:max]
	}
	return fmt.Sprintf("[%v%s]", xs, suffix)
}

func TestBoyerMooreMajority(t *testing.T) {
	// Correctness.
	cases := []struct {
		in    []int
		want  int
		found bool
	}{
		{[]int{1, 2, 1, 2, 1}, 1, true},
		{[]int{1, 2, 3}, 0, false},      // no majority
		{[]int{7}, 7, true},             // single element
		{[]int{}, 0, false},             // empty
		{[]int{4, 4, 4, 4}, 4, true},    // all equal
		{[]int{1, 1, 2, 2}, 0, false},   // exact tie
		{[]int{2, 2, 2, 1, 3}, 2, true}, // majority must be verified
	}
	for _, c := range cases {
		got, found := BoyerMooreMajority(c.in)
		if found != c.found || (found && got != c.want) {
			t.Errorf("BoyerMooreMajority(%v) = (%d, %v), want (%d, %v)", c.in, got, found, c.want, c.found)
		}
	}
	for _, n := range []int{1, 2, 50, 501} {
		withMaj := append(complexity.RandInts(n, 9), make([]int, n/2+1)...)
		for i := n; i < len(withMaj); i++ {
			withMaj[i] = 42
		}
		if got, found := BoyerMooreMajority(withMaj); !found || got != 42 {
			t.Errorf("BoyerMooreMajority(with majority) = (%d, %v), want (42, true)", got, found)
		}
	}
	// Complexity.
	assertExponent(t, "O(n)", complexity.FitExponent(20000, 5, func(n int) {
		BoyerMooreMajority(complexity.RandInts(n, 42))
	}), 0.8, 1.3)
}

func TestKMP(t *testing.T) {
	checkStringSearch(t, "kmp", KMP)
	assertExponent(t, "O(n)", complexity.FitExponent(20000, 5, func(n int) {
		KMP(complexity.RandStr(n, 42), "abcd")
	}), 0.8, 1.4)
}

func TestRabinKarp(t *testing.T) {
	checkStringSearch(t, "rabin-karp", RabinKarp)
	assertExponent(t, "O(n)", complexity.FitExponent(20000, 5, func(n int) {
		RabinKarp(complexity.RandStr(n, 42), "abcd")
	}), 0.8, 1.4)
}

func checkStringSearch(t *testing.T, name string, fn func(string, string) int) {
	t.Helper()
	text := "abracadabra"
	if got, want := fn(text, "cada"), strings.Index(text, "cada"); got != want {
		t.Errorf("%s: got %d, want %d", name, got, want)
	}
	if got := fn(text, "zzz"); got != -1 {
		t.Errorf("%s: absent pattern = %d, want -1", name, got)
	}
	if got := fn("", "abc"); got != -1 {
		t.Errorf("%s: empty text = %d, want -1", name, got)
	}
	if got, want := fn("aaa", ""), strings.Index("aaa", ""); got != want {
		t.Errorf("%s: empty pattern = %d, want %d", name, got, want)
	}
	for _, n := range []int{0, 1, 17, 1000} {
		for seed := int64(1); seed <= 3; seed++ {
			txt := complexity.RandStr(n, seed)
			if n >= 4 {
				start := int(seed*3) % (n - 3)
				pat := txt[start : start+3]
				if got, want := fn(txt, pat), strings.Index(txt, pat); got != want {
					t.Errorf("%s(%q, %q) = %d, want %d", name, txt, pat, got, want)
				}
			}
			randPat := complexity.RandStr(3, seed+100)
			if got, want := fn(txt, randPat), strings.Index(txt, randPat); got != want {
				t.Errorf("%s(%q, %q) = %d, want %d", name, txt, randPat, got, want)
			}
		}
	}
}
