package graphs

import "alGorithms/internal/graph"

// [easy] BFS walks a graph breadth first from a source vertex and returns
// the visit order together with the number of edges on the shortest path
// to every vertex.
//
// Background. Breadth-first search explores the graph in layers. The
// source is layer zero; every vertex reachable by one edge is layer one;
// every vertex first reached from layer one is layer two; and so on. A
// queue enforces the order: vertices enter at the back and leave at the
// front, so all vertices of one layer leave before any vertex of the
// next layer. Because the first time a vertex is reached is along a
// shortest path, the layer number is the shortest-path distance in
// edges.
//
// Definitions:
//   - queue: the list of discovered but not yet expanded vertices. New
//     vertices are appended at the back and removed from the front.
//   - discovered: a vertex is discovered when it is put in the queue. It
//     is never put in the queue twice.
//   - expanded: a vertex is expanded when it is removed from the queue and
//     its outgoing edges are examined.
//   - distance: dist[v] is the number of edges on a shortest path from the
//     source to v, or -1 when v is not reachable.
//   - visit order: the order in which vertices are expanded. The source
//     is first. Vertices of one layer appear before vertices of the next
//     layer.
//
// Steps:
//  1. Set every distance to -1. Set dist[src] to zero, mark src as
//     discovered, and put src in the queue.
//  2. Remove the front vertex u and append it to the visit order. For
//     each outgoing edge u->v in the order the edges appear in the
//     adjacency list, if v is not discovered, set dist[v] to dist[u]+1,
//     mark v as discovered, and append v to the queue.
//  3. Repeat step 2 until the queue is empty. Return the visit order and
//     the distance array.
//
// Notes:
//   - The graph is used as stored. A directed graph follows edge
//     direction only; an undirected graph built with AddUndirected has
//     edges in both directions.
//   - Only vertices reachable from src appear in the visit order.
//   - src must lie in 0..N-1.
//   - Every vertex enters the queue at most once and every edge is
//     examined once.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func BFS(g *graph.Graph, src int) (order []int, dist []int) {
	// TODO: implement
	return nil, nil
}
