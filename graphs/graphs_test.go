package graphs

import (
	"fmt"
	"math/bits"
	"math/rand"
	"slices"
	"testing"

	"alGorithms/internal/complexity"
	"alGorithms/internal/graph"
)

// ---- shared helpers ----

func assertExponent(t *testing.T, name string, p, lo, hi float64) {
	t.Helper()
	if p > hi {
		t.Errorf("%s: measured exponent %.2f, expected at most %.2f", name, p, hi)
	} else if p < lo {
		t.Logf("%s: exponent %.2f below expected range [%.2f, %.2f] — faster than expected, ok", name, p, lo, hi)
	} else {
		t.Logf("%s: exponent %.2f ok (expected [%.2f, %.2f])", name, p, lo, hi)
	}
}

func okPrint(format string, args ...any) {
	fmt.Printf("ok: "+format+"\n", args...)
}

func addEdges(g *graph.Graph, edges [][3]int, undirected bool) {
	for _, e := range edges {
		if undirected {
			g.AddUndirected(e[0], e[1], e[2])
		} else {
			g.AddEdge(e[0], e[1], e[2])
		}
	}
}

// buildUndirected returns a connected undirected graph: a random spanning
// tree plus extra random edges. Weights are in 1..maxW.
func buildUndirected(n, extra, maxW int, seed int64) *graph.Graph {
	g := graph.NewGraph(n)
	if n == 0 {
		return g
	}
	rng := rand.New(rand.NewSource(seed))
	for v := 1; v < n; v++ {
		g.AddUndirected(rng.Intn(v), v, 1+rng.Intn(maxW))
	}
	for i := 0; i < extra; i++ {
		u, v := rng.Intn(n), rng.Intn(n)
		if u == v {
			continue
		}
		g.AddUndirected(u, v, 1+rng.Intn(maxW))
	}
	return g
}

// buildSimpleUndirected is buildUndirected without parallel edges.
func buildSimpleUndirected(n, extra, maxW int, seed int64) *graph.Graph {
	g := graph.NewGraph(n)
	if n == 0 {
		return g
	}
	rng := rand.New(rand.NewSource(seed))
	seen := make(map[[2]int]bool)
	add := func(u, v, w int) {
		if u == v {
			return
		}
		if u > v {
			u, v = v, u
		}
		k := [2]int{u, v}
		if seen[k] {
			return
		}
		seen[k] = true
		g.AddUndirected(u, v, w)
	}
	for v := 1; v < n; v++ {
		add(rng.Intn(v), v, 1+rng.Intn(maxW))
	}
	for i := 0; i < extra; i++ {
		add(rng.Intn(n), rng.Intn(n), 1+rng.Intn(maxW))
	}
	return g
}

// buildDAG returns a directed acyclic graph with edges only from a
// smaller index to a larger one.
func buildDAG(n, extra, maxW int, seed int64) *graph.Graph {
	g := graph.NewGraph(n)
	rng := rand.New(rand.NewSource(seed))
	for v := 1; v < n; v++ {
		g.AddEdge(rng.Intn(v), v, 1+rng.Intn(maxW))
	}
	for i := 0; i < extra; i++ {
		u, v := rng.Intn(n), rng.Intn(n)
		if u == v {
			continue
		}
		if u > v {
			u, v = v, u
		}
		g.AddEdge(u, v, 1+rng.Intn(maxW))
	}
	return g
}

// buildSignedDAG is buildDAG with weights in -5..10.
func buildSignedDAG(n, extra int, seed int64) *graph.Graph {
	g := graph.NewGraph(n)
	rng := rand.New(rand.NewSource(seed))
	for v := 1; v < n; v++ {
		g.AddEdge(rng.Intn(v), v, rng.Intn(16)-5)
	}
	for i := 0; i < extra; i++ {
		u, v := rng.Intn(n), rng.Intn(n)
		if u == v {
			continue
		}
		if u > v {
			u, v = v, u
		}
		g.AddEdge(u, v, rng.Intn(16)-5)
	}
	return g
}

// buildDirected returns a directed graph with extra random edges and
// weights in 1..maxW.
func buildDirected(n, extra, maxW int, seed int64) *graph.Graph {
	g := graph.NewGraph(n)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < extra; i++ {
		g.AddEdge(rng.Intn(n), rng.Intn(n), 1+rng.Intn(maxW))
	}
	return g
}

// buildBipartite returns a connected bipartite undirected graph. The two
// sides are 0..half-1 and half..n-1.
func buildBipartite(n, extra int, seed int64) *graph.Graph {
	g := graph.NewGraph(n)
	if n == 0 {
		return g
	}
	half := n / 2
	if half == 0 {
		return g
	}
	rng := rand.New(rand.NewSource(seed))
	for v := 1; v < n; v++ {
		u := v - 1
		if (u < half) == (v < half) {
			u = v - 2
		}
		if u < 0 {
			if v < half {
				u = half
			} else {
				u = 0
			}
		}
		g.AddUndirected(u, v, 1)
	}
	for i := 0; i < extra; i++ {
		a := rng.Intn(half)
		b := half + rng.Intn(n-half)
		g.AddUndirected(a, b, 1)
	}
	return g
}

type wedge struct{ u, v, w int }

// uniqueEdges returns every undirected edge once, with u < v. Self-loops
// are skipped.
func uniqueEdges(g *graph.Graph) []wedge {
	var out []wedge
	for u := 0; u < g.N; u++ {
		for _, e := range g.Adj[u] {
			if u < e.To {
				out = append(out, wedge{u, e.To, e.Weight})
			}
		}
	}
	return out
}

// refAllPairs is a test-side Floyd-Warshall. It returns -1 for
// unreachable pairs.
func refAllPairs(g *graph.Graph) [][]int {
	const inf = 1 << 60
	d := make([][]int, g.N)
	for i := range d {
		d[i] = make([]int, g.N)
		for j := range d[i] {
			d[i][j] = inf
		}
		d[i][i] = 0
	}
	for u := 0; u < g.N; u++ {
		for _, e := range g.Adj[u] {
			if e.Weight < d[u][e.To] {
				d[u][e.To] = e.Weight
			}
		}
	}
	for k := 0; k < g.N; k++ {
		for i := 0; i < g.N; i++ {
			if d[i][k] == inf {
				continue
			}
			for j := 0; j < g.N; j++ {
				if d[k][j] == inf {
					continue
				}
				if nd := d[i][k] + d[k][j]; nd < d[i][j] {
					d[i][j] = nd
				}
			}
		}
	}
	out := make([][]int, g.N)
	for i := range out {
		out[i] = make([]int, g.N)
		for j := range out[i] {
			if d[i][j] == inf {
				out[i][j] = -1
			} else {
				out[i][j] = d[i][j]
			}
		}
	}
	return out
}

