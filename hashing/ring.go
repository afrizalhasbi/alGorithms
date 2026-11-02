package hashing

// [hard] Ring is consistent hashing: assign keys to nodes so that adding/removing
// a node only remaps ~1/n of the keys.
//
// Each node is hashed to many positions on a circle; a key belongs to the
// first node position clockwise from its own hash.
// Steps:
//  1. NewRing(nodes): for each node, place `replicas` (e.g. 100) points on
//     the ring: hash the node's name with a suffix ("node-3#17") using FNV1a.
//     Keep a sorted slice of all ring positions (uint32 wraps around: that
//     wraparound IS the circle).
//  2. GetNode(key): hash the key, binary-search the sorted positions for the
//     smallest position >= key's hash; wrap to positions[0] if none.
//  3. Map the position back to its node (keep a parallel slice of owners).
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
