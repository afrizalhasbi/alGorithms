package graphs

import "alGorithms/internal/graph"

// [hard] ArticulationPointsBridges returns the articulation points and
// bridges of an undirected graph.
//
// Background. Removing a vertex can split a component into more pieces;
// such a vertex is an articulation point. Removing an edge can do the
// same; such an edge is a bridge. One depth-first search finds both. A
// vertex is articulated when one of its children in the search forest
// cannot reach any vertex discovered before the vertex by a back edge.
// The same condition, with equality replaced by strict inequality, makes
// the tree edge to that child a bridge. The root of a search is special:
// it is articulated exactly when it has more than one child.
//
// Definitions:
//   - discovery time: disc[u] is the position of u in the depth-first
//     search, starting at zero.
//   - lowlink: low[u] is the smallest discovery time reachable from u's
//     subtree by following tree edges and then at most one back edge to
//     an ancestor.
//   - search child: a vertex entered directly from u by a tree edge.
//   - root: a vertex that starts a depth-first search. A graph with
//     several components has several roots.
//   - articulation point: a vertex whose removal increases the number of
//     connected components. Isolated vertices are not articulation
//     points.
//   - bridge: an edge whose removal increases the number of connected
//     components.
//   - simple graph: a graph with at most one edge between any pair of
//     vertices. Parallel edges are outside the contract.
//
// Steps:
//  1. Set disc and low of every vertex to -1 and the discovery counter to
//     zero. Create empty point and bridge lists.
//  2. For each vertex r in increasing index order that has disc -1,
//     start a search from r with no parent.
//  3. Entering u with parent p: set disc[u] and low[u] to the current
//     counter and increase the counter. Set a child counter to zero.
//  4. For each outgoing edge u->v in adjacency order, skip the single
//     edge back to p. If disc[v] is -1, enter v with parent u, increase
//     the child counter, then set low[u] to the smaller of low[u] and
//     low[v]. If p is not -1 and low[v] is at least disc[u], mark u as an
//     articulation point. If low[v] is greater than disc[u], mark the
//     edge u-v as a bridge. If disc[v] is not -1 and v is not p, set
//     low[u] to the smaller of low[u] and disc[v].
//  5. After all edges of u are examined, if p is -1 and the child counter
//     is greater than one, mark u as an articulation point.
//  6. Return the articulation points sorted in increasing order. Return
//     each bridge once with its smaller endpoint first, sorted by first
//     endpoint and then by second endpoint.
//
// Notes:
//   - The graph must be undirected, built with AddUndirected, and must
//     not contain parallel edges. Parallel edges would require tracking
//     edge identities rather than the parent vertex.
//   - A self-loop is never a bridge and does not make its vertex an
//     articulation point.
//   - Recursion depth can reach V on a path-shaped graph.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func ArticulationPointsBridges(g *graph.Graph) (points []int, bridges [][2]int) {
	// TODO: implement
	return nil, nil
}