// refBellmanFord is a test-side Bellman-Ford. It returns -1 for
// unreachable vertices and assumes no reachable negative cycle.
func refBellmanFord(g *graph.Graph, src int) []int {
	const inf = 1 << 60
	d := make([]int, g.N)
	for i := range d {
		d[i] = inf
	}
	d[src] = 0
	for pass := 0; pass < g.N-1; pass++ {
		changed := false
		for u := 0; u < g.N; u++ {
			if d[u] == inf {
				continue
			}
			for _, e := range g.Adj[u] {
				if nd := d[u] + e.Weight; nd < d[e.To] {
					d[e.To] = nd
					changed = true
				}
			}
		}
		if !changed {
			break
		}
	}
	out := make([]int, g.N)
	for i := range out {
		if d[i] == inf {
			out[i] = -1
		} else {
			out[i] = d[i]
		}
	}
	return out
}

// bruteShortest enumerates simple paths from s to t. It is used only on
// tiny graphs without negative cycles.
func bruteShortest(g *graph.Graph, s, t int) int {
	const inf = 1 << 60
	best := inf
	seen := make([]bool, g.N)
	var dfs func(u, dist int)
	dfs = func(u, dist int) {
		if u == t {
			if dist < best {
				best = dist
			}
			return
		}
		seen[u] = true
		for _, e := range g.Adj[u] {
			if !seen[e.To] {
				dfs(e.To, dist+e.Weight)
			}
		}
		seen[u] = false
	}
	dfs(s, 0)
	if best == inf {
		return -1
	}
	return best
}

func reaches(g *graph.Graph, s, t int) bool {
	seen := make([]bool, g.N)
	stack := []int{s}
	seen[s] = true
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if u == t {
			return true
		}
		for _, e := range g.Adj[u] {
			if !seen[e.To] {
				seen[e.To] = true
				stack = append(stack, e.To)
			}
		}
	}
	return false
}

func hasCycleRef(g *graph.Graph) bool {
	for v := 0; v < g.N; v++ {
		for _, e := range g.Adj[v] {
			if e.To == v || reaches(g, e.To, v) {
				return true
			}
		}
	}
	return false
}

// componentCount counts components after optionally removing one vertex
// and one undirected edge. skipV < 0 means no vertex is removed;
// skipEdge[0] < 0 means no edge is removed.
func componentCount(g *graph.Graph, skipV int, skipEdge [2]int) int {
	seen := make([]bool, g.N)
	count := 0
	for s := 0; s < g.N; s++ {
		if s == skipV || seen[s] {
			continue
		}
		count++
		seen[s] = true
		stack := []int{s}
		for len(stack) > 0 {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, e := range g.Adj[u] {
				if e.To == skipV {
					continue
				}
				if skipEdge[0] >= 0 &&
					((u == skipEdge[0] && e.To == skipEdge[1]) ||
						(u == skipEdge[1] && e.To == skipEdge[0])) {
					continue
				}
				if !seen[e.To] {
					seen[e.To] = true
					stack = append(stack, e.To)
				}
			}
		}
	}
	return count
}

func validTopo(g *graph.Graph, order []int) bool {
	if len(order) != g.N {
		return false
	}
	pos := make([]int, g.N)
	for i, v := range order {
		if v < 0 || v >= g.N {
			return false
		}
		pos[v] = i
	}
	for u := 0; u < g.N; u++ {
		for _, e := range g.Adj[u] {
			if pos[u] >= pos[e.To] {
				return false
			}
		}
	}
	return true
}

func normalizeComponents(comps [][]int) [][]int {
	out := make([][]int, len(comps))
	for i, c := range comps {
		out[i] = slices.Clone(c)
		slices.Sort(out[i])
	}
	slices.SortFunc(out, func(a, b []int) int { return slices.Compare(a, b) })
	return out
}

func equalComponents(a, b [][]int) bool {
	return slices.EqualFunc(a, b, func(x, y []int) bool { return slices.Equal(x, y) })
}

// sccByReach groups vertices by mutual reachability. It is the oracle
// for TarjanSCC on small graphs.
func sccByReach(g *graph.Graph) [][]int {
	assigned := make([]bool, g.N)
	var out [][]int
	for i := 0; i < g.N; i++ {
		if assigned[i] {
			continue
		}
		group := []int{i}
		assigned[i] = true
		for j := i + 1; j < g.N; j++ {
			if !assigned[j] && reaches(g, i, j) && reaches(g, j, i) {
				assigned[j] = true
				group = append(group, j)
			}
		}
		out = append(out, group)
	}
	return out
}

func bruteMST(g *graph.Graph) (int, bool) {
	n := g.N
	if n <= 1 {
		return 0, true
	}
	edges := uniqueEdges(g)
	const inf = 1 << 60
	best := inf
	for mask := 0; mask < 1<<len(edges); mask++ {
		if bits.OnesCount(uint(mask)) != n-1 {
			continue
		}
		parent := make([]int, n)
		for i := range parent {
			parent[i] = i
		}
		var find func(int) int
		find = func(x int) int {
			for parent[x] != x {
				parent[x] = parent[parent[x]]
				x = parent[x]
			}
			return x
		}
		sum := 0
		ok := true
		for i, e := range edges {
			if mask>>i&1 == 0 {
				continue
			}
			ra, rb := find(e.u), find(e.v)
			if ra == rb {
				ok = false
				break
			}
			parent[ra] = rb
			sum += e.w
		}
		if !ok {
			continue
		}
		r := find(0)
		for v := 1; v < n; v++ {
			if find(v) != r {
				ok = false
				break
			}
		}
		if ok && sum < best {
			best = sum
		}
	}
	if best == inf {
		return 0, false
	}
	return best, true
}

func refKruskal(g *graph.Graph) int {
	edges := uniqueEdges(g)
	slices.SortFunc(edges, func(a, b wedge) int {
		if a.w != b.w {
			return a.w - b.w
		}
		if a.u != b.u {
			return a.u - b.u
		}
		return a.v - b.v
	})
	parent := make([]int, g.N)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	total := 0
	for _, e := range edges {
		ra, rb := find(e.u), find(e.v)
		if ra != rb {
			parent[ra] = rb
			total += e.w
		}
	}
	return total
}

func refPrim(g *graph.Graph, src int) int {
	if src < 0 || src >= g.N {
		return 0
	}
	const inf = 1 << 60
	in := make([]bool, g.N)
	in[src] = true
	total := 0
	for {
		best := inf
		next := -1
		for u := 0; u < g.N; u++ {
			if !in[u] {
				continue
			}
			for _, e := range g.Adj[u] {
				if !in[e.To] && e.Weight < best {
					best = e.Weight
					next = e.To
				}
			}
		}
		if next == -1 {
			break
		}
		in[next] = true
		total += best
	}
	return total
}

