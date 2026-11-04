package sorting

// [medium] HeapSort returns xs sorted in ascending order.
//
// Background. A heap is a way of organizing items so that the largest
// one is always at the front. It is a binary tree drawn inside the
// slice itself: the item at each position has its two children a
// fixed distance further back, so no pointers or extra memory are
// needed, the positions in the slice are the tree. The heap
// condition says every item must be at least as large as its
// children. That condition does not sort the slice, but it forces
// the largest item of all to sit at the front.
//
// The heap does not exist at the start; the input is in arbitrary
// order. The algorithm has two phases: first rearrange the slice
// into a heap, then repeatedly remove the largest item. Removing
// shrinks the heap and grows a sorted region at the back of the
// slice, until the heap is gone and the slice is sorted.
//
// Definitions:
//  - heap condition: every item is at least as large as its children.
//  - front / back of the heap: the front is position 0 (the largest
//     item); the back is the last position the current heap size
//     covers. The heap shrinks from the back as items are extracted.
//  - children: the items at 2i+1 and 2i+2 for an item at position i,
//     valid only while those positions are inside the current heap
//     size. An item with no child inside the heap is a leaf.
//  - parent: the item at (i-1)/2 for an item at position i. Every
//     non-leaf position is a parent.
//  - repair (sift-down): the operation in step 1. It assumes both
//     subtrees below the item already satisfy the heap condition; that
//     assumption is what the build order in step 2 provides.
//
// Steps:
//  1. Repairing an item: compare an item with its children. If it is
//     smaller than one of them, swap it with the larger of the two
//     and repeat the comparison at the child's new position. Stop
//     when the item has no children, or when it is at least as large
//     as both of them. Each repair fixes one path from top to bottom.
//  2. Build the heap: repair every item that has children, starting
//     from the last parent and moving toward the front of the slice.
//     Working from the back matters: when an item is repaired,
//     everything below it has already been repaired, so each repair
//     has less and less to fix. When this phase ends the whole slice
//     satisfies the heap condition.
//  3. Extract: the front of the heap is the largest remaining item.
//     Swap it with the last item of the heap, count the heap as one
//     item shorter, and repair the front item (the swap may have put
//     a small item at the front, breaking the heap condition there).
//     The extracted item is now in its final sorted position, behind
//     the heap.
//  4. Repeat the extraction until the heap is empty. The sorted
//     region then covers the whole slice.
//
// Notes:
//   - Children of an item sit at fixed positions (double the
//     position, plus one and plus two), which is why no pointers
//     are needed.
//   - O(n log n) on every input shape, O(1) auxiliary space.
//   - Not stable.
//
// Expected: best/average/worst O(n log n)
// Input: mutated in place
func HeapSort(xs []int) []int {
	if len(xs) == 0 {
		return xs
	}

	// heap size
	s := len(xs)
	// index of last parent
	i := s/2 - 1

	for i >= 0 {
		repair(xs, i, s)
		// parents are contiguous
		i--
	}

	for s > 0 {
		// swap heap front with heap end
		head := xs[0]
		xs[0] = xs[s-1]
		xs[s-1] = head

		// heap is 1 item smaller
		s--
		// repair the front
		repair(xs, 0, s)
	}

	return xs
}

func repair(xs []int, i int, s int) {
	par := xs[i]
	p := 2*i + 1
	q := 2*i + 2

	// early stop if no child
	if p >= s {
		return
	}
	ch1 := xs[p]
	r := p
	var cand = ch1

	if q < s {
		// choose whichever child is bigger
		if ch1 < xs[q] {
			cand = xs[q]
			r = q
		}
	} else {
	}
	if par < cand {
		xs[r] = par
		xs[i] = cand
		repair(xs, r, s)
	} else {
	}
}
