package graphs

import "alGorithms/internal/graph"

// [medium] IsBipartite reports whether an undirected graph can be split
// into two sides with every edge crossing between them, and returns one
// such split when it exists.
//
// Background. Assigning a side to every vertex is the same as assigning
// one of two colors. An edge forces its endpoints to have different
// colors. Starting from any vertex with color zero and propagating the
// opposite color across every edge either succeeds everywhere or reaches
// an edge whose endpoints already have the same color. That conflict
// means an odd cycle: the colors alternate around the cycle and the last
// edge closes it with equal colors. Conversely, a graph with no odd cycle
// can always be colored this way.
//
// Definitions:
//   - color: color[v] is 0 or 1 for a colored vertex and -1 for an
//     uncolored one.
//   - bipartite: a graph is bipartite when its vertices can be colored so
//     that every edge has endpoints of different colors.
//   - conflict: an edge whose two endpoints have the same color.
//   - component: a maximal set of vertices connected by paths. Each
//     component is colored independently.
//
// Steps:
//  1. Set every color to -1.
//  2. For v from 0 to N-1 in increasing order, if color[v] is -1, set it
//     to zero and traverse its component. During the traversal, for each
//     edge u->w: if w is uncolored, set color[w] to 1-color[u] and
//     continue from w; if w is colored and color[w] equals color[u],
//     return false together with the colors assigned so far.
//  3. If every component is colored without a conflict, return true and
//     the color array.
//
// Notes:
//   - The graph must be undirected, built with AddUndirected.
//   - A self-loop is a conflict and makes the graph non-bipartite.
//   - When the result is false, the returned colors are a partial
//     assignment and are not a valid bipartition.
//   - Isolated vertices receive color zero.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func IsBipartite(g *graph.Graph) (bool, []int) {
	// TODO: implement
	return false, nil
}
