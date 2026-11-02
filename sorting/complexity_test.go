package sorting

import (
	"testing"

	"alGorithms/internal/complexity"
)

func fit(t *testing.T, n0 int, f func(int)) float64 {
	t.Helper()
	return complexity.FitExponent(n0, 5, f)
}

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	if p < lo || p > hi {
		t.Errorf("%s: measured exponent %.2f, expected in [%.2f, %.2f]", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

func TestBubbleSortComplexity(t *testing.T) {
	assertExponent(t, "bubble avg O(n^2)",
		fit(t, 250, func(n int) { BubbleSort(complexity.RandInts(n, 42)) }), 1.7, 2.4)
	assertExponent(t, "bubble best O(n)",
		fit(t, 250, func(n int) { BubbleSort(complexity.SortedInts(n)) }), 0.7, 1.3)
}

func TestInsertionSortComplexity(t *testing.T) {
	assertExponent(t, "insertion avg O(n^2)",
		fit(t, 250, func(n int) { InsertionSort(complexity.RandInts(n, 42)) }), 1.7, 2.4)
	assertExponent(t, "insertion best O(n)",
		fit(t, 250, func(n int) { InsertionSort(complexity.SortedInts(n)) }), 0.7, 1.3)
}

func TestSelectionSortComplexity(t *testing.T) {
	// No best case: quadratic on every input shape.
	assertExponent(t, "selection O(n^2)",
		fit(t, 250, func(n int) { SelectionSort(complexity.RandInts(n, 42)) }), 1.7, 2.4)
	assertExponent(t, "selection sorted still O(n^2)",
		fit(t, 250, func(n int) { SelectionSort(complexity.SortedInts(n)) }), 1.7, 2.4)
}

func TestMergeSortComplexity(t *testing.T) {
	assertExponent(t, "merge O(n log n)",
		fit(t, 500, func(n int) { MergeSort(complexity.RandInts(n, 42)) }), 0.9, 1.4)
}

func TestQuickSortComplexity(t *testing.T) {
	assertExponent(t, "quick avg O(n log n)",
		fit(t, 500, func(n int) { QuickSort(complexity.RandInts(n, 42)) }), 0.9, 1.4)
	// Random pivots must rescue the adversarial (sorted) shape.
	assertExponent(t, "quick on sorted stays O(n log n)",
		fit(t, 500, func(n int) { QuickSort(complexity.SortedInts(n)) }), 0.9, 1.6)
}

func TestHeapSortComplexity(t *testing.T) {
	assertExponent(t, "heap O(n log n)",
		fit(t, 500, func(n int) { HeapSort(complexity.RandInts(n, 42)) }), 0.9, 1.4)
}

func TestRadixSortComplexity(t *testing.T) {
	// Values in [0, n) -> k ~ log n, so O(n*k) drifts slightly above 1.
	assertExponent(t, "radix ~O(n)",
		fit(t, 500, func(n int) { RadixSort(complexity.RandInts(n, 42)) }), 0.9, 1.6)
}

func TestTimSortComplexity(t *testing.T) {
	assertExponent(t, "timsort avg O(n log n)",
		fit(t, 500, func(n int) { TimSort(complexity.RandInts(n, 42)) }), 0.9, 1.5)
	assertExponent(t, "timsort best O(n)",
		fit(t, 500, func(n int) { TimSort(complexity.SortedInts(n)) }), 0.7, 1.3)
}
