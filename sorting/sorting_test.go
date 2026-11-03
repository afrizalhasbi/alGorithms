package sorting

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"alGorithms/internal/complexity"
)

// ---- shared helpers ----

func fitExp(t *testing.T, n0 int, f func(int)) float64 {
	t.Helper()
	return complexity.FitExponent(n0, 5, f)
}

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	// Being faster than expected is fine (measurement noise, machine
	// variance); only being slower than the upper bound fails.
	if p > hi {
		t.Errorf("%s: measured exponent %.2f, expected at most %.2f", name, p, hi)
	} else if p < lo {
		t.Logf("%s: exponent %.2f below expected range [%.2f, %.2f] — faster than expected, ok", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

// preview renders at most 8 elements to keep failure output readable.
func preview(xs []int) string {
	const max = 8
	suffix := ""
	if len(xs) > max {
		suffix = ", ..."
		xs = xs[:max]
	}
	return fmt.Sprintf("[%s%s]", strings.Trim(fmt.Sprint(xs), "[]"), suffix)
}

var shapes = []struct {
	name string
	gen  func(int) []int
}{
	{"rand", func(n int) []int { return complexity.RandInts(n, 7) }},
	{"sorted", complexity.SortedInts},
	{"reversed", complexity.ReversedInts},
	{"allequal", complexity.AllEqual},
}

var fixedCases = [][]int{
	{},                    // empty
	{1},                   // single
	{2, 1},                // pair
	{3, 1, 2},             // odd size
	{2, 2, 2},             // all equal
	{3, 1, 3, 1},          // duplicates
	{5, 4, 3, 2, 1},       // reversed
	{1, 2, 3, 4, 5},       // sorted
	{1, 9, 1, 9, 1, 9, 1}, // alternating duplicates
}

// okPrint prints a success line styled like the failure lines: no
// file:line prefix or indentation, always visible.
func okPrint(format string, args ...any) {
	fmt.Printf("ok: "+format+"\n", args...)
}

// checkSorting runs every correctness case against the stdlib oracle.
// Sorters mutate their input in place.
func checkSorting(t *testing.T, name string, sort func([]int) []int) {
	t.Helper()
	check := func(in []int) {
		inStr := preview(in) // captured before the in-place sort
		want := slices.Clone(in)
		slices.Sort(want)
		got := sort(in)
		if !slices.Equal(got, want) {
			// Print manually (then just fail) so the output has no
			// file:line prefix or indentation.
			fmt.Printf("%s: in: %s, got: %s, want: %s\n", name, inStr, preview(got), preview(want))
			t.Fail()
		}
		okPrint("%s %s/%s/%s", name, inStr, preview(got), preview(want))
	}
	for _, in := range fixedCases {
		check(in)
	}
	sizes := []int{0, 1, 2, 17, 1000}
	for _, sh := range shapes {
		for _, n := range sizes {
			check(sh.gen(n))
		}
	}
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestBubbleSort(t *testing.T) {
	checkSorting(t, "bubble", BubbleSort)
	assertExponent(t, "avg O(n^2)", fitExp(t, 250, func(n int) { BubbleSort(complexity.RandInts(n, 42)) }), 1.7, 2.4)
	assertExponent(t, "best O(n)", fitExp(t, 250, func(n int) { BubbleSort(complexity.SortedInts(n)) }), 0.7, 1.3)
}

func TestInsertionSort(t *testing.T) {
	checkSorting(t, "insertion", InsertionSort)
	assertExponent(t, "avg O(n^2)", fitExp(t, 250, func(n int) { InsertionSort(complexity.RandInts(n, 42)) }), 1.7, 2.4)
	assertExponent(t, "best O(n)", fitExp(t, 250, func(n int) { InsertionSort(complexity.SortedInts(n)) }), 0.7, 1.3)
}

func TestSelectionSort(t *testing.T) {
	checkSorting(t, "selection", SelectionSort)
	// No best case: quadratic on every input shape.
	assertExponent(t, "rand O(n^2)", fitExp(t, 250, func(n int) { SelectionSort(complexity.RandInts(n, 42)) }), 1.7, 2.4)
	assertExponent(t, "sorted still O(n^2)", fitExp(t, 250, func(n int) { SelectionSort(complexity.SortedInts(n)) }), 1.7, 2.4)
}

func TestMergeSort(t *testing.T) {
	checkSorting(t, "merge", MergeSort)
	assertExponent(t, "O(n log n)", fitExp(t, 500, func(n int) { MergeSort(complexity.RandInts(n, 42)) }), 0.9, 1.4)
}

func TestQuickSort(t *testing.T) {
	checkSorting(t, "quick", QuickSort)
	assertExponent(t, "avg O(n log n)", fitExp(t, 500, func(n int) { QuickSort(complexity.RandInts(n, 42)) }), 0.9, 1.4)
	// Random pivots must rescue the adversarial (sorted) shape.
	assertExponent(t, "sorted stays O(n log n)", fitExp(t, 500, func(n int) { QuickSort(complexity.SortedInts(n)) }), 0.9, 1.6)
}

func TestHeapSort(t *testing.T) {
	checkSorting(t, "heap", HeapSort)
	assertExponent(t, "O(n log n)", fitExp(t, 500, func(n int) { HeapSort(complexity.RandInts(n, 42)) }), 0.9, 1.4)
}

func TestRadixSort(t *testing.T) {
	checkSorting(t, "radix", RadixSort)
	// Values in [0, n) -> k ~ log n, so O(n*k) drifts slightly above 1.
	assertExponent(t, "~O(n)", fitExp(t, 500, func(n int) { RadixSort(complexity.RandInts(n, 42)) }), 0.9, 1.6)
}

func TestTimSort(t *testing.T) {
	checkSorting(t, "timsort", TimSort)
	assertExponent(t, "avg O(n log n)", fitExp(t, 500, func(n int) { TimSort(complexity.RandInts(n, 42)) }), 0.9, 1.5)
	assertExponent(t, "best O(n)", fitExp(t, 500, func(n int) { TimSort(complexity.SortedInts(n)) }), 0.7, 1.3)
}
