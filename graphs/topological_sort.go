package graphs

import "alGorithms/internal/graph"

// [medium] TopologicalSort returns an ordering of the vertices of a
// directed acyclic graph in which every edge points forward.
//
// Background. A topological order is a linear order that respects every
// directed edge: if u points to v, then u must appear before v. Such an
// order exists exactly when the graph has no directed cycle. Kahn's
// method builds the order from the sources outward. A source is a vertex
// with no incoming edge. Removing a source and its outgoing edges can
// create new sources, which are appended to the order in turn. If the
// process stalls before every vertex is emitted, the remaining vertices
// lie on or behind a cycle.
//
// Definitions:
//   - indegree: indegree[v] is the number of incoming edges of v.
//   - source: a vertex whose indegree is zero.
//   - topological order: a permutation of all vertices such that for
//     every edge u->v, u appears before v.
//   - cycle: a directed path of at least one edge that returns to its
//     start. A self-loop is a cycle.
//   - queue: the vertices currently known to be sources. Vertices are
//     removed from the front and new sources are appended at the back.
//
// Steps:
//  1. Compute the indegree of every vertex by counting every incoming
//     edge.
//  2. Put every vertex with indegree zero into the queue, in increasing
//     index order. Create an empty order.
//  3. Remove the front vertex u, append u to the order, and examine each
//     outgoing edge u->v in adjacency order. Decrease indegree[v] by one.
//     If it becomes zero, append v to the queue.
//  4. Repeat step 3 until the queue is empty. If the order contains every
//     vertex, return it with true. Otherwise a cycle exists; return nil
//     with false.
//
// Notes:
//   - The graph must be directed.
//   - The order is not unique. Any order satisfying the definition is
//     correct; the queue policy above only makes the reference result
//     deterministic.
//   - An empty graph returns an empty order and true.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func TopologicalSort(g *graph.Graph) ([]int, bool) {
	// TODO: implement
	return nil, false
}
