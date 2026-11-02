package hashing

// [hard] Ring is consistent hashing: assign keys to nodes so that adding/removing
// a node only remaps ~1/n of the keys.
//
// Each node is hashed to many positions on a circle; a key belongs to the
// first node position clockwise from its own hash.
// Steps:
//  1. Construction: imagine a circle of hash values. Place many marks
//     on the circle for each node — for each node, stamp the given
//     number of copies of its name (a small distinguishing suffix per
//     copy), hash each, and record the hash as a mark owned by that
//     node. Keep all marks together in one sorted list, with a
//     parallel list remembering each mark's owner. The hash naturally
//     wraps around at its maximum: that wraparound is the circle.
//  2. Lookup: hash the key to a position on the circle. Find the
//     first mark at or after that position, going clockwise — a
//     binary search over the sorted marks. If the key's position is
//     past the last mark, the search wraps around to the very first
//     mark. The owner of the mark found is the node that gets the key.
//  3. The many marks per node (virtual nodes) are what spread the
//     load evenly: with one mark per node, whichever node draws the
//     largest empty stretch of circle takes a hugely disproportionate
//     share of keys.
//
// Notes:
//   - This is where packages meet: step 2 is binary search from scanning/,
//     step 1 uses FNV1a. Algorithms compose.
//   - Replicas (virtual nodes) smooth out load imbalance; try replicas=1
//     first and watch how unfair it is, then bump it up.
//
// Expected: O(log n) lookup in ring positions
// Input: not mutated
type Ring struct {
	// TODO: fields
}

func NewRing(nodes []string, replicas int) *Ring {
	// TODO: implement
	return nil
}

func (r *Ring) GetNode(key string) string {
	// TODO: implement
	return ""
}
