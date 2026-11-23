package sorting

import (
	"math/rand"
)

// [medium] QuickSort returns xs sorted in ascending order.
//
// Background. Pick one item of the slice and call it the pivot. A
// partition rearranges the slice so that every item smaller than the
// pivot ends up before it and every item larger ends up after it.
// After one partition the pivot is somewhere between the two groups,
// and everything on each side of it still needs sorting, but no item
// can ever cross the pivot again. Sorting is then just the same
// operation applied to each side, and each side is smaller than the
// whole, so the recursion stops.
//
// The pivot choice decides everything. If the pivot always lands
// near the middle, the halves halve each time and the total work is
// about n log n. If the pivot is always the smallest or largest
// item (which happens deterministically on sorted input when the
// pivot is always taken from the front), one side is empty and the
// work is quadratic. Choosing the pivot at random makes that input
// harmless.
//
// This partition is the Hoare scheme: two cursors start at the two
// ends and move inward. The left cursor stops at the first item that
// is at least the pivot; the right cursor stops at the first item
// that is at most the pivot. When both are stopped, those two items
// are on the wrong sides of each other, so they are swapped. The
// cursors resume. When they cross, the partition is done: the two
// final cursor positions split the slice into the two sides to recurse
// on, or leave one item between the sides when they meet on it.
//
// Definitions:
//   - pivot: the item chosen in step 2. Partitions and comparisons use
//     its value; the item itself moves during partitioning like any
//     other.
//   - partition: one application of the two-cursor scheme, rearranging
//     a range so that no item larger than the pivot value sits before
//     the split and none smaller sits after it.
//   - cursors: the two moving positions that scan the range inward.
//     The left one advances past items smaller than the pivot value;
//     the right one advances past items larger. Each stops at the
//     first item that belongs on the other side.
//   - split: where the partition cuts the range, determined by the
//     cursors' final positions. It is NOT necessarily the pivot's final
//     position — unlike the Lomuto scheme, no item is guaranteed to be
//     in its final position after one partition. When the cursors pass
//     each other the split is a single boundary; when they meet on the
//     same item the two ranges are separated by that item, which lies
//     in neither of them.
//
// Steps:
//
//  1. Exit condition: a range of one or zero items is sorted.
//
//  2. Choose the pivot uniformly at random from the current range, using
//     one generator with a fixed seed created once and shared by the
//     whole recursion, so runs are reproducible.
//
//  3. Partition the range with the two-cursor scheme described above,
//     comparing against the pivot's value. A cursor never moves past the
//     other one: a scan that reaches the opposite cursor stops there, so
//     neither cursor can read outside the range. After swapping the two
//     stopped items, advance both cursors one step inward before
//     resuming; without that advance, equal items can leave both cursors
//     stopped forever.
//
//  4. When the cursors cross, the partition is done: swap nothing, and
//     keep both final positions. Let k be the right cursor's final
//     position and j the left cursor's. Recurse on xs[:k+1] and xs[j:].
//
//     Usually the cursors pass each other, so j == k+1 and the two
//     ranges are divided by the single boundary k. But the cursors can
//     also meet on one item whose value equals the pivot: that pair is
//     swapped (a no-op) and advanced, leaving k = m-1 and j = m+1 around
//     the item at m. That item is not smaller or larger than the pivot,
//     so it belongs to neither side and both ranges must exclude it —
//     which is why the right range starts at j, not at k+1.
//
//     In the extreme case the meeting item is at index 0 and the right
//     cursor ends at k = -1. Then xs[:k+1] is empty and xs[j:] is the
//     rest of the range, which is correct and still shorter than the
//     whole. No item is guaranteed to be in its final position by the
//     split itself.
//
// Notes:
//   - Sorted or all-equal input is the classic worst case for a
//     naive pivot (front or back item). Random pivots fix the sorted
//     case; all-equal input still stresses the partition, which is
//     why it is one of the tested shapes.
//   - Partitioning happens inside the slice, so no extra allocation
//     is needed.
//   - Not stable.
//
// Expected: best/average O(n log n), worst O(n^2) (rare with random pivots)
// Input: mutated in place
func QuickSort(xs []int) []int {

	r := rand.New(rand.NewSource(100))
	return quickSort(r, xs)

}

func quickSort(r *rand.Rand, xs []int) []int {
	if len(xs) <= 1 {
		return xs
	}
	pi := r.Intn(len(xs))
	piv := xs[pi]
	j := 0
	k := len(xs) - 1

	for j <= k {
		// left candidate
		for xs[j] < piv && j <= k {
			j++
		}

		// right candidate
		for xs[k] > piv && j <= k {
			k--
		}

		if j > k {
			break
		}

		tmp := xs[j]
		xs[j] = xs[k]
		xs[k] = tmp

		j++
		k--
	}

	quickSort(r, xs[:k+1])
	quickSort(r, xs[j:])
	return xs
}
