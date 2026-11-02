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
	// TODO: implement
	return 0
}
