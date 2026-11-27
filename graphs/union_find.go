package graphs

// [easy] UnionFind maintains disjoint sets over the integers 0..n-1 and
// answers whether two of them belong to the same set.
//
// Background. Each set is stored as a tree of parents: every element
// points to a parent, and following parents eventually reaches a root.
// The root is the name of the set. Two elements are in the same set
// exactly when they have the same root. Joining two sets means making
// one root point at the other. Searching for a root can be made almost
// constant by flattening the path: after a search, every element on the
// path points directly at the root.
//
// Definitions:
//   - parent: parent[x] is the next element toward the root of x's set.
//   - root: an element x with parent[x] == x. Every set has one root.
//   - representative: the root of the set containing x.
//   - rank: an upper bound on the height of the tree rooted at an
//     element. It starts at zero for every element.
//   - path compression: after Find(x), every element visited on the path
//     from x to the root has its parent changed to the root.
//   - union by rank: the root with the smaller rank is attached under the
//     root with the larger rank. When the ranks are equal, either root
//     may be attached, and the surviving root's rank increases by one.
//
// Steps:
//  1. Construction: parent[x] = x and rank[x] = 0 for every x.
//  2. Find(x): follow parent links from x until a root is reached. Then
//     walk the same path again and set every visited element's parent to
//     that root. Return the root.
//  3. Union(a, b): find the roots of a and b. If they are equal, the
//     sets are already joined; return false. Otherwise attach the root
//     with the smaller rank under the root with the larger rank, or, if
//     the ranks are equal, attach the root of b under the root of a and
//     increase the rank of a by one. Return true.
//  4. Connected(a, b): return whether Find(a) equals Find(b).
//
// Notes:
//   - Find must be iterative or otherwise avoid deep recursion; a chain
//     of n elements is possible before compression.
//   - The input index of Find and the endpoints of Union must lie in
//     0..n-1.
//   - With both path compression and union by rank, the amortized cost
//     per operation is O(alpha(n)), where alpha is the inverse Ackermann
//     function.
//
// Expected: O(alpha(n)) amortized per operation, O(n) space
// Input: mutated in place
type UnionFind struct {
	// TODO: fields
}

// NewUnionFind returns n singleton sets.
func NewUnionFind(n int) *UnionFind {
	// TODO: implement
	return nil
}

// Find returns the representative of x's set.
func (u *UnionFind) Find(x int) int {
	// TODO: implement
	return 0
}

// Union joins the sets containing a and b. It reports whether the sets
// were different before the call.
func (u *UnionFind) Union(a, b int) bool {
	// TODO: implement
	return false
}

// Connected reports whether a and b belong to the same set.
func (u *UnionFind) Connected(a, b int) bool {
	// TODO: implement
	return false
}
