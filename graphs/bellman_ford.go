package graphs

import "alGorithms/internal/graph"

// [medium] BellmanFord returns the shortest-path distance from a source
// vertex to every vertex in a directed graph that may contain negative
// edge weights, and reports whether a negative cycle is reachable from
// the source.
//
// Background. The algorithm relaxes every edge repeatedly. After k full
// passes, every shortest path that uses at most k edges is known, so V-1
// passes suffice for any shortest path that does not repeat a vertex. A
// shortest path never needs to repeat a vertex unless a negative cycle
// makes it profitable to loop. One extra pass therefore detects a
// negative cycle: if any edge can still be relaxed after V-1 passes, the
// graph contains a reachable cycle whose total weight is negative.
//
// Definitions:
//   - tentative distance: dist[v] is the best known distance from src to
//     v, or -1 when no path is known yet.
//   - relaxation: for an edge u->v with weight w, the candidate distance
//     dist[u]+w. It applies only when dist[u] is not -1. If v is
//     unreached or the candidate is smaller than dist[v], dist[v]
//     becomes the candidate.
//   - pass: one sweep over every edge in the graph, in any fixed order.
//   - negative cycle: a directed cycle whose total weight is negative.
//   - reachable negative cycle: a negative cycle with at least one vertex
//     whose distance from src is not -1.
//
// Steps:
//  1. Set every distance to -1 and dist[src] to zero.
//  2. Perform up to V-1 passes. In each pass, examine every edge u->v in
//     the graph. If dist[u] is not -1 and relaxation improves dist[v],
//     update it and record that the pass changed something. If a pass
//     changes nothing, stop early; all distances are final.
//  3. Perform one more pass over every edge. If any edge can still be
//     relaxed, a negative cycle is reachable from src; return the
//     distance array with false.
//  4. If no edge can be relaxed, return the distance array with true.
//
// Notes:
//   - Unreachable vertices keep distance -1.
//   - A negative cycle that is not reachable from src does not affect the
//     result: the extra pass only relaxes edges whose source has a known
//     distance.
//   - When false is returned, the distances are not meaningful.
//   - src must lie in 0..N-1.
//
// Expected: O(V*E) time, O(V) space
// Input: not mutated
func BellmanFord(g *graph.Graph, src int) ([]int, bool) {
	// TODO: implement
	return nil, false
}