func minCutBrute(g *graph.Graph, s, t int) int {
	if s == t || s < 0 || t < 0 || s >= g.N || t >= g.N {
		return 0
	}
	const inf = 1 << 60
	best := inf
	for mask := 0; mask < 1<<g.N; mask++ {
		if mask>>s&1 == 0 || mask>>t&1 == 1 {
			continue
		}
		cut := 0
		for u := 0; u < g.N; u++ {
			if mask>>u&1 == 0 {
				continue
			}
			for _, e := range g.Adj[u] {
				if mask>>e.To&1 == 0 {
					cut += e.Weight
				}
			}
		}
		if cut < best {
			best = cut
		}
	}
	return best
}

func validateTree(t *testing.T, g *graph.Graph, edges []graph.TreeEdge, wantWeight int) {
	t.Helper()
	parent := make([]int, g.N)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		for parent[x] != x {
			parent[x] = parent[parent[x]]
			x = parent[x]
		}
		return x
	}
	sum := 0
	for _, e := range edges {
		if e.U < 0 || e.V >= g.N || e.U >= e.V {
			t.Errorf("tree edge %+v is not normalized or in range", e)
			continue
		}
		ra, rb := find(e.U), find(e.V)
		if ra == rb {
			t.Errorf("tree edge %+v closes a cycle", e)
		}
		parent[ra] = rb
		sum += e.W
	}
	if sum != wantWeight {
		t.Errorf("tree edge weights sum to %d, want %d", sum, wantWeight)
	}
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(6)
	for i := 0; i < 6; i++ {
		if !uf.Connected(i, i) {
			t.Errorf("Connected(%d, %d) = false, want true", i, i)
		}
	}
	unions := [][2]int{{0, 1}, {2, 3}, {1, 3}, {5, 5}, {0, 3}}
	want := []bool{true, true, true, false, false}
	for i, u := range unions {
		if got := uf.Union(u[0], u[1]); got != want[i] {
			t.Errorf("Union(%d, %d) = %v, want %v", u[0], u[1], got, want[i])
		} else {
			okPrint("UnionFind Union(%d, %d)/%v/%v", u[0], u[1], got, want[i])
		}
	}
	pairs := [][2]int{{0, 2}, {1, 2}, {0, 3}, {0, 4}, {4, 5}}
	wantPairs := []bool{true, true, true, false, false}
	for i, p := range pairs {
		if got := uf.Connected(p[0], p[1]); got != wantPairs[i] {
			t.Errorf("Connected(%d, %d) = %v, want %v", p[0], p[1], got, wantPairs[i])
		}
	}
	okPrint("UnionFind 5 unions and 5 connectivity probes")

	// Random union sequence checked against naive labels.
	const n = 100
	rng := rand.New(rand.NewSource(7))
	uf2 := NewUnionFind(n)
	label := make([]int, n)
	for i := range label {
		label[i] = i
	}
	for step := 0; step < 500; step++ {
		a, b := rng.Intn(n), rng.Intn(n)
		wantMerged := label[a] != label[b]
		if got := uf2.Union(a, b); got != wantMerged {
			t.Fatalf("Union(%d, %d) = %v, want %v", a, b, got, wantMerged)
		}
		if wantMerged {
			old, newID := label[b], label[a]
			for i := range label {
				if label[i] == old {
					label[i] = newID
				}
			}
		}
		x, y := rng.Intn(n), rng.Intn(n)
		if got, want := uf2.Connected(x, y), label[x] == label[y]; got != want {
			t.Fatalf("Connected(%d, %d) = %v, want %v", x, y, got, want)
		}
	}
	okPrint("UnionFind 500 random unions against naive labels")

	// Complexity: a fixed number of Find calls on a cached structure.
	cache := make(map[int]*UnionFind)
	build := func(n int) *UnionFind {
		if u, ok := cache[n]; ok {
			return u
		}
		u := NewUnionFind(n)
		for i := 1; i < n; i++ {
			u.Union(i, i/2)
		}
		cache[n] = u
		return u
	}
	const probes = 100_000
	assertExponent(t, "O(alpha(n))", complexity.FitExponent(1000, 3, func(n int) {
		u := build(n)
		for i := 0; i < probes; i++ {
			u.Find((i * 7919) % n)
		}
	}), -0.3, 0.4)
}

func TestBFS(t *testing.T) {
	g := graph.NewGraph(6)
	addEdges(g, [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 3, 1}, {3, 4, 1}}, false)
	order, dist := BFS(g, 0)
	if len(dist) != 6 {
		t.Fatalf("BFS returned %d distances, want 6", len(dist))
	}
	if want := []int{0, 1, 1, 2, 3, -1}; !slices.Equal(dist, want) {
		t.Errorf("BFS dist = %v, want %v", dist, want)
	} else {
		okPrint("BFS distances/%v/%v", dist, want)
	}
	if want := []int{0, 1, 2, 3, 4}; !slices.Equal(order, want) {
		t.Errorf("BFS order = %v, want %v", order, want)
	} else {
		okPrint("BFS order/%v/%v", order, want)
	}

	for _, n := range []int{10, 20, 30} {
		rg := buildDirected(n, 3*n, 1, int64(n))
		ref := refAllPairs(rg)
		for src := 0; src < n; src++ {
			_, got := BFS(rg, src)
			if len(got) != n {
				t.Fatalf("BFS(n=%d, src=%d) returned %d distances", n, src, len(got))
			}
			if !slices.Equal(got, ref[src]) {
				t.Fatalf("BFS(n=%d, src=%d) dist = %v, want %v", n, src, got, ref[src])
			}
		}
	}
	okPrint("BFS 3 random directed graphs, all sources/got=Floyd-Warshall/want=Floyd-Warshall")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 1, 11)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		BFS(build(n), 0)
	}), 0.7, 1.4)
}

