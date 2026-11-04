package sorting

// [easy] SelectionSort returns xs sorted in ascending order.
//
// Background. Keep a sorted region at the front of the slice and an
// unsorted remainder behind it. The sorted region starts empty.
// Each round, find the smallest item in the remainder and move it to
// the boundary, which grows the sorted region by exactly one item.
// After as many rounds as there are items, the region is the whole
// slice.
//
// Moving the smallest item to the boundary is a swap with whatever
// currently sits at the boundary. That swap is the reason this sort
// is not stable: the swapped item jumps over the rest of the
// remainder, and if it is equal to something it jumps past, their
// relative order is lost. (A stable variant exists that shifts items
// instead, at the cost of more writes.)
//
// Definitions:
//  - sorted region / remainder: a final, sorted prefix of the slice,
//     and the rest. The region starts EMPTY (unlike insertion sort) —
//     nothing is known to be sorted until the first round finishes.
//  - round: one scan of the whole remainder plus one swap.
//  - the scan: the full pass over the remainder that finds its
//     smallest item. Ties are resolved by keeping the FIRST smallest
//     seen; the scan never overwrites on equality.
//  - boundary: the first position of the remainder. Swapping the
//     smallest item into it grows the region by one.
//
// Steps:
//  1. Each round, scan the whole remainder once to find its smallest
//     item. If several items tie for smallest, the scan naturally
//     finds the first of them.
//  2. Swap that item with the first item of the remainder. The sorted
//     region grows by one.
//  3. Stop once the remainder is down to a single item. The last
//     remaining item is necessarily the largest of what is left, so
//     scanning for it would learn nothing.
//
// Notes:
//   - Exactly O(n^2) comparisons on every input shape: there is no
//     best case, because the remainder is fully scanned every round
//     no matter what.
//   - Minimizes swaps (at most one per round), useful when writes
//     are expensive.
//
// Expected: best/average/worst O(n^2)
// Input: mutated in place
func SelectionSort(xs []int) []int {
	var i = 0
	for i < len(xs) {
		var s = xs[i]
		var si = i

		var j = i
		for j < len(xs) {
			// check if smaller, overwrite if so
			if xs[j] < s {
				s = xs[j]
				si = j
			}
			j += 1
		}
		xs[si] = xs[i]
		xs[i] = s
		i += 1
	}
	return xs
}
