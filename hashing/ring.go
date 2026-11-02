package hashing

// [hard] Ring is consistent hashing: assign keys to nodes so that adding/removing
// a node only remaps ~1/n of the keys.
//
// Background. The problem: a cluster of cache servers, and every key
// must go to one of them. The obvious answer, hash the key and take
// the remainder against the node count, breaks when a node joins or
// leaves: the remainder changes for nearly every key, so nearly
// every key moves to a different server and caches go cold
// everywhere. Consistent hashing gets the remapping down to about
// one nth of the keys.
//
// The trick is to stop dividing by the node count. Instead, hash
// both the nodes and the keys onto the same number line that wraps
// around at the top of the hash range (that wraparound is what makes
// it a circle). Each key belongs to the first node mark it meets
// moving clockwise from its own position. A node now "owns" the arc
// between its mark and the previous node's mark. When a node
// leaves, only the keys on its arc move to the next node; everyone
// else is untouched. Adding a node steals one arc's worth of keys,
// same idea.
//
// One mark per node is still not enough: the marks land at random
// spots on the circle, and with few nodes one of them is likely to
// draw a huge arc and take a hugely disproportionate share of keys.
// The fix is virtual nodes: stamp each node on the circle many
// times, under slightly different names, so each node owns many
// small arcs instead of one big one. Randomness averages out and
// the loads even up.
//
// Steps:
//  1. Construction: for each node, hash its name with a distinguishing
//     suffix once per replica, producing that many marks. Keep all
//     marks in one sorted list (hash values are plain unsigned
//     integers; the wraparound at the top of the range is the circle
//     closing) with a parallel list of which node owns each mark.
//  2. Lookup: hash the key and find the first mark at or after that
//     position in the sorted list. The owner of that mark is the
//     answer. If the key's hash is past the last mark, wrap around to
//     the first mark.
//  3. The lookup is a binary search over the sorted marks, so it gets
//     slower only logarithmically as nodes are added.
//
// Notes:
//   - This is where packages meet: the lookup is a binary search
//     from scanning/, the hashing is FNV1a from this package.
//     Algorithms compose.
//   - Try replicas of one first and watch how unfair the load is,
//     then raise it. The test requires no node to exceed half the
//     keys with 100 replicas.
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