func TestDFS(t *testing.T) {
	g := graph.NewGraph(5)
	addEdges(g, [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 3, 1}, {3, 4, 1}}, false)
	got := DFS(g, 0)
	want := []int{0, 1, 3, 4, 2}
	if !slices.Equal(got, want) {
		t.Errorf("DFS order = %v, want %v", got, want)
	} else {
		okPrint("DFS order/%v/%v", got, want)
	}

	for _, n := range []int{10, 20, 30} {
		rg := buildDirected(n, 3*n, 1, int64(n))
		ref := refAllPairs(rg)
		for src := 0; src < n; src++ {
			order := DFS(rg, src)
			if len(order) == 0 || order[0] != src {
				t.Fatalf("DFS(n=%d, src=%d) starts with %v, want %d", n, src, order, src)
			}
			seen := make(map[int]bool, len(order))
			for _, v := range order {
				if seen[v] {
					t.Fatalf("DFS(n=%d, src=%d) visits %d twice", n, src, v)
				}
				seen[v] = true
			}
			for v := 0; v < n; v++ {
				if seen[v] != (ref[src][v] != -1) {
					t.Fatalf("DFS(n=%d, src=%d): visited[%d] = %v, reachable = %v",
						n, src, v, seen[v], ref[src][v] != -1)
				}
			}
		}
	}
	okPrint("DFS 3 random directed graphs, all sources/got=reachable set/want=Floyd-Warshall")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 1, 12)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		DFS(build(n), 0)
	}), 0.7, 1.4)
}

func TestConnectedComponents(t *testing.T) {
	g := graph.NewGraph(6)
	addEdges(g, [][3]int{{0, 1, 1}, {1, 2, 1}, {3, 4, 1}}, true)
	count, label := ConnectedComponents(g)
	if count != 3 {
		t.Errorf("ConnectedComponents count = %d, want 3", count)
	}
	if want := []int{0, 0, 0, 1, 1, 2}; !slices.Equal(label, want) {
		t.Errorf("ConnectedComponents labels = %v, want %v", label, want)
	} else {
		okPrint("ConnectedComponents labels/%v/%v", label, want)
	}

	for _, n := range []int{5, 12, 30} {
		rg := buildUndirected(n, 2*n, 1, int64(n))
		gotCount, gotLabel := ConnectedComponents(rg)
		if len(gotLabel) != n {
			t.Fatalf("ConnectedComponents(n=%d) returned %d labels", n, len(gotLabel))
		}
		for i := 0; i < n; i++ {
			if gotLabel[i] < 0 || gotLabel[i] >= gotCount {
				t.Fatalf("label[%d] = %d outside 0..%d", i, gotLabel[i], gotCount-1)
			}
		}
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				connected := reaches(rg, i, j) && reaches(rg, j, i)
				if (gotLabel[i] == gotLabel[j]) != connected {
					t.Fatalf("n=%d: labels[%d]=%d labels[%d]=%d but connected=%v",
						n, i, gotLabel[i], j, gotLabel[j], connected)
				}
			}
		}
	}
	okPrint("ConnectedComponents 3 random graphs/got=same label iff reachable/want=reachability")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 1, 13)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		ConnectedComponents(build(n))
	}), 0.7, 1.4)
}

func TestIsBipartite(t *testing.T) {
	even := graph.NewGraph(4)
	addEdges(even, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {3, 0, 1}}, true)
	if ok, colors := IsBipartite(even); !ok || len(colors) != 4 {
		t.Errorf("IsBipartite(even cycle) = (%v, %v), want (true, 4 colors)", ok, colors)
	} else {
		okPrint("IsBipartite even cycle/%v/true", ok)
	}

	odd := graph.NewGraph(3)
	addEdges(odd, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}}, true)
	if ok, _ := IsBipartite(odd); ok {
		t.Error("IsBipartite(triangle) = true, want false")
	} else {
		okPrint("IsBipartite triangle/%v/false", ok)
	}

	disconnected := graph.NewGraph(5)
	addEdges(disconnected, [][3]int{{0, 1, 1}, {2, 3, 1}, {3, 4, 1}, {4, 2, 1}}, true)
	if ok, _ := IsBipartite(disconnected); ok {
		t.Error("IsBipartite(disconnected with odd cycle) = true, want false")
	} else {
		okPrint("IsBipartite disconnected odd cycle/%v/false", ok)
	}

	selfLoop := graph.NewGraph(2)
	addEdges(selfLoop, [][3]int{{0, 0, 1}, {0, 1, 1}}, true)
	if ok, _ := IsBipartite(selfLoop); ok {
		t.Error("IsBipartite(self-loop) = true, want false")
	} else {
		okPrint("IsBipartite self-loop/%v/false", ok)
	}

	isolated := graph.NewGraph(3)
	if ok, colors := IsBipartite(isolated); !ok || len(colors) != 3 {
		t.Errorf("IsBipartite(isolated) = (%v, %v), want (true, 3 colors)", ok, colors)
	} else {
		okPrint("IsBipartite isolated vertices/%v/true", ok)
	}

	for _, n := range []int{1, 2, 5, 10} {
		rg := buildSimpleUndirected(n, 2*n, 1, int64(n))
		got, colors := IsBipartite(rg)
		want := isBipartiteBrute(rg)
		if got != want {
			t.Fatalf("IsBipartite(n=%d) = %v, want %v", n, got, want)
		}
		if got {
			if len(colors) != n {
				t.Fatalf("IsBipartite(n=%d) returned %d colors", n, len(colors))
			}
			for u := 0; u < n; u++ {
				for _, e := range rg.Adj[u] {
					if colors[u] == colors[e.To] {
						t.Fatalf("IsBipartite(n=%d): edge %d-%d has equal colors", n, u, e.To)
					}
				}
			}
		}
	}
	okPrint("IsBipartite 4 random graphs/got=brute force 2-coloring/want=brute force")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildBipartite(n, 2*n, 14)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		IsBipartite(build(n))
	}), 0.7, 1.4)
}

func isBipartiteBrute(g *graph.Graph) bool {
	for mask := 0; mask < 1<<g.N; mask++ {
		ok := true
		for u := 0; u < g.N && ok; u++ {
			for _, e := range g.Adj[u] {
				if (mask >> u & 1) == (mask >> e.To & 1) {
					ok = false
					break
				}
			}
		}
		if ok {
			return true
		}
	}
	return false
}

func TestCycleDetectionDirected(t *testing.T) {
	dag := graph.NewGraph(4)
	addEdges(dag, [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 3, 1}}, false)
	if HasCycleDirected(dag) {
		t.Error("HasCycleDirected(diamond DAG) = true, want false")
	} else {
		okPrint("HasCycleDirected diamond DAG/%v/false", false)
	}

	triangle := graph.NewGraph(3)
	addEdges(triangle, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}}, false)
	if !HasCycleDirected(triangle) {
		t.Error("HasCycleDirected(triangle) = false, want true")
	} else {
		okPrint("HasCycleDirected triangle/%v/true", true)
	}

	selfLoop := graph.NewGraph(1)
	selfLoop.AddEdge(0, 0, 1)
	if !HasCycleDirected(selfLoop) {
		t.Error("HasCycleDirected(self-loop) = false, want true")
	} else {
		okPrint("HasCycleDirected self-loop/%v/true", true)
	}

	disconnected := graph.NewGraph(4)
	addEdges(disconnected, [][3]int{{0, 1, 1}, {2, 3, 1}, {3, 2, 1}}, false)
	if !HasCycleDirected(disconnected) {
		t.Error("HasCycleDirected(disconnected cycle) = false, want true")
	} else {
		okPrint("HasCycleDirected disconnected cycle/%v/true", true)
	}

	for _, n := range []int{1, 5, 12, 25} {
		rg := buildDirected(n, 2*n, 1, int64(n))
		if got, want := HasCycleDirected(rg), hasCycleRef(rg); got != want {
			t.Fatalf("HasCycleDirected(n=%d) = %v, want %v", n, got, want)
		}
	}
	okPrint("HasCycleDirected 4 random graphs/got=reachability oracle/want=reachability oracle")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildDAG(n, 2*n, 1, 15)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		HasCycleDirected(build(n))
	}), 0.7, 1.4)
}

