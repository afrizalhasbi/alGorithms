package graphs

import "alGorithms/internal/graph"

// [hard] EdmondsKarp returns the maximum flow from a source to a sink in
// a directed graph whose edge weights are capacities.
//
// Background. A flow assigns a non-negative amount to every edge, never
// more than the edge's capacity, and conserves it at every vertex except
// the source and the sink: what enters equals what leaves. The value of
// the flow is the net amount leaving the source. The residual graph
// records how much capacity remains on each edge and how much flow can be
// canceled by pushing back along an edge. Any path from source to sink in
// the residual graph can carry more flow, by the smallest residual
// capacity on the path. When no such path exists, the flow is maximum,
// and the set of vertices reachable from the source in the residual graph
// forms a minimum cut. Edmonds-Karp chooses the augmenting path with a
// breadth-first search, which bounds the number of augmentations.
//
// Definitions:
//   - capacity: capacity[u][v] is the total capacity of all edges u->v.
//     Parallel edges add their capacities.
//   - flow: flow[u][v] is the amount currently sent along u->v.
//   - residual capacity: capacity[u][v] minus flow[u][v] in the forward
//     direction, plus flow[v][u] in the backward direction. It is the
//     amount that can still be pushed from u to v.
//   - augmenting path: a path from source to sink in which every edge has
//     positive residual capacity.
//   - bottleneck: the smallest residual capacity on an augmenting path.
//   - maximum flow: a flow of the largest possible value.
//
// Steps:
//  1. Build a capacity table. For every edge u->v with capacity c, add c
//     to capacity[u][v]. If s or t is outside 0..N-1, or s equals t,
//     return zero.
//  2. Start with zero flow on every edge and a total flow of zero.
//  3. Search from s with a breadth-first search that follows only edges
//     with positive residual capacity. Record, for every reached vertex,
//     the edge by which it was reached. If t is not reached, stop.
//  4. Recover the augmenting path by following the recorded edges back
//     from t to s. Find its bottleneck, the smallest residual capacity on
//     the path.
//  5. Increase the flow along every edge of the path by the bottleneck,
//     decrease the flow along every reverse edge by the same amount, and
//     add the bottleneck to the total flow.
//  6. Repeat from step 3. When t is unreachable, return the total flow.
//
// Notes:
//   - Capacities must be non-negative integers.
//   - The graph is directed. An undirected edge with capacity c is
//     represented by two directed edges, each with capacity c.
//   - The capacity table uses O(V^2) space regardless of the number of
//     edges.
//
// Expected: O(V*E^2) time, O(V^2) space
// Input: not mutated
func EdmondsKarp(g *graph.Graph, s, t int) int {
	// TODO: implement
	return 0
}
