package graphs

import "alGorithms/internal/graph"

// [medium] Prim returns a minimum spanning tree of the connected
// component containing a source vertex, together with its total weight.
//
// Background. Prim's method grows one tree from a start vertex. At every
// step it takes the lightest edge with one endpoint inside the tree and
// one endpoint outside. The cut property guarantees that this edge is
// safe: it belongs to some minimum spanning tree of the component. The
// tree grows by one vertex per step until no edge leaves it. A min-heap
// of frontier edges makes the lightest choice efficient. Edges whose
// outside endpoint has already entered the tree are stale and are
// skipped.
//
// Definitions:
//   - tree: the set of vertices already connected to the source by chosen
//     edges.
//   - frontier edge: an edge with exactly one endpoint in the tree. It is
//     a candidate for the next step.
//   - stale entry: a heap entry whose outside endpoint is already in the
//     tree. It is discarded.
//   - edge order: increasing weight, then increasing endpoint inside the
//     tree, then increasing endpoint outside the tree. This tie order is
//     only for determinism; any minimum tree is correct.
//   - component: a maximal set of vertices connected by paths.
//
// Steps:
//  1. If src is outside 0..N-1, return weight zero and an empty list.
//  2. Mark src as in the tree. Push every edge from src to an unvisited
//     vertex into the min-heap. Initialize the total weight to zero and
//     the chosen edge list to empty.
//  3. Remove the smallest edge (u, v, w) from the heap. If v is already
//     in the tree, the entry is stale; discard it and repeat.
//  4. Otherwise add the edge to the chosen list, add w to the total, mark
//     v as in the tree, and push every edge from v to a vertex not yet in
//     the tree into the heap.
//  5. Repeat steps 3 and 4 until the heap is empty. Return the total
//     weight and the chosen edges.
//
// Notes:
//   - The graph must be undirected, built with AddUndirected.
//   - Only the component containing src is spanned. Vertices in other
//     components are ignored.
//   - The returned TreeEdge has U smaller than V.
//   - The heap can be a local binary heap; no library is required.
//
// Expected: O(E log V) time, O(V+E) space
// Input: not mutated
func Prim(g *graph.Graph, src int) (weight int, edges []graph.TreeEdge) {
	// TODO: implement
	return 0, nil
}