func TestCycleDetectionUndirected(t *testing.T) {
	tree := graph.NewGraph(4)
	addEdges(tree, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, true)
	if HasCycleUndirected(tree) {
		t.Error("HasCycleUndirected(path) = true, want false")
	} else {
		okPrint("HasCycleUndirected path/%v/false", false)
	}

	triangle := graph.NewGraph(3)
	addEdges(triangle, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}}, true)
	if !HasCycleUndirected(triangle) {
		t.Error("HasCycleUndirected(triangle) = false, want true")
	} else {
		okPrint("HasCycleUndirected triangle/%v/true", true)
	}

	selfLoop := graph.NewGraph(1)
	selfLoop.AddUndirected(0, 0, 1)
	if !HasCycleUndirected(selfLoop) {
		t.Error("HasCycleUndirected(self-loop) = false, want true")
	} else {
		okPrint("HasCycleUndirected self-loop/%v/true", true)
	}

	disconnected := graph.NewGraph(5)
	addEdges(disconnected, [][3]int{{0, 1, 1}, {2, 3, 1}, {3, 4, 1}, {4, 2, 1}}, true)
	if !HasCycleUndirected(disconnected) {
		t.Error("HasCycleUndirected(disconnected cycle) = false, want true")
	} else {
		okPrint("HasCycleUndirected disconnected cycle/%v/true", true)
	}

	for _, n := range []int{1, 5, 12, 20} {
		rg := buildSimpleUndirected(n, 2*n, 1, int64(n))
		components := componentCount(rg, -1, [2]int{-1, -1})
		want := len(uniqueEdges(rg)) > n-components
		if got := HasCycleUndirected(rg); got != want {
			t.Fatalf("HasCycleUndirected(n=%d) = %v, want %v (E=%d, V-C=%d)",
				n, got, want, len(uniqueEdges(rg)), n-components)
		}
	}
	okPrint("HasCycleUndirected 4 random graphs/got=E > V-C/want=E > V-C")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildSimpleUndirected(n, 0, 1, 16)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		HasCycleUndirected(build(n))
	}), 0.7, 1.4)
}

func TestTopologicalSort(t *testing.T) {
	dag := graph.NewGraph(4)
	addEdges(dag, [][3]int{{0, 1, 1}, {0, 2, 1}, {1, 3, 1}, {2, 3, 1}}, false)
	order, ok := TopologicalSort(dag)
	if !ok || !validTopo(dag, order) {
		t.Errorf("TopologicalSort(diamond) = (%v, %v), want a valid order", order, ok)
	} else {
		okPrint("TopologicalSort diamond/%v/valid order", order)
	}

	cycle := graph.NewGraph(3)
	addEdges(cycle, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}}, false)
	if order, ok := TopologicalSort(cycle); ok {
		t.Errorf("TopologicalSort(cycle) = (%v, true), want false", order)
	} else {
		okPrint("TopologicalSort cycle/%v/false", ok)
	}

	selfLoop := graph.NewGraph(1)
	selfLoop.AddEdge(0, 0, 1)
	if _, ok := TopologicalSort(selfLoop); ok {
		t.Error("TopologicalSort(self-loop) = true, want false")
	} else {
		okPrint("TopologicalSort self-loop/%v/false", ok)
	}

	empty := graph.NewGraph(0)
	if order, ok := TopologicalSort(empty); !ok || len(order) != 0 {
		t.Errorf("TopologicalSort(empty) = (%v, %v), want ([], true)", order, ok)
	} else {
		okPrint("TopologicalSort empty/%v/true", ok)
	}

	for _, n := range []int{1, 5, 20, 30} {
		rg := buildDAG(n, 2*n, 1, int64(n))
		order, ok := TopologicalSort(rg)
		if !ok || !validTopo(rg, order) {
			t.Fatalf("TopologicalSort(n=%d) = (%v, %v), want a valid order", n, order, ok)
		}
	}
	okPrint("TopologicalSort 4 random DAGs/got=every edge points forward/want=valid order")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildDAG(n, 2*n, 1, 17)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(2000, 3, func(n int) {
		TopologicalSort(build(n))
	}), 0.7, 1.4)
}

func TestDijkstra(t *testing.T) {
	g := graph.NewGraph(6)
	addEdges(g, [][3]int{
		{0, 1, 4}, {0, 2, 1}, {2, 1, 2}, {1, 3, 1}, {2, 3, 5}, {3, 4, 3},
	}, false)
	got := Dijkstra(g, 0)
	want := []int{0, 3, 1, 4, 7, -1}
	if len(got) != len(want) || !slices.Equal(got, want) {
		t.Errorf("Dijkstra dist = %v, want %v", got, want)
	} else {
		okPrint("Dijkstra distances/%v/%v", got, want)
	}

	zero := graph.NewGraph(3)
	addEdges(zero, [][3]int{{0, 1, 0}, {1, 2, 0}}, false)
	if got := Dijkstra(zero, 0); !slices.Equal(got, []int{0, 0, 0}) {
		t.Errorf("Dijkstra(zero weights) = %v, want [0 0 0]", got)
	} else {
		okPrint("Dijkstra zero weights/%v/[0 0 0]", got)
	}

	for _, n := range []int{10, 15, 20} {
		rg := buildDirected(n, 3*n, 9, int64(n))
		ref := refAllPairs(rg)
		for src := 0; src < n; src++ {
			got := Dijkstra(rg, src)
			if len(got) != n {
				t.Fatalf("Dijkstra(n=%d, src=%d) returned %d distances", n, src, len(got))
			}
			if !slices.Equal(got, ref[src]) {
				t.Fatalf("Dijkstra(n=%d, src=%d) = %v, want %v", n, src, got, ref[src])
			}
		}
	}
	okPrint("Dijkstra 3 random graphs, all sources/got=Floyd-Warshall/want=Floyd-Warshall")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 9, 18)
		cache[n] = g
		return g
	}
	assertExponent(t, "O((V+E) log V)", complexity.FitExponent(1000, 3, func(n int) {
		Dijkstra(build(n), 0)
	}), 0.9, 1.5)
}

