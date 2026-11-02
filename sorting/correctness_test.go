package sorting

import (
	"fmt"
	"slices"
	"testing"

	"alGorithms/internal/complexity"
)

var sorters = []struct {
	name string
	sort func([]int) []int
}{
	{"bubble", BubbleSort},
	{"insertion", InsertionSort},
	{"selection", SelectionSort},
	{"merge", MergeSort},
	{"quick", QuickSort},
	{"heap", HeapSort},
	{"radix", RadixSort},
	{"timsort", TimSort},
}

// checkSorted compares got against the stdlib oracle for "sorted version of in".
func checkSorted(t *testing.T, name string, in, got []int) {
	t.Helper()
	want := slices.Clone(in)
	slices.Sort(want)
	if !slices.Equal(got, want) {
		t.Errorf("%s: sorted(%v) = %v, want %v", name, in, got, want)
	}
}

func TestSortingFixedCases(t *testing.T) {
	cases := [][]int{
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
	for _, s := range sorters {
		for _, in := range cases {
			in := slices.Clone(in)
			checkSorted(t, fmt.Sprintf("%s/%v", s.name, in), in, s.sort(in))
		}
	}
}

func TestSortingShapeSweeps(t *testing.T) {
	shapes := []struct {
		name string
		gen  func(int) []int
	}{
		{"rand", func(n int) []int { return complexity.RandInts(n, 7) }},
		{"sorted", complexity.SortedInts},
		{"reversed", complexity.ReversedInts},
		{"allequal", complexity.AllEqual},
	}
	sizes := []int{0, 1, 2, 17, 1000}
	for _, s := range sorters {
		for _, sh := range shapes {
			for _, n := range sizes {
				in := sh.gen(n)
				checkSorted(t, fmt.Sprintf("%s/%s/%d", s.name, sh.name, n), in, s.sort(in))
			}
		}
	}
}

func TestSortingDoesNotMutateInput(t *testing.T) {
	// Every sort here is documented "Input: not mutated".
	in := complexity.RandInts(50, 3)
	orig := slices.Clone(in)
	for _, s := range sorters {
		work := slices.Clone(in)
		s.sort(work)
		if !slices.Equal(work, orig) {
			t.Errorf("%s mutated its input", s.name)
		}
	}
}
