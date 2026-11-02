package sorting

// [easy] InsertionSort returns xs sorted in ascending order.
//
// Grow a sorted prefix one item at a time, like sorting a hand of cards.
// Steps:
//  1. Treat the front of the slice as a sorted prefix — it starts as
//     just the first item — and extend it one item per round.
//  2. Each round takes the first item of the unsorted remainder (the
//     key) and holds it aside, opening a gap where it used to be.
//  3. Walk backwards through the sorted prefix, moving every item
//     larger than the key one slot toward the gap, until you meet an
//     item that is not larger. Two things can stop the walk: such an
//     item is found, or the front of the slice is reached (the key is
//     smaller than everything in the prefix). Equal items are never
//     moved, so equal keys land after their twins — that is what keeps
//     the sort stable.
//  4. Drop the key into the slot the walk stopped at, and continue
//     with the next item of the remainder. The sort ends when the
//     remainder is empty.
//
// Notes:
//   - Shifting instead of swapping halves the number of writes vs bubble sort.
//   - Very fast on nearly-sorted input; still O(n^2) on random input.
//   - Stable.
//
// Expected: best O(n), average/worst O(n^2)
// Input: mutated in place
func InsertionSort(xs []int) []int {
	for i := range len(xs) {
		k := xs[i]
		// reverse iteration
		for j := i - 1; j >= 0; j-- {
			if xs[j] > k {
				xs[j+1] = xs[j]
				if j == 0 {
					xs[j] = k
				}
			} else {
				xs[j+1] = k
				break
			}
		}
	}
	return xs
}
