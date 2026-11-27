package graphs

import "alGorithms/internal/graph"

// [hard] TarjanSCC partitions a directed graph into its strongly
// connected components.
//
// Background. Two vertices are strongly connected when each can reach
// the other. This is an equivalence relation, so it partitions the
// vertices into strongly connected components. Tarjan's method uses one
// depth-first search. It numbers vertices in discovery order and keeps a
// stack of vertices whose component is not yet decided. For each vertex
// it computes a lowlink: the smallest discovery number reachable from
// its subtree using at most one edge back to a vertex still on the
// stack. A vertex whose lowlink equals its own discovery number is the
// root of a component, because nothing in its subtree can reach higher
// on the search path. Popping the stack down to that vertex yields
// exactly the component.
//
// Definitions:
//   - discovery number: index[v] is the position of v in the depth-first
//     search, starting at zero.
//   - lowlink: low[v] is the smallest discovery number reachable from v's
//     subtree by following tree edges and then at most one edge to a
//     vertex that is still on the stack.
//   - stack: the vertices that have been discovered but whose component
//     has not yet been completed. A vertex is on the stack from its
//     discovery until its component is popped.
//   - onStack: whether a vertex is currently on the stack.
//   - strongly connected component: a maximal set of vertices that are
//     pairwise reachable from one another. A single vertex with no
//     self-loop forms a component by itself.
//
// Steps:
//  1. Set index and lowlink of every vertex to -1, create an empty stack,
//     and set every onStack flag to false. Start a counter at zero.
//  2. For each vertex v in increasing index order that has index -1,
//     enter v.
//  3. Entering v: set index[v] and low[v] to the current counter and
//     increase the counter. Push v on the stack and set onStack[v].
//  4. Examine each outgoing edge v->w in adjacency order. If index[w] is
//     -1, enter w, then set low[v] to the smaller of low[v] and low[w].
//     If w is on the stack, set low[v] to the smaller of low[v] and
//     index[w]. If w is neither, ignore it.
//  5. After all edges of v are examined, if low[v] equals index[v], pop
//     vertices from the stack until v has been popped. The popped
//     vertices form one component; append it to the result.
//  6. Return the list of components.
//
// Notes:
//   - The graph must be directed.
//   - The order of components and the order of vertices inside a
//     component are not specified.
//   - A self-loop does not change the partition; its vertex is still a
//     component of size one.
//   - Recursion depth can reach V on a path-shaped graph.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func TarjanSCC(g *graph.Graph) [][]int {
	// TODO: implement
	return nil
}