func TestBellmanFord(t *testing.T) {
	g := graph.NewGraph(4)
	addEdges(g, [][3]int{{0, 1, 4}, {0, 2, 5}, {1, 2, -3}, {2, 3, 2}}, false)
	got, ok := BellmanFord(g, 0)
	if !ok {
		t.Error("BellmanFord(negative DAG) reported a negative cycle")
	} else if want := []int{0, 4, 1, 3}; !slices.Equal(got, want) {
		t.Errorf("BellmanFord(negative DAG) = %v, want %v", got, want)
	} else {
		okPrint("BellmanFord negative DAG/%v/%v", got, want)
	}

	negCycle := graph.NewGraph(3)
	addEdges(negCycle, [][3]int{{0, 1, 1}, {1, 2, -2}, {2, 1, -2}}, false)
	if _, ok := BellmanFord(negCycle, 0); ok {
		t.Error("BellmanFord(reachable negative cycle) = true, want false")
	} else {
		okPrint("BellmanFord reachable negative cycle/%v/false", ok)
	}

	unreachableCycle := graph.NewGraph(4)
	addEdges(unreachableCycle, [][3]int{{0, 1, 1}, {2, 3, -2}, {3, 2, -2}}, false)
	got, ok = BellmanFord(unreachableCycle, 0)
	if !ok {
		t.Error("BellmanFord(unreachable negative cycle) = false, want true")
	} else if want := []int{0, 1, -1, -1}; !slices.Equal(got, want) {
		t.Errorf("BellmanFord(unreachable negative cycle) = %v, want %v", got, want)
	} else {
		okPrint("BellmanFord unreachable negative cycle/%v/%v", got, want)
	}

	for _, n := range []int{10, 15, 20} {
		rg := buildSignedDAG(n, 2*n, int64(n))
		ref := refAllPairs(rg)
		for src := 0; src < n; src++ {
			got, ok := BellmanFord(rg, src)
			if !ok {
				t.Fatalf("BellmanFord(signed DAG, n=%d, src=%d) reported a cycle", n, src)
			}
			if !slices.Equal(got, ref[src]) {
				t.Fatalf("BellmanFord(n=%d, src=%d) = %v, want %v", n, src, got, ref[src])
			}
		}
	}
	okPrint("BellmanFord 3 signed DAGs, all sources/got=Floyd-Warshall/want=Floyd-Warshall")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildDAG(n, 2*n, 9, 19)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V*E)", complexity.FitExponent(150, 3, func(n int) {
		BellmanFord(build(n), 0)
	}), 1.7, 2.3)
}

func TestFloydWarshall(t *testing.T) {
	g := graph.NewGraph(4)
	addEdges(g, [][3]int{{0, 1, 3}, {0, 2, 8}, {1, 2, -2}, {2, 3, 1}, {3, 0, 4}}, false)
	got := FloydWarshall(g)
	if len(got) != 4 {
		t.Fatalf("FloydWarshall returned %d rows, want 4", len(got))
	}
	for i := 0; i < 4; i++ {
		if len(got[i]) != 4 {
			t.Fatalf("FloydWarshall row %d has %d entries, want 4", i, len(got[i]))
		}
		for j := 0; j < 4; j++ {
			want := bruteShortest(g, i, j)
			if got[i][j] != want {
				t.Errorf("FloydWarshall[%d][%d] = %d, want %d", i, j, got[i][j], want)
			}
		}
	}
	okPrint("FloydWarshall 4x4 signed graph/got=brute simple paths/want=brute simple paths")

	unreachable := graph.NewGraph(3)
	addEdges(unreachable, [][3]int{{0, 1, 2}}, false)
	got = FloydWarshall(unreachable)
	if got[0][2] != -1 || got[2][0] != -1 || got[2][2] != 0 {
		t.Errorf("FloydWarshall(unreachable) = %v, want -1 off-component and 0 diagonal", got)
	} else {
		okPrint("FloydWarshall unreachable/%d/%d", got[0][2], -1)
	}

	parallel := graph.NewGraph(2)
	parallel.AddEdge(0, 1, 5)
	parallel.AddEdge(0, 1, 2)
	got = FloydWarshall(parallel)
	if got[0][1] != 2 {
		t.Errorf("FloydWarshall(parallel edges) = %d, want 2", got[0][1])
	} else {
		okPrint("FloydWarshall parallel edges/%d/%d", got[0][1], 2)
	}

	for _, n := range []int{5, 10, 15, 20} {
		rg := buildSignedDAG(n, 2*n, int64(n))
		got := FloydWarshall(rg)
		for i := 0; i < n; i++ {
			ref := refBellmanFord(rg, i)
			if !slices.Equal(got[i], ref) {
				t.Fatalf("FloydWarshall(n=%d) row %d = %v, want %v", n, i, got[i], ref)
			}
			if n <= 6 {
				for j := 0; j < n; j++ {
					if want := bruteShortest(rg, i, j); got[i][j] != want {
						t.Fatalf("FloydWarshall(n=%d)[%d][%d] = %d, want %d",
							n, i, j, got[i][j], want)
					}
				}
			}
		}
	}
	okPrint("FloydWarshall 4 signed DAGs/got=Bellman-Ford/want=Bellman-Ford")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 9, 20)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V^3)", complexity.FitExponent(20, 3, func(n int) {
		FloydWarshall(build(n))
	}), 2.6, 3.4)
}

