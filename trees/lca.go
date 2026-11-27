package trees

// [hard] LCA answers lowest-common-ancestor queries on a fixed binary
// tree using binary lifting.
//
// Background. The lowest common ancestor of two nodes is the deepest node
// that is an ancestor of both. A direct answer walks both ancestor chains
// upward, which costs the tree height per query. Binary lifting
// preprocesses the tree so that a query costs O(log n). Every node stores
// its 1st, 2nd, 4th, 8th, ... ancestor. A query first lifts the deeper
// node to the same depth as the other. If the nodes are then equal, that
// node is the answer. Otherwise both are lifted together from the largest
// power of two downward, as long as their ancestors at that distance
// differ. When no further common lift is possible, their immediate
// parents are the answer.
//
// Definitions:
//   - ancestor: a node reachable by following parent links upward,
//     including the node itself.
//   - lowest common ancestor: the deepest ancestor shared by both nodes.
//   - depth: the number of edges from the root to a node. The root has
//     depth zero.
//   - up[k][v]: the ancestor of v at distance 2^k, or the root when 2^k is
//     larger than v's depth. The root's ancestor at every level is the
//     root itself.
//   - LOG: the smallest number of levels such that 2^LOG is greater than
//     the maximum depth. A value such as the number of bits in n is
//     sufficient.
//
// Steps:
//  1. Construction: traverse the tree from the root. Record every node's
//     value, its depth, and its immediate parent. The root's parent is the
//     root itself, and its depth is zero.
//  2. Build up[0][v] as the parent of v. For k from 1 to LOG-1, set
//     up[k][v] to up[k-1][ up[k-1][v] ].
//  3. Query(a, b): if a and b are the same node, return a. If depth[a] is
//     less than depth[b], exchange them. Lift a by the difference in
//     depths, using the binary digits of the difference: for each set bit
//     k, set a to up[k][a]. If a now equals b, return a.
//  4. Otherwise, for k from LOG-1 down to zero, if up[k][a] differs from
//     up[k][b], set both a and b to those ancestors. After the loop, a
//     and b are distinct children of the answer; return up[0][a].
//
// Notes:
//   - Node values must be unique, and both queried values must be present
//     in the tree.
//   - The tree is not modified after construction.
//   - Construction visits every node once per level, so it takes
//     O(n log n) time and space.
//
// Expected: O(n log n) construction, O(log n) per query, O(n log n) space
// Input: not mutated
type LCA struct {
	// TODO: fields
}

// NewLCA preprocesses root for ancestor queries.
func NewLCA(root *TreeNode) *LCA {
	// TODO: implement
	return nil
}

// Query returns the value of the lowest common ancestor of a and b.
func (l *LCA) Query(a, b int) int {
	// TODO: implement
	return 0
}
