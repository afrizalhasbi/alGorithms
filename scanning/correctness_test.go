package scanning

import (
	"slices"
	"strings"
	"testing"

	"alGorithms/internal/complexity"
)

func TestLinearSearchCorrect(t *testing.T) {
	xs := []int{5, 3, 8, 3, 1}
	if got := LinearSearch(xs, 8); got != 2 {
		t.Errorf("LinearSearch(xs, 8) = %d, want 2", got)
	}
	if got := LinearSearch(xs, 3); got != 1 {
		t.Errorf("LinearSearch: must return the FIRST match, got %d", got)
	}
	if got := LinearSearch(xs, 42); got != -1 {
		t.Errorf("LinearSearch(xs, 42) = %d, want -1", got)
	}
	if got := LinearSearch(nil, 1); got != -1 {
		t.Errorf("LinearSearch(nil, 1) = %d, want -1", got)
	}
	// Oracle: for random inputs, presence in the slice must match the result.
	for _, n := range []int{0, 1, 17, 1000} {
		in := complexity.RandInts(n, 11)
		for target := range 10 {
			got := LinearSearch(in, target)
			want := slices.Index(in, target)
			if got != want {
				t.Errorf("LinearSearch(%v, %d) = %d, want %d", in, target, got, want)
			}
		}
	}
}

func TestBinarySearchCorrect(t *testing.T) {
	xs := []int{1, 3, 5, 7, 9, 11}
	for i, want := range []int{0, 1, 2, 3, 4, 5} {
		if got := BinarySearch(xs, xs[i]); got != want {
			t.Errorf("BinarySearch(xs, %d) = %d, want %d", xs[i], got, want)
		}
	}
	if got := BinarySearch(xs, 6); got != -1 {
		t.Errorf("BinarySearch absent value = %d, want -1", got)
	}
	if got := BinarySearch(nil, 1); got != -1 {
		t.Errorf("BinarySearch(nil, 1) = %d, want -1", got)
	}
	// Oracle on random sorted input: found ⇔ actually present; a found index
	// must point at the target (any occurrence, duplicates allowed).
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

func TestKadaneCorrect(t *testing.T) {
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
	// Oracle on random input.
	for _, n := range []int{1, 2, 50, 200} {
		in := complexity.RandInts(n, 5)
		if got, want := Kadane(in), kadaneBrute(in); got != want {
			t.Errorf("Kadane(%v) = %d, want %d", in, got, want)
		}
	}
}

func TestTwoPointerDedupCorrect(t *testing.T) {
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
		if !slices.Equal(in[:c.want], slices.Compact(slices.Clone(c.in))) {
			t.Errorf("TwoPointerDedup(%v): deduped prefix = %v, want %v",
				c.in, in[:c.want], slices.Compact(slices.Clone(c.in)))
		}
	}
	// Oracle on random duplicate-heavy input.
	for _, n := range []int{1, 17, 1000} {
		in := complexity.DuplicatedInts(n)
		orig := slices.Clone(in)
		k := TwoPointerDedup(in)
		want := slices.Compact(orig)
		if k != len(want) || !slices.Equal(in[:k], want) {
			t.Errorf("TwoPointerDedup: got k=%d prefix %v, want k=%d %v", k, in[:k], len(want), want)
		}
	}
}

func TestBoyerMooreMajorityCorrect(t *testing.T) {
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
	// Oracle: brute-force count on random inputs, with and without majorities.
	for _, n := range []int{1, 2, 50, 501} {
		withMaj := append(complexity.RandInts(n, 9), make([]int, n/2+1)...)
		for i := n; i < len(withMaj); i++ {
			withMaj[i] = 42
		}
		if got, found := BoyerMooreMajority(withMaj); !found || got != 42 {
			t.Errorf("BoyerMooreMajority(with majority) = (%d, %v), want (42, true)", got, found)
		}
		if _, found := BoyerMooreMajority(complexity.RandInts(n, 9)); found {
			in := complexity.RandInts(n, 9)
			var counts int
			for _, x := range in {
				if x == 42 {
					counts++
				}
			}
			if counts*2 > len(in) {
				t.Errorf("BoyerMooreMajority(%v): claimed no majority but 42 wins", in)
			}
		}
	}
}

func TestStringSearchCorrect(t *testing.T) {
	text := "abracadabra"
	type searchFn func(string, string) int
	for name, fn := range map[string]searchFn{"kmp": KMP, "rabin-karp": RabinKarp} {
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
		// Oracle: random texts, pattern present (a substring) and random patterns.
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
}