func TestKruskal(t *testing.T) {
	g := graph.NewGraph(4)
	addEdges(g, [][3]int{
		{0, 1, 1}, {0, 2, 4}, {1, 2, 2}, {1, 3, 5}, {2, 3, 3},
	}, true)
	weight, edges := Kruskal(g)
	if weight != 6 || len(edges) != 3 {
		t.Errorf("Kruskal = (%d, %d edges), want (6, 3 edges)", weight, len(edges))
	} else {
		okPrint("Kruskal weight and edge count/%d, %d/6, 3", weight, len(edges))
	}
	validateTree(t, g, edges, 6)

	forest := graph.NewGraph(4)
	addEdges(forest, [][3]int{{0, 1, 1}, {2, 3, 2}}, true)
	weight, edges = Kruskal(forest)
	if weight != 3 || len(edges) != 2 {
		t.Errorf("Kruskal(disconnected) = (%d, %d edges), want (3, 2 edges)", weight, len(edges))
	} else {
		okPrint("Kruskal disconnected weight and edge count/%d, %d/3, 2", weight, len(edges))
	}
	validateTree(t, forest, edges, 3)

	for _, n := range []int{3, 4, 5, 6} {
		rg := buildSimpleUndirected(n, 2*n, 9, int64(n))
		want, ok := bruteMST(rg)
		if !ok {
			t.Fatalf("bruteMST(n=%d) found no spanning tree", n)
		}
		weight, edges := Kruskal(rg)
		if weight != want || len(edges) != n-1 {
			t.Fatalf("Kruskal(n=%d) = (%d, %d edges), want (%d, %d edges)",
				n, weight, len(edges), want, n-1)
		}
		validateTree(t, rg, edges, want)
	}
	okPrint("Kruskal 4 random graphs/got=brute force edge subsets/want=brute force")

	large := buildUndirected(50, 100, 9, 21)
	gotWeight, _ := Kruskal(large)
	wantWeight := refPrim(large, 0)
	if gotWeight != wantWeight {
		t.Errorf("Kruskal(50 vertices) weight = %d, want Prim weight %d", gotWeight, wantWeight)
	} else {
		okPrint("Kruskal 50 vertices/%d/%d", gotWeight, wantWeight)
	}

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 9, 22)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(E log E)", complexity.FitExponent(1000, 3, func(n int) {
		Kruskal(build(n))
	}), 0.9, 1.5)
}

func TestPrim(t *testing.T) {
	g := graph.NewGraph(4)
	addEdges(g, [][3]int{
		{0, 1, 1}, {0, 2, 4}, {1, 2, 2}, {1, 3, 5}, {2, 3, 3},
	}, true)
	weight, edges := Prim(g, 0)
	if weight != 6 || len(edges) != 3 {
		t.Errorf("Prim = (%d, %d edges), want (6, 3 edges)", weight, len(edges))
	} else {
		okPrint("Prim weight and edge count/%d, %d/6, 3", weight, len(edges))
	}
	validateTree(t, g, edges, 6)

	forest := graph.NewGraph(4)
	addEdges(forest, [][3]int{{0, 1, 1}, {2, 3, 2}}, true)
	weight, edges = Prim(forest, 0)
	if weight != 1 || len(edges) != 1 {
		t.Errorf("Prim(disconnected, src=0) = (%d, %d edges), want (1, 1 edge)", weight, len(edges))
	} else {
		okPrint("Prim disconnected src 0/%d, %d/1, 1", weight, len(edges))
	}
	weight, edges = Prim(forest, 2)
	if weight != 2 || len(edges) != 1 {
		t.Errorf("Prim(disconnected, src=2) = (%d, %d edges), want (2, 1 edge)", weight, len(edges))
	} else {
		okPrint("Prim disconnected src 2/%d, %d/2, 1", weight, len(edges))
	}
	if weight, edges := Prim(forest, 99); weight != 0 || len(edges) != 0 {
		t.Errorf("Prim(invalid src) = (%d, %d edges), want (0, 0 edges)", weight, len(edges))
	} else {
		okPrint("Prim invalid src/%d, %d/0, 0", weight, len(edges))
	}

	for _, n := range []int{3, 4, 5, 6} {
		rg := buildSimpleUndirected(n, 2*n, 9, int64(n))
		want, ok := bruteMST(rg)
		if !ok {
			t.Fatalf("bruteMST(n=%d) found no spanning tree", n)
		}
		weight, edges := Prim(rg, 0)
		if weight != want || len(edges) != n-1 {
			t.Fatalf("Prim(n=%d) = (%d, %d edges), want (%d, %d edges)",
				n, weight, len(edges), want, n-1)
		}
		validateTree(t, rg, edges, want)
	}
	okPrint("Prim 4 random graphs/got=brute force edge subsets/want=brute force")

	large := buildUndirected(50, 100, 9, 23)
	gotWeight, _ := Prim(large, 0)
	wantWeight := refKruskal(large)
	if gotWeight != wantWeight {
		t.Errorf("Prim(50 vertices) weight = %d, want Kruskal weight %d", gotWeight, wantWeight)
	} else {
		okPrint("Prim 50 vertices/%d/%d", gotWeight, wantWeight)
	}

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildUndirected(n, 2*n, 9, 24)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(E log V)", complexity.FitExponent(1000, 3, func(n int) {
		Prim(build(n), 0)
	}), 0.9, 1.5)
}

func TestTarjanSCC(t *testing.T) {
	g := graph.NewGraph(6)
	addEdges(g, [][3]int{
		{0, 1, 1}, {1, 2, 1}, {2, 0, 1}, {2, 3, 1}, {3, 4, 1}, {4, 3, 1},
	}, false)
	got := normalizeComponents(TarjanSCC(g))
	want := normalizeComponents([][]int{{0, 1, 2}, {3, 4}, {5}})
	if !equalComponents(got, want) {
		t.Errorf("TarjanSCC = %v, want %v", got, want)
	} else {
		okPrint("TarjanSCC two cycles and an isolated vertex/%v/%v", got, want)
	}

	dag := graph.NewGraph(4)
	addEdges(dag, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, false)
	got = normalizeComponents(TarjanSCC(dag))
	want = normalizeComponents([][]int{{0}, {1}, {2}, {3}})
	if !equalComponents(got, want) {
		t.Errorf("TarjanSCC(DAG) = %v, want %v", got, want)
	} else {
		okPrint("TarjanSCC DAG/%v/%v", got, want)
	}

	selfLoop := graph.NewGraph(1)
	selfLoop.AddEdge(0, 0, 1)
	if got := normalizeComponents(TarjanSCC(selfLoop)); !equalComponents(got, [][]int{{0}}) {
		t.Errorf("TarjanSCC(self-loop) = %v, want [[0]]", got)
	} else {
		okPrint("TarjanSCC self-loop/%v/[[0]]", got)
	}

	for _, n := range []int{1, 5, 12, 25} {
		rg := buildDirected(n, 2*n, 1, int64(n))
		got := normalizeComponents(TarjanSCC(rg))
		want := normalizeComponents(sccByReach(rg))
		if !equalComponents(got, want) {
			t.Fatalf("TarjanSCC(n=%d) = %v, want %v", n, got, want)
		}
	}
	okPrint("TarjanSCC 4 random graphs/got=mutual reachability/want=mutual reachability")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildDirected(n, 3*n, 1, 25)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(1000, 3, func(n int) {
		TarjanSCC(build(n))
	}), 0.7, 1.4)
}

