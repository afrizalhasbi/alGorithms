package scanning

// [easy] LinearSearch returns the index of target in xs, or -1 if absent.
//
// Background. Without any structure on the input (the slice is not
// sorted, so nothing can be skipped), the only correct method is to
// look at every item until the target turns up. That is linear
// search, and it is the baseline every faster search is measured
// against. This is also the fully implemented reference for the
// exercise pattern: read the doc, make the complexity test pass,
// keep the implementation minimal.
//
// Definitions:
//  - early exit: returning at the first equal item. What makes an
//     early target cheap and an absent one expensive; not an
//     optimization, the definition of the method.
//  - worst case: target absent or last — the full scan. The shape the
//     complexity test measures.
//
// Steps:
//  1. Walk the slice from front to back, comparing each item with the
//     target.
//  2. The first equal item ends the search immediately; return its
//     position. Early exit matters: it is what makes an early target
//     cheap and a late or absent target expensive.
//  3. If the walk reaches the end without a match, the target is not
//     present: return -1.
//
// Notes:
//   - The worst case is target absent (or last): it forces a full
//     scan, and that is exactly what the complexity test measures.
//
// Expected: O(n)
// Input: not mutated
func LinearSearch(xs []int, target int) int {
	for i, x := range xs {
		if x == target {
			return i
		}
	}
	return -1
}
