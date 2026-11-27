package graphs

import "alGorithms/internal/graph"

// [medium] Dijkstra returns the shortest-path distance from a source
// vertex to every vertex in a directed graph with non-negative edge
// weights.
//
// Background. Dijkstra's method grows a set of vertices whose shortest
// distance is final. It repeatedly takes the unsettled vertex with the
// smallest tentative distance, finalizes it, and relaxes its outgoing
// edges. The greedy choice is correct only because weights are
// non-negative: any other path to the chosen vertex would have to pass
// through an unsettled vertex whose tentative distance is at least as
// large, so it cannot be shorter. A min-heap keyed by tentative distance
// provides the next vertex efficiently. A vertex may be pushed more than
// once when a shorter distance is found; older entries are recognized as
// stale when their stored distance is larger than the current one.
//
// Definitions:
//   - tentative distance: dist[v] is the best known distance from src to
//     v, or -1 when no path is known yet.
//   - settled: a vertex whose tentative distance has been removed from
//     the heap as the current minimum and is therefore final.
//   - relaxation: for an edge u->v with weight w, the candidate distance
//     dist[u]+w. If v is unreached or the candidate is smaller than
//     dist[v], dist[v] becomes the candidate.
//   - stale entry: a heap entry (d, v) with d greater than dist[v]. It is
//     ignored.
//   - min-heap: a priority queue that returns the entry with the smallest
//     distance first.
//
// Steps:
//  1. Set every distance to -1. Set dist[src] to zero and push the entry
//     (0, src) into the min-heap.
//  2. Remove the smallest entry (d, u) from the heap. If d is greater
//     than dist[u], the entry is stale; discard it and repeat.
//  3. For each outgoing edge u->v with weight w, compute d+w. If dist[v]
//     is -1 or d+w is smaller than dist[v], set dist[v] to d+w and push
//     (dist[v], v) into the heap.
//  4. Repeat steps 2 and 3 until the heap is empty. Return the distance
//     array.
//
// Notes:
//   - Every weight must be non-negative. Negative weights invalidate the
//     greedy choice and are outside the contract.
//   - Unreachable vertices keep distance -1.
//   - src must lie in 0..N-1.
//   - The heap can be a local binary heap; no library is required.
//
// Expected: O((V+E) log V) time, O(V+E) space
// Input: not mutated
func Dijkstra(g *graph.Graph, src int) []int {
	// TODO: implement
	return nil
}