func TestArticulationPointsBridges(t *testing.T) {
	path := graph.NewGraph(4)
	addEdges(path, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}}, true)
	points, bridges := ArticulationPointsBridges(path)
	if want := []int{1, 2}; !slices.Equal(points, want) {
		t.Errorf("ArticulationPointsBridges(path) points = %v, want %v", points, want)
	}
	if want := [][2]int{{0, 1}, {1, 2}, {2, 3}}; !slices.Equal(bridges, want) {
		t.Errorf("ArticulationPointsBridges(path) bridges = %v, want %v", bridges, want)
	} else {
		okPrint("ArticulationPointsBridges path/%v, %v/%v, %v", points, bridges, want, want)
	}

	tail := graph.NewGraph(4)
	addEdges(tail, [][3]int{{0, 1, 1}, {1, 2, 1}, {2, 0, 1}, {2, 3, 1}}, true)
	points, bridges = ArticulationPointsBridges(tail)
	if want := []int{2}; !slices.Equal(points, want) {
		t.Errorf("ArticulationPointsBridges(triangle with tail) points = %v, want %v", points, want)
	}
	if want := [][2]int{{2, 3}}; !slices.Equal(bridges, want) {
		t.Errorf("ArticulationPointsBridges(triangle with tail) bridges = %v, want %v", bridges, want)
	} else {
		okPrint("ArticulationPointsBridges triangle with tail/%v, %v/%v, %v", points, bridges, want, want)
	}

	disconnected := graph.NewGraph(5)
	addEdges(disconnected, [][3]int{{0, 1, 1}, {2, 3, 1}, {3, 4, 1}, {4, 2, 1}}, true)
	points, bridges = ArticulationPointsBridges(disconnected)
	if len(points) != 0 {
		t.Errorf("ArticulationPointsBridges(disconnected) points = %v, want none", points)
	}
	if want := [][2]int{{0, 1}}; !slices.Equal(bridges, want) {
		t.Errorf("ArticulationPointsBridges(disconnected) bridges = %v, want %v", bridges, want)
	} else {
		okPrint("ArticulationPointsBridges disconnected/%v, %v/%v, %v", points, bridges, want, want)
	}

	selfLoop := graph.NewGraph(2)
	addEdges(selfLoop, [][3]int{{0, 0, 1}, {0, 1, 1}}, true)
	points, bridges = ArticulationPointsBridges(selfLoop)
	if len(points) != 0 || len(bridges) != 0 {
		t.Errorf("ArticulationPointsBridges(self-loop) = (%v, %v), want none", points, bridges)
	} else {
		okPrint("ArticulationPointsBridges self-loop/%v, %v/none, none", points, bridges)
	}

	for _, n := range []int{4, 6, 9, 12} {
		rg := buildSimpleUndirected(n, 2*n, 1, int64(n))
		original := componentCount(rg, -1, [2]int{-1, -1})
		var wantPoints []int
		for v := 0; v < n; v++ {
			if componentCount(rg, v, [2]int{-1, -1}) > original {
				wantPoints = append(wantPoints, v)
			}
		}
		var wantBridges [][2]int
		for _, e := range uniqueEdges(rg) {
			if componentCount(rg, -1, [2]int{e.u, e.v}) > original {
				wantBridges = append(wantBridges, [2]int{e.u, e.v})
			}
		}
		slices.SortFunc(wantBridges, func(a, b [2]int) int {
			if a[0] != b[0] {
				return a[0] - b[0]
			}
			return a[1] - b[1]
		})
		points, bridges := ArticulationPointsBridges(rg)
		if !slices.Equal(points, wantPoints) {
			t.Fatalf("ArticulationPointsBridges(n=%d) points = %v, want %v", n, points, wantPoints)
		}
		if !slices.Equal(bridges, wantBridges) {
			t.Fatalf("ArticulationPointsBridges(n=%d) bridges = %v, want %v", n, bridges, wantBridges)
		}
	}
	okPrint("ArticulationPointsBridges 4 random graphs/got=vertex and edge removal/want=removal")

	cache := make(map[int]*graph.Graph)
	build := func(n int) *graph.Graph {
		if g, ok := cache[n]; ok {
			return g
		}
		g := buildSimpleUndirected(n, 0, 1, 26)
		cache[n] = g
		return g
	}
	assertExponent(t, "O(V+E)", complexity.FitExponent(1000, 3, func(n int) {
		ArticulationPointsBridges(build(n))
	}), 0.7, 1.4)
}

func TestEdmondsKarp(t *testing.T) {
	clrs := graph.NewGraph(6)
	addEdges(clrs, [][3]int{
		{0, 1, 16}, {0, 2, 13}, {1, 2, 10}, {2, 1, 4},
		{1, 3, 12}, {3, 2, 9}, {2, 4, 14}, {4, 3, 7},
		{3, 5, 20}, {4, 5, 4},
	}, false)
	if got := EdmondsKarp(clrs, 0, 5); got != 23 {
		t.Errorf("EdmondsKarp(CLRS network) = %d, want 23", got)
	} else {
		okPrint("EdmondsKarp CLRS network/%d/23", got)
	}

	simple := graph.NewGraph(3)
	addEdges(simple, [][3]int{{0, 1, 3}, {1, 2, 2}}, false)
	if got := EdmondsKarp(simple, 0, 2); got != 2 {
		t.Errorf("EdmondsKarp(simple path) = %d, want 2", got)
	} else {
		okPrint("EdmondsKarp simple path/%d/2", got)
	}

	disconnected := graph.NewGraph(4)
	addEdges(disconnected, [][3]int{{0, 1, 5}, {2, 3, 5}}, false)
	if got := EdmondsKarp(disconnected, 0, 3); got != 0 {
		t.Errorf("EdmondsKarp(disconnected) = %d, want 0", got)
	} else {
		okPrint("EdmondsKarp disconnected/%d/0", got)
	}

	parallel := graph.NewGraph(2)
	parallel.AddEdge(0, 1, 1)
	parallel.AddEdge(0, 1, 1)
	if got := EdmondsKarp(parallel, 0, 1); got != 2 {
		t.Errorf("EdmondsKarp(parallel edges) = %d, want 2", got)
	} else {
		okPrint("EdmondsKarp parallel edges/%d/2", got)
	}

	if got := EdmondsKarp(clrs, 0, 0); got != 0 {
		t.Errorf("EdmondsKarp(s == t) = %d, want 0", got)
	} else {
		okPrint("EdmondsKarp s == t/%d/0", got)
	}

	for _, n := range []int{4, 5, 6, 7, 8} {
		rg := buildDirected(n, 3*n, 5, int64(n))
		want := minCutBrute(rg, 0, n-1)
		if got := EdmondsKarp(rg, 0, n-1); got != want {
			t.Fatalf("EdmondsKarp(n=%d) = %d, want min cut %d", n, got, want)
		}
	}
	okPrint("EdmondsKarp 5 random graphs/got=brute force min cut/want=min cut")
}
