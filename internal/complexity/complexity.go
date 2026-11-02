// Package complexity provides deterministic input generators and empirical
// big-O estimation used by all algorithm exercises.
package complexity

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// RandInts returns n deterministic pseudo-random ints in [0, n).
// Same seed -> same slice, every run.
func RandInts(n int, seed int64) []int {
	rng := rand.New(rand.NewSource(seed))
	xs := make([]int, n)
	for i := range xs {
		xs[i] = rng.Intn(n)
	}
	return xs
}

// SortedInts returns 0,1,...,n-1 (best case for bubble/insertion,
// worst case for naive quicksort).
func SortedInts(n int) []int {
	xs := make([]int, n)
	for i := range xs {
		xs[i] = i
	}
	return xs
}

// ReversedInts returns n-1,...,0.
func ReversedInts(n int) []int {
	xs := make([]int, n)
	for i := range xs {
		xs[i] = n - 1 - i
	}
	return xs
}

// DuplicatedInts returns a sorted slice with duplicates: 0,1,...,n/2,0,1,...
// (for TwoPointerDedup-style exercises).
func DuplicatedInts(n int) []int {
	xs := make([]int, n)
	m := n/2 + 1
	for i := range xs {
		xs[i] = i % m
	}
	return xs
}

// AllEqual returns n copies of the same value.
func AllEqual(n int) []int {
	xs := make([]int, n)
	for i := range xs {
		xs[i] = 7
	}
	return xs
}

// RandStr returns a deterministic pseudo-random string of length n over a
// 4-letter alphabet (forces frequent near-matches, good for string search).
func RandStr(n int, seed int64) string {
	rng := rand.New(rand.NewSource(seed))
	b := make([]byte, n)
	for i := range b {
		b[i] = byte('a' + rng.Intn(4))
	}
	return string(b)
}

// FitExponent runs f on doubling input sizes (n0, 2*n0, 4*n0, 8*n0), taking
// the median time of `runs` calls at each size, then returns the fitted
// exponent p in the model t(n) = c * n^p.
//
// Rule of thumb for interpreting p:
//
//	O(1)       -> ~0
//	O(log n)   -> ~0
//	O(n)       -> ~1
//	O(n log n) -> ~1.1 .. 1.3
//	O(n^2)     -> ~2
func FitExponent(n0, runs int, f func(n int)) float64 {
	const sizes = 4
	ns := make([]int, sizes)
	ts := make([]float64, sizes)
	for i := range ns {
		n := n0 << i
		ns[i] = n
		f(n) // warmup: populate caches, grow internal buffers
		seen := make([]float64, runs)
		for r := range runs {
			start := time.Now()
			f(n)
			seen[r] = float64(time.Since(start))
		}
		sort.Float64s(seen)
		ts[i] = seen[runs/2]
	}
	var lx, ly, lxx, lxy float64
	for i := range ns {
		x := math.Log(float64(ns[i]))
		y := math.Log(ts[i])
		lx += x
		ly += y
		lxx += x * x
		lxy += x * y
	}
	k := float64(sizes)
	return (k*lxy - lx*ly) / (k*lxx - lx*lx)
}
