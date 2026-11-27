package graphs

import "alGorithms/internal/graph"

// [easy] ConnectedComponents labels the connected components of an
// undirected graph.
//
// Background. In an undirected graph, two vertices are connected when a
// path between them exists. Connectivity is an equivalence relation, so
// it partitions the vertices into components. A single traversal from
// any unlabeled vertex reaches exactly its whole component, so repeating
// a traversal from the next unlabeled vertex labels every component
// once.
//
// Definitions:
//   - component: a maximal set of vertices that are pairwise connected by
//     paths. A vertex with no edges forms a component by itself.
//   - label: label[v] is the identifier of v's component. Identifiers are
//     0, 1, ..., count-1.
//   - scan order: vertices are considered in increasing index order.
//   - traversal: a breadth-first or depth-first walk from a start vertex
//     that visits every vertex reachable from it.
//
// Steps:
//  1. Set every label to -1 and count to zero.
//  2. For v from 0 to N-1 in increasing order, if label[v] is -1, run a
//     traversal from v. Assign the current count to every vertex the
//     traversal reaches, including v. Then increase count by one.
//  3. Return count and the label array.
//
// Notes:
//   - The graph must be undirected, built with AddUndirected. A directed
//     graph would split into weak or strong components depending on the
//     traversal, which is not what this function reports.
//   - An isolated vertex is a component of size one.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func ConnectedComponents(g *graph.Graph) (count int, label []int) {
	// TODO: implement
	return 0, nil
}
