package scanning

import (
	"testing"

	"alGorithms/internal/complexity"
)

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	if p < lo || p > hi {
		t.Errorf("%s: measured exponent %.2f, expected in [%.2f, %.2f]", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

func TestLinearSearchComplexity(t *testing.T) {
	// Search for an absent value: forces the full scan (worst case).
	assertExponent(t, "linear search O(n)",
		complexity.FitExponent(20000, 5, func(n int) {
			LinearSearch(complexity.RandInts(n, 42), -1)
		}), 0.8, 1.3)
}

func TestBinarySearchComplexity(t *testing.T) {
	// Single searches are nanoseconds — too noisy to fit. Amplify by doing
	// many searches per measured call; O(log n) still shows exponent ~0.
	const probes = 20_000
	assertExponent(t, "binary search O(log n)",
		complexity.FitExponent(20000, 5, func(n int) {
			xs := complexity.SortedInts(n)
			for i := range probes {
				BinarySearch(xs, i%n)
			}
		}), -0.3, 0.3)
}

func TestKadaneComplexity(t *testing.T) {
	assertExponent(t, "kadane O(n)",
		complexity.FitExponent(20000, 5, func(n int) {
			Kadane(complexity.RandInts(n, 42))
		}), 0.8, 1.3)
}

func TestBoyerMooreComplexity(t *testing.T) {
	assertExponent(t, "boyer-moore majority O(n)",
		complexity.FitExponent(20000, 5, func(n int) {
			BoyerMooreMajority(complexity.RandInts(n, 42))
		}), 0.8, 1.3)
}

func TestKMPComplexity(t *testing.T) {
	assertExponent(t, "kmp O(n)",
		complexity.FitExponent(20000, 5, func(n int) {
			KMP(complexity.RandStr(n, 42), "abcd")
		}), 0.8, 1.4)
}

func TestRabinKarpComplexity(t *testing.T) {
	assertExponent(t, "rabin-karp O(n)",
		complexity.FitExponent(20000, 5, func(n int) {
			RabinKarp(complexity.RandStr(n, 42), "abcd")
		}), 0.8, 1.4)
}

func TestTwoPointerDedupComplexity(t *testing.T) {
	assertExponent(t, "two-pointer dedup O(n)",
		complexity.FitExponent(20000, 5, func(n int) {
			TwoPointerDedup(complexity.DuplicatedInts(n))
		}), 0.8, 1.3)
}
