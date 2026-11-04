package sorting

// [easy] InsertionSort returns xs sorted in ascending order.
//
// Background. Keep a sorted region at the front of the slice and an
// unsorted remainder behind it. The sorted region starts as just the
// first item. Each round, take the first item of the remainder (the
// key) and insert it where it belongs in the sorted region, the way
// one sorts a hand of cards: pick up a card, slide it left past
// every card larger than it, and set it down.
//
// The sliding is what makes this different from a swap-based sort.
// The key is held aside (opening a gap in the region), every larger
// item moves one slot toward the gap, and the key drops into the
// slot that remains. Each item involved moves once instead of the
// three writes a swap costs, so the total number of writes is about
// half of bubble sort's.
//
// Definitions:
//  - sorted region / unsorted remainder: a prefix that is sorted and
//     whose items are final, and the rest of the slice. The region
//     starts as the first item alone and grows by one per round.
//  - key: the item being inserted this round — the first item of the
//     remainder. It is held aside before the walk starts, which opens
//     the gap.
//  - gap: the slot vacated by lifting the key out. Items slide INTO
//     the gap; the gap moves left as they do.
//  - the walk: the backwards scan in step 2. It stops at an item that
//     is not greater than the key, or at the front of the slice.
//
// Steps:
//  1. Each round, take the first item of the unsorted remainder as
//     the key.
//  2. Walk backwards through the sorted region, moving every item
//     larger than the key one slot toward the gap. Two things can
//     stop the walk: an item that is not larger (the key belongs just
//     after it), or the front of the slice (the key is smaller than
//     everything in the region). Items equal to the key are not
//     moved, so equal keys land after their twins.
//  3. Drop the key into the slot the walk stopped at. The sorted
//     region has grown by one.
//  4. Stop when the remainder is empty.
//
// Notes:
//   - On already-sorted input every walk stops immediately, so the
//     whole sort is one linear scan. That is the linear best case
//     the test measures.
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
