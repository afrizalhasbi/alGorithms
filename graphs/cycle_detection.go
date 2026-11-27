package graphs

import "alGorithms/internal/graph"

// [medium] HasCycleDirected reports whether a directed graph contains a
// directed cycle.
//
// Background. A depth-first search of a directed graph classifies every
// edge by the state of its destination. An edge to an unvisited vertex
// continues the search; an edge to a finished vertex points into a
// subtree that is already complete; an edge to a vertex that is still on
// the current recursion path closes a cycle. The last case is the only
// one that proves a cycle, because the path from that vertex to the
// current one plus the closing edge forms a directed loop.
//
// Definitions:
//   - white: a vertex that has not been entered.
//   - gray: a vertex that has been entered but whose outgoing edges have
//     not all been examined. Gray vertices are exactly the current
//     recursion path.
//   - black: a vertex whose outgoing edges have all been examined and
//     whose search has returned.
//   - back edge: an edge from the current vertex to a gray vertex.
//   - cycle: a path of at least one edge that returns to its start. A
//     self-loop is a cycle.
//
// Steps:
//  1. Mark every vertex white.
//  2. For each vertex v in increasing index order that is still white,
//     enter v: mark it gray, then examine its outgoing edges in adjacency
//     order.
//  3. For an edge v->w: if w is white, enter w and continue; if w is
//     gray, return true; if w is black, continue with the next edge.
//  4. After all edges of v are examined, mark v black and return from v.
//  5. If every search finishes without finding a gray destination,
//     return false.
//
// Notes:
//   - The graph is used as stored and must be directed.
//   - Recursion depth can reach V on a path-shaped graph.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func HasCycleDirected(g *graph.Graph) bool {
	// TODO: implement
	return false
}

// [medium] HasCycleUndirected reports whether an undirected graph
// contains a cycle.
//
// Background. A depth-first search of an undirected graph builds a
// forest of discovery edges. Every edge that is not the edge back to the
// parent either leads to an unvisited vertex, which extends the forest,
// or leads to an already visited vertex, which closes a loop. In an
// undirected graph the parent edge is the only edge that must be
// ignored; every other edge to a visited vertex proves a cycle.
//
// Definitions:
//   - visited: a vertex that has been entered by the search.
//   - parent: the vertex from which the search entered the current
//     vertex.
//   - discovery edge: an edge from a vertex to an unvisited neighbor; it
//     extends the search forest.
//   - cycle: a path of at least one edge that returns to its start. A
//     self-loop is a cycle.
//   - simple graph: a graph with at most one edge between any pair of
//     vertices. Parallel edges are outside the contract.
//
// Steps:
//  1. Mark every vertex unvisited.
//  2. For each vertex v in increasing index order that is unvisited,
//     enter v with no parent.
//  3. From a vertex u with parent p, examine each outgoing edge u->w in
//     adjacency order. If w equals p, skip that edge. If w is
//     unvisited, enter w with parent u. If w is visited and not p,
//     return true.
//  4. If every search finishes without finding such an edge, return
//     false.
//
// Notes:
//   - The graph must be undirected, built with AddUndirected, and must
//     not contain parallel edges. A self-loop is detected as an edge to
//     the current vertex.
//   - Recursion depth can reach V on a path-shaped graph.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func HasCycleUndirected(g *graph.Graph) bool {
	// TODO: implement
	return false
}
