// Package graph provides the adjacency-list graph type shared by the
// graphs/ exercises. It is a support layer: the exercises implement the
// algorithms, not the representation.
package graph

// Edge is one adjacency entry: the destination vertex and the weight
// (or capacity) of the edge.
type Edge struct {
	To     int
	Weight int
}

// TreeEdge is one edge of a spanning tree, reported with both endpoints.
// U is always the smaller endpoint.
type TreeEdge struct {
	U, V, W int
}

// Graph is a directed adjacency-list graph with vertices numbered
// 0..N-1. An undirected graph stores every edge in both adjacency
// lists; AddUndirected does that.
type Graph struct {
	N   int
	Adj [][]Edge
}

// NewGraph returns a graph with n vertices and no edges.
func NewGraph(n int) *Graph {
	if n < 0 {
		n = 0
	}
	return &Graph{N: n, Adj: make([][]Edge, n)}
}

// AddEdge adds the directed edge u->v with the given weight. Endpoints
// outside 0..N-1 are ignored.
func (g *Graph) AddEdge(u, v, w int) {
	if u < 0 || u >= g.N || v < 0 || v >= g.N {
		return
	}
	g.Adj[u] = append(g.Adj[u], Edge{To: v, Weight: w})
}

// AddUndirected adds both u->v and v->u with the same weight.
func (g *Graph) AddUndirected(u, v, w int) {
	g.AddEdge(u, v, w)
	g.AddEdge(v, u, w)
}
