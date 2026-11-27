package graphs

import "alGorithms/internal/graph"

// [medium] FloydWarshall returns the shortest-path distance between
// every ordered pair of vertices in a directed graph that may contain
// negative edge weights.
//
// Background. The algorithm builds shortest paths by allowing more and
// more intermediate vertices. Let d[i][j] be the best known distance
// from i to j. Before any intermediate vertex is allowed, d holds the
// direct edges. After allowing vertices 0..k as intermediates, every
// path that uses only those intermediates has been considered. Allowing
// vertex k as the next intermediate adds the possibility of going from
// i to k and then from k to j, so the update is
// d[i][j] = min(d[i][j], d[i][k]+d[k][j]). After all vertices have been
// allowed, d holds the shortest distances.
//
// Definitions:
//   - d[i][j]: the best known distance from i to j. The diagonal is zero.
//   - direct edge: an edge i->j. When several edges join the same pair,
//     the smallest weight is the direct distance.
//   - intermediate vertex: a vertex allowed to appear in the middle of a
//     path.
//   - unreachable: no path exists. It is represented by -1 in the result
//     and by a separate infinity value during the computation.
//   - negative cycle: a directed cycle whose total weight is negative.
//
// Steps:
//  1. Create an N by N table. Set d[i][i] to zero for every i. For every
//     edge u->v with weight w, set d[u][v] to the smaller of its current
//     value and w.
//  2. For k from 0 to N-1 in increasing order, for every i, and for
//     every j: if both d[i][k] and d[k][j] are reachable, the candidate
//     distance is their sum. If j is currently unreachable or the
//     candidate is smaller, set d[i][j] to the candidate.
//  3. After all k have been processed, replace every remaining
//     unreachable entry with -1 and return the table.
//
// Notes:
//   - Negative edges are allowed. A negative cycle makes the shortest
//     distance undefined for pairs that can reach and leave it; behavior
//     on such graphs is outside the contract.
//   - The diagonal is zero even when the graph has a negative cycle.
//   - The result is indexed as result[i][j].
//
// Expected: O(V^3) time, O(V^2) space
// Input: not mutated
func FloydWarshall(g *graph.Graph) [][]int {
	// TODO: implement
	return nil
}
