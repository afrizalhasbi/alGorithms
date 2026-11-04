package scanning

// [medium] TwoPointerDedup removes duplicates from the sorted slice xs in place and
// returns the length k of the deduplicated prefix.
//
// Background. The slice is sorted, so all copies of a value sit next
// to each other. Deduplication therefore never needs to compare an
// item with more than its immediate predecessor. The result is
// produced in place, inside the same slice: the front of the slice
// becomes the deduplicated data (the kept region) and the rest is
// overwritten with leftovers. The function returns how long the kept
// region is; the caller knows only that much of the slice is
// meaningful.
//
// Two positions are tracked. The boundary of the kept region moves
// forward only when a genuinely new value is found. The reading
// position moves forward over every item. The name of the algorithm
// comes from these two positions moving at different speeds.
//
// Definitions:
//  - kept region: the prefix of the slice holding the deduplicated
//     data so far. Starts as the first item alone; its length is the
//     returned value.
//  - read position: where the scan currently is. Moves forward over
//     every remaining item, one at a time.
//  - duplicate: an item equal to the LAST KEPT value, not to the
//     previously read item — those coincide only because the input is
//     sorted.
//  - write slot: the position just past the kept region. A new value
//     is copied into it; a duplicate leaves it untouched.
//
// Steps:
//  1. An empty slice returns zero.
//  2. The kept region starts as just the first item (one item is
//     always unique).
//  3. Read the remaining items in order. An item equal to the last
//     kept one is a duplicate: skip it and leave whatever sits in the
//     unexamined area untouched. A different value is copied into the
//     slot right after the kept region, and the region grows by one.
//  4. When every item has been read, the kept region is the answer;
//     return its length. Nothing past it is meaningful.
//
// Notes:
//   - This is the canonical "write index" pattern; it shows up
//     everywhere (partitioning, compaction, stable partition
//     variants).
//   - Input must be sorted, otherwise equal items are not adjacent
//     and the single-predecessor comparison is not enough.
//
// Expected: O(n)
// Input: mutated in place (only up to the returned length is meaningful)
func TwoPointerDedup(xs []int) int {
	if len(xs) == 0 {
		return 0
	}

	// retained len
	// consider this the begin
	// index of the dupe slice
	i := 1
	// scanner len
	// consider this the ending
	// index of the dupe slice
	j := 1
	// if i == j, then theres no dupe slice

	var prev = xs[0]
	for j < len(xs) {
		new := xs[j]
		if new != prev {
			xs[i] = new
			prev = new
			i++
		}
		j++
	}

	return i
}
