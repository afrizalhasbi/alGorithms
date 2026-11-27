package graphs

import "alGorithms/internal/graph"

// [easy] DFS walks a graph depth first from a source vertex and returns
// the preorder visit sequence.
//
// Background. Depth-first search follows one edge as far as possible
// before backing up. When a vertex has no unvisited neighbor left, the
// search returns to the vertex it came from and tries that vertex's next
// neighbor. A stack, explicit or provided by recursion, holds the
// vertices on the current path. The preorder sequence records a vertex
// when the search first enters it, before any of its neighbors.
//
// Definitions:
//   - visited: a vertex is visited when the search first enters it. A
//     visited vertex is never entered again.
//   - preorder: the sequence of vertices in the order they are first
//     entered.
//   - adjacency order: the order of outgoing edges in Adj[u]. It is the
//     order in which the neighbors of u are tried.
//   - backtrack: returning from a vertex to the vertex that entered it,
//     after all of its neighbors have been tried.
//
// Steps:
//  1. Mark src visited and append src to the preorder sequence.
//  2. For each outgoing edge src->v in adjacency order, if v is not
//     visited, enter v and repeat the same procedure from v.
//  3. After all edges of a vertex have been tried, backtrack. The search
//     ends when it backtracks out of src.
//
// Notes:
//   - The graph is used as stored. A directed graph follows edge
//     direction only; an undirected graph built with AddUndirected has
//     edges in both directions.
//   - Only vertices reachable from src appear in the result.
//   - src must lie in 0..N-1.
//   - Recursion depth can reach V on a path-shaped graph.
//
// Expected: O(V+E) time, O(V) space
// Input: not mutated
func DFS(g *graph.Graph, src int) []int {
	// TODO: implement
	return nil
}
