package trees

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"

	"alGorithms/internal/complexity"
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

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// completeTree returns a complete binary tree with values 0..n-1 in
// level order.
func completeTree(n int) *TreeNode {
	if n == 0 {
		return nil
	}
	nodes := make([]*TreeNode, n)
	for i := range nodes {
		nodes[i] = &TreeNode{Val: i}
	}
	for i := 0; i < n; i++ {
		if 2*i+1 < n {
			nodes[i].Left = nodes[2*i+1]
		}
		if 2*i+2 < n {
			nodes[i].Right = nodes[2*i+2]
		}
	}
	return nodes[0]
}

// randomBinaryTree attaches each new value to a random free child. Values
// are a random permutation of 0..n-1, so they are unique.
func randomBinaryTree(n int, seed int64) *TreeNode {
	if n == 0 {
		return nil
	}
	rng := rand.New(rand.NewSource(seed))
	vals := rng.Perm(n)
	root := &TreeNode{Val: vals[0]}
	nodes := []*TreeNode{root}
	for i := 1; i < n; i++ {
		v := vals[i]
		for {
			p := nodes[rng.Intn(len(nodes))]
			if rng.Intn(2) == 0 {
				if p.Left == nil {
					p.Left = &TreeNode{Val: v}
					nodes = append(nodes, p.Left)
					break
				}
			} else if p.Right == nil {
				p.Right = &TreeNode{Val: v}
				nodes = append(nodes, p.Right)
				break
			}
		}
	}
	return root
}

func walkInorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := walkInorder(root.Left)
	out = append(out, root.Val)
	return append(out, walkInorder(root.Right)...)
}

func walkPreorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := []int{root.Val}
	out = append(out, walkPreorder(root.Left)...)
	return append(out, walkPreorder(root.Right)...)
}

func walkPostorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	out := walkPostorder(root.Left)
	out = append(out, walkPostorder(root.Right)...)
	return append(out, root.Val)
}

func walkLevel(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var out []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		out = append(out, n.Val)
		if n.Left != nil {
			queue = append(queue, n.Left)
		}
		if n.Right != nil {
			queue = append(queue, n.Right)
		}
	}
	return out
}

// snapshot records a preorder shape with -1 for every nil child. Values
// in the tests are non-negative, so -1 is unambiguous.
func snapshot(root *TreeNode) []int {
	var out []int
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			out = append(out, -1)
			return
		}
		out = append(out, n.Val)
		walk(n.Left)
		walk(n.Right)
	}
	walk(root)
	return out
}

func checkAVL(t *testing.T, n *AVLNode) int {
	t.Helper()
	if n == nil {
		return 0
	}
	lh := checkAVL(t, n.Left)
	rh := checkAVL(t, n.Right)
	want := 1 + maxInt(lh, rh)
	if n.Height != want {
		t.Errorf("AVL node %d height = %d, want %d", n.Val, n.Height, want)
	}
	if diff := lh - rh; diff < -1 || diff > 1 {
		t.Errorf("AVL node %d is unbalanced: left height %d, right height %d", n.Val, lh, rh)
	}
	return n.Height
}

func checkTreap(t *testing.T, n *TreapNode) {
	t.Helper()
	if n == nil {
		return
	}
	if n.Priority != treapPriority(n.Val) {
		t.Errorf("treap node %d priority = %d, want %d", n.Val, n.Priority, treapPriority(n.Val))
	}
	if n.Left != nil && n.Left.Priority > n.Priority {
		t.Errorf("treap node %d priority %d is below left child %d priority %d",
			n.Val, n.Priority, n.Left.Val, n.Left.Priority)
	}
	if n.Right != nil && n.Right.Priority > n.Priority {
		t.Errorf("treap node %d priority %d is below right child %d priority %d",
			n.Val, n.Priority, n.Right.Val, n.Right.Priority)
	}
	checkTreap(t, n.Left)
	checkTreap(t, n.Right)
}

func randomWords(n, length int, seed int64) []string {
	rng := rand.New(rand.NewSource(seed))
	out := make([]string, n)
	for i := range out {
		b := make([]byte, length)
		for j := range b {
			b[j] = byte('a' + rng.Intn(26))
		}
		out[i] = string(b)
	}
	return out
}

func naiveLCA(root *TreeNode, a, b int) int {
	parent := make(map[int]int)
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		if n.Left != nil {
			parent[n.Left.Val] = n.Val
			walk(n.Left)
		}
		if n.Right != nil {
			parent[n.Right.Val] = n.Val
			walk(n.Right)
		}
	}
	walk(root)

	ancestors := make(map[int]bool)
	for x := a; ; {
		ancestors[x] = true
		p, ok := parent[x]
		if !ok {
			break
		}
		x = p
	}
	for x := b; ; {
		if ancestors[x] {
			return x
		}
		p, ok := parent[x]
		if !ok {
			break
		}
		x = p
	}
	return root.Val
}

// ---- one test per algorithm: correctness + complexity, always together ----

func TestTraversals(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}

	in, pre, post, level := Traversals(root)
	cases := []struct {
		name string
		got  []int
		want []int
	}{
		{"inorder", in, []int{4, 2, 5, 1, 6, 3}},
		{"preorder", pre, []int{1, 2, 4, 5, 3, 6}},
		{"postorder", post, []int{4, 5, 2, 6, 3, 1}},
		{"levelorder", level, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, c := range cases {
		if !slices.Equal(c.got, c.want) {
			t.Errorf("Traversals %s = %v, want %v", c.name, c.got, c.want)
		} else {
			okPrint("Traversals %s/%v/%v", c.name, c.got, c.want)
		}
	}

	if in, pre, post, level := Traversals(nil); len(in)+len(pre)+len(post)+len(level) != 0 {
		t.Errorf("Traversals(nil) = %v %v %v %v, want four empty sequences", in, pre, post, level)
	} else {
		okPrint("Traversals nil/four empty sequences/four empty sequences")
	}

	for _, n := range []int{1, 7, 50} {
		tree := randomBinaryTree(n, int64(n))
		in, pre, post, level := Traversals(tree)
		if !slices.Equal(in, walkInorder(tree)) ||
			!slices.Equal(pre, walkPreorder(tree)) ||
			!slices.Equal(post, walkPostorder(tree)) ||
			!slices.Equal(level, walkLevel(tree)) {
			t.Fatalf("Traversals(n=%d) disagrees with recursive reference", n)
		}
	}
	okPrint("Traversals 3 random trees/got=recursive reference/want=recursive reference")

	cache := make(map[int]*TreeNode)
	build := func(n int) *TreeNode {
		if root, ok := cache[n]; ok {
			return root
		}
		root := completeTree(n)
		cache[n] = root
		return root
	}
	assertExponent(t, "O(n)", complexity.FitExponent(1000, 3, func(n int) {
		Traversals(build(n))
	}), 0.7, 1.4)
}

func TestBST(t *testing.T) {
	b := &BST{}
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9, 2, 6} {
		b.Insert(v)
	}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if got := b.InOrder(); !slices.Equal(got, want) {
		t.Errorf("BST InOrder = %v, want %v", got, want)
	} else {
		okPrint("BST InOrder/%v/%v", got, want)
	}
	for _, v := range want {
		if !b.Search(v) {
			t.Errorf("BST Search(%d) = false, want true", v)
		}
	}
	if b.Search(42) {
		t.Error("BST Search(42) = true, want false")
	} else {
		okPrint("BST search present and absent values/all present/true, 42/false")
	}

	b.Insert(5)
	if got := b.InOrder(); !slices.Equal(got, want) {
		t.Errorf("BST duplicate insert changed InOrder to %v, want %v", got, want)
	} else {
		okPrint("BST duplicate insert ignored/%v/%v", got, want)
	}

	cases := []struct {
		name   string
		values []int
		del    int
		want   []int
	}{
		{"leaf", []int{5, 3, 8}, 3, []int{5, 8}},
		{"one child", []int{5, 3, 1}, 3, []int{1, 5}},
		{"two children", []int{5, 3, 8, 1, 4, 7, 9}, 5, []int{1, 3, 4, 7, 8, 9}},
		{"root one child", []int{5, 3}, 5, []int{3}},
		{"missing", []int{5, 3, 8}, 42, []int{3, 5, 8}},
	}
	for _, c := range cases {
		tree := &BST{}
		for _, v := range c.values {
			tree.Insert(v)
		}
		tree.Delete(c.del)
		if got := tree.InOrder(); !slices.Equal(got, c.want) {
			t.Errorf("BST Delete(%s, %d) = %v, want %v", c.name, c.del, got, c.want)
		} else {
			okPrint("BST delete %s/%v/%v", c.name, got, c.want)
		}
	}

	const n = 100
	rng := rand.New(rand.NewSource(32))
	vals := rng.Perm(n)
	random := &BST{}
	for _, v := range vals {
		random.Insert(v)
	}
	if got := random.InOrder(); !slices.Equal(got, sequence(n)) {
		t.Fatalf("BST random inserts InOrder = %v, want 0..%d", got, n-1)
	}
	remaining := make([]int, 0, n)
	for v := 0; v < n; v++ {
		if v%3 == 0 {
			random.Delete(v)
		} else {
			remaining = append(remaining, v)
		}
	}
	if got := random.InOrder(); !slices.Equal(got, remaining) {
		t.Fatalf("BST random deletes InOrder = %v, want %v", got, remaining)
	}
	okPrint("BST 100 random inserts and deletes/got=sorted remainder/want=sorted remainder")

	cache := make(map[int][]int)
	perm := func(n int) []int {
		if p, ok := cache[n]; ok {
			return p
		}
		p := rand.New(rand.NewSource(int64(n))).Perm(n)
		cache[n] = p
		return p
	}
	assertExponent(t, "O(n log n) random inserts", complexity.FitExponent(1000, 3, func(n int) {
		tree := &BST{}
		for _, v := range perm(n) {
			tree.Insert(v)
		}
	}), 0.9, 1.6)
}

func sequence(n int) []int {
	out := make([]int, n)
	for i := range out {
		out[i] = i
	}
	return out
}

func rangeInts(lo, hi int) []int {
	out := make([]int, 0, hi-lo+1)
	for v := lo; v <= hi; v++ {
		out = append(out, v)
	}
	return out
}

func TestMorrisTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}

	before := snapshot(root)
	got := MorrisInOrder(root)
	want := []int{4, 2, 5, 1, 6, 3}
	if !slices.Equal(got, want) {
		t.Errorf("MorrisInOrder = %v, want %v", got, want)
	} else {
		okPrint("MorrisInOrder/%v/%v", got, want)
	}
	if after := snapshot(root); !slices.Equal(before, after) {
		t.Errorf("MorrisInOrder changed the tree: before %v, after %v", before, after)
	} else {
		okPrint("MorrisInOrder tree restored/before=after/before=after")
	}

	if got := MorrisInOrder(nil); len(got) != 0 {
		t.Errorf("MorrisInOrder(nil) = %v, want empty", got)
	} else {
		okPrint("MorrisInOrder nil/empty/empty")
	}

	for _, n := range []int{1, 7, 50} {
		tree := randomBinaryTree(n, int64(n))
		before := snapshot(tree)
		got := MorrisInOrder(tree)
		if want := walkInorder(tree); !slices.Equal(got, want) {
			t.Fatalf("MorrisInOrder(n=%d) = %v, want %v", n, got, want)
		}
		if after := snapshot(tree); !slices.Equal(before, after) {
			t.Fatalf("MorrisInOrder(n=%d) did not restore the tree", n)
		}
	}
	okPrint("MorrisInOrder 3 random trees/got=recursive inorder/want=recursive inorder")

	cache := make(map[int]*TreeNode)
	build := func(n int) *TreeNode {
		if root, ok := cache[n]; ok {
			return root
		}
		root := completeTree(n)
		cache[n] = root
		return root
	}
	assertExponent(t, "O(n)", complexity.FitExponent(1000, 3, func(n int) {
		MorrisInOrder(build(n))
	}), 0.7, 1.4)
}

func TestAVL(t *testing.T) {
	a := &AVL{}
	for v := 1; v <= 100; v++ {
		a.Insert(v)
	}
	if got := a.InOrder(); !slices.Equal(got, rangeInts(1, 100)) {
		t.Errorf("AVL sorted inserts InOrder = %v, want 1..100", got)
	} else {
		okPrint("AVL sorted inserts InOrder/1..100/1..100")
	}
	if got := checkAVL(t, a.Root); got != a.Height() {
		t.Errorf("AVL Height() = %d, but root Height field is %d", a.Height(), got)
	}
	if h := a.Height(); float64(h) > 1.45*math.Log2(102) {
		t.Errorf("AVL height %d exceeds the AVL bound for 100 nodes", h)
	} else {
		okPrint("AVL height after 100 sorted inserts/%d/<=%d", h, int(1.45*math.Log2(102)))
	}

	for v := 2; v <= 100; v += 2 {
		a.Delete(v)
	}
	wantOdd := make([]int, 0, 50)
	for v := 1; v <= 99; v += 2 {
		wantOdd = append(wantOdd, v)
	}
	if got := a.InOrder(); !slices.Equal(got, wantOdd) {
		t.Errorf("AVL after deleting evens InOrder = %v, want %v", got, wantOdd)
	} else {
		okPrint("AVL delete evens/%v/%v", got, wantOdd)
	}
	checkAVL(t, a.Root)
	okPrint("AVL balance after deletes/balanced/balanced")

	rng := rand.New(rand.NewSource(33))
	const n = 200
	vals := rng.Perm(n)
	random := &AVL{}
	for _, v := range vals {
		random.Insert(v)
	}
	if got := random.InOrder(); !slices.Equal(got, sequence(n)) {
		t.Fatalf("AVL random inserts InOrder = %v, want 0..%d", got, n-1)
	}
	checkAVL(t, random.Root)
	for v := 0; v < n; v += 2 {
		random.Delete(v)
	}
	wantOdd = wantOdd[:0]
	for v := 1; v < n; v += 2 {
		wantOdd = append(wantOdd, v)
	}
	if got := random.InOrder(); !slices.Equal(got, wantOdd) {
		t.Fatalf("AVL random deletes InOrder = %v, want %v", got, wantOdd)
	}
	checkAVL(t, random.Root)
	okPrint("AVL 200 random inserts and 100 deletes/balanced/balanced")

	dup := &AVL{}
	dup.Insert(5)
	dup.Insert(5)
	if got := dup.InOrder(); !slices.Equal(got, []int{5}) {
		t.Errorf("AVL duplicate insert InOrder = %v, want [5]", got)
	} else {
		okPrint("AVL duplicate insert ignored/%v/[5]", got)
	}

	cache := make(map[int][]int)
	perm := func(n int) []int {
		if p, ok := cache[n]; ok {
			return p
		}
		p := rand.New(rand.NewSource(int64(n))).Perm(n)
		cache[n] = p
		return p
	}
	assertExponent(t, "O(n log n) random inserts", complexity.FitExponent(1000, 3, func(n int) {
		tree := &AVL{}
		for _, v := range perm(n) {
			tree.Insert(v)
		}
	}), 0.9, 1.6)
}

func TestTreap(t *testing.T) {
	rng := rand.New(rand.NewSource(34))
	const n = 200
	vals := rng.Perm(n)
	tr := &Treap{}
	for _, v := range vals {
		tr.Insert(v)
	}
	if got := tr.InOrder(); !slices.Equal(got, sequence(n)) {
		t.Fatalf("Treap InOrder = %v, want 0..%d", got, n-1)
	} else {
		okPrint("Treap 200 random inserts InOrder/0..199/0..199")
	}
	checkTreap(t, tr.Root)
	okPrint("Treap heap and priority properties/hold/hold")

	for v := 0; v < n; v += 2 {
		tr.Erase(v)
	}
	want := make([]int, 0, n/2)
	for v := 1; v < n; v += 2 {
		want = append(want, v)
	}
	if got := tr.InOrder(); !slices.Equal(got, want) {
		t.Fatalf("Treap after erasing evens InOrder = %v, want %v", got, want)
	}
	checkTreap(t, tr.Root)
	okPrint("Treap 100 erases/heap and priority hold/hold")

	dup := &Treap{}
	dup.Insert(5)
	dup.Insert(5)
	if got := dup.InOrder(); !slices.Equal(got, []int{5}) {
		t.Errorf("Treap duplicate insert InOrder = %v, want [5]", got)
	} else {
		okPrint("Treap duplicate insert ignored/%v/[5]", got)
	}

	cache := make(map[int][]int)
	perm := func(n int) []int {
		if p, ok := cache[n]; ok {
			return p
		}
		p := rand.New(rand.NewSource(int64(n))).Perm(n)
		cache[n] = p
		return p
	}
	assertExponent(t, "O(n log n) random inserts", complexity.FitExponent(1000, 3, func(n int) {
		tree := &Treap{}
		for _, v := range perm(n) {
			tree.Insert(v)
		}
	}), 0.9, 1.6)
}

func TestSegmentTree(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	s := NewSegmentTree(xs)
	cases := []struct {
		l, r int
		want int
	}{
		{0, 5, 15},
		{1, 4, 9},
		{2, 2, 0},
		{0, 0, 0},
		{4, 5, 5},
		{0, 1, 1},
	}
	for _, c := range cases {
		if got := s.Query(c.l, c.r); got != c.want {
			t.Errorf("SegmentTree.Query(%d, %d) = %d, want %d", c.l, c.r, got, c.want)
		} else {
			okPrint("SegmentTree Query(%d, %d)/%d/%d", c.l, c.r, got, c.want)
		}
	}

	s.Update(2, 10)
	if got := s.Query(0, 5); got != 22 {
		t.Errorf("SegmentTree after Update(2, 10), Query(0, 5) = %d, want 22", got)
	} else {
		okPrint("SegmentTree update and full query/%d/22", got)
	}
	if got := s.Query(2, 3); got != 10 {
		t.Errorf("SegmentTree.Query(2, 3) after update = %d, want 10", got)
	} else {
		okPrint("SegmentTree updated leaf/%d/10", got)
	}

	empty := NewSegmentTree(nil)
	if got := empty.Query(0, 0); got != 0 {
		t.Errorf("SegmentTree(empty).Query(0, 0) = %d, want 0", got)
	} else {
		okPrint("SegmentTree empty/%d/0", got)
	}

	const n = 100
	naive := make([]int, n)
	for i := range naive {
		naive[i] = i % 17
	}
	random := NewSegmentTree(naive)
	rng := rand.New(rand.NewSource(35))
	for step := 0; step < 500; step++ {
		if rng.Intn(2) == 0 {
			i := rng.Intn(n)
			v := rng.Intn(100) - 50
			naive[i] = v
			random.Update(i, v)
		} else {
			l := rng.Intn(n)
			r := l + rng.Intn(n-l+1)
			want := 0
			for i := l; i < r; i++ {
				want += naive[i]
			}
			if got := random.Query(l, r); got != want {
				t.Fatalf("SegmentTree random Query(%d, %d) = %d, want %d", l, r, got, want)
			}
		}
	}
	okPrint("SegmentTree 500 random updates and queries/got=naive sums/want=naive sums")

	cache := make(map[int][]int)
	values := func(n int) []int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := make([]int, n)
		for i := range v {
			v[i] = i % 17
		}
		cache[n] = v
		return v
	}
	assertExponent(t, "O(n log n) build and queries", complexity.FitExponent(1000, 3, func(n int) {
		tree := NewSegmentTree(values(n))
		for i := 0; i < n; i++ {
			l := i % n
			r := l + 10
			if r > n {
				r = n
			}
			tree.Query(l, r)
		}
	}), 0.9, 1.5)
}

func TestFenwick(t *testing.T) {
	f := NewFenwick(5)
	for i := 0; i < 5; i++ {
		f.Add(i, i+1)
	}
	prefix := []int{0, 1, 3, 6, 10, 15}
	for i, want := range prefix {
		if got := f.PrefixSum(i); got != want {
			t.Errorf("Fenwick.PrefixSum(%d) = %d, want %d", i, got, want)
		} else {
			okPrint("Fenwick PrefixSum(%d)/%d/%d", i, got, want)
		}
	}
	if got := f.RangeSum(1, 4); got != 9 {
		t.Errorf("Fenwick.RangeSum(1, 4) = %d, want 9", got)
	} else {
		okPrint("Fenwick RangeSum(1, 4)/%d/9", got)
	}
	if got := f.RangeSum(0, 5); got != 15 {
		t.Errorf("Fenwick.RangeSum(0, 5) = %d, want 15", got)
	} else {
		okPrint("Fenwick RangeSum(0, 5)/%d/15", got)
	}

	const n = 100
	random := NewFenwick(n)
	naive := make([]int, n)
	rng := rand.New(rand.NewSource(36))
	for step := 0; step < 500; step++ {
		if rng.Intn(2) == 0 {
			i := rng.Intn(n)
			d := rng.Intn(21) - 10
			naive[i] += d
			random.Add(i, d)
		} else {
			l := rng.Intn(n + 1)
			r := l + rng.Intn(n-l+1)
			want := 0
			for i := l; i < r; i++ {
				want += naive[i]
			}
			if got := random.RangeSum(l, r); got != want {
				t.Fatalf("Fenwick random RangeSum(%d, %d) = %d, want %d", l, r, got, want)
			}
			if got := random.PrefixSum(r); got != prefixSum(naive, r) {
				t.Fatalf("Fenwick random PrefixSum(%d) = %d, want %d", r, got, prefixSum(naive, r))
			}
		}
	}
	okPrint("Fenwick 500 random adds and queries/got=naive prefix sums/want=naive prefix sums")

	empty := NewFenwick(0)
	if got := empty.PrefixSum(0); got != 0 {
		t.Errorf("Fenwick(0).PrefixSum(0) = %d, want 0", got)
	} else {
		okPrint("Fenwick empty/%d/0", got)
	}

	cache := make(map[int][]int)
	values := func(n int) []int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := make([]int, n)
		for i := range v {
			v[i] = i % 13
		}
		cache[n] = v
		return v
	}
	assertExponent(t, "O(n log n) adds and queries", complexity.FitExponent(1000, 3, func(n int) {
		tree := NewFenwick(n)
		for i, v := range values(n) {
			tree.Add(i, v)
		}
		for i := 0; i < n; i++ {
			tree.PrefixSum((i * 7919) % (n + 1))
		}
	}), 0.9, 1.5)
}

func prefixSum(xs []int, i int) int {
	sum := 0
	for j := 0; j < i; j++ {
		sum += xs[j]
	}
	return sum
}

func TestTrie(t *testing.T) {
	tr := NewTrie()
	for _, w := range []string{"apple", "app", "ape", "bat", "app"} {
		tr.Insert(w)
	}
	for _, w := range []string{"apple", "app", "ape", "bat"} {
		if !tr.Search(w) {
			t.Errorf("Trie.Search(%q) = false, want true", w)
		}
	}
	if tr.Search("ap") {
		t.Error("Trie.Search(\"ap\") = true, want false (prefix only)")
	} else {
		okPrint("Trie exact search present and prefix-only word/all present/true, ap/false")
	}
	if !tr.StartsWith("ap") || !tr.StartsWith("bat") {
		t.Error("Trie.StartsWith failed for a present prefix")
	}
	if tr.StartsWith("cat") {
		t.Error("Trie.StartsWith(\"cat\") = true, want false")
	} else {
		okPrint("Trie prefix search present and absent/true, false/true, false")
	}
	counts := []struct {
		prefix string
		want   int
	}{
		{"", 5},
		{"ap", 4},
		{"app", 3},
		{"appl", 1},
		{"b", 1},
		{"cat", 0},
	}
	for _, c := range counts {
		if got := tr.CountPrefix(c.prefix); got != c.want {
			t.Errorf("Trie.CountPrefix(%q) = %d, want %d", c.prefix, got, c.want)
		} else {
			okPrint("Trie CountPrefix(%q)/%d/%d", c.prefix, got, c.want)
		}
	}

	words := randomWords(300, 5, 37)
	random := NewTrie()
	countsByPrefix := make(map[string]int)
	for _, w := range words {
		random.Insert(w)
		for i := 0; i <= len(w); i++ {
			countsByPrefix[w[:i]]++
		}
	}
	for _, w := range words {
		if !random.Search(w) {
			t.Fatalf("Trie.Search(%q) = false after insertion", w)
		}
	}
	for _, p := range []string{"", "a", "ab", "abc", "zzzzz", "q"} {
		if got, want := random.CountPrefix(p), countsByPrefix[p]; got != want {
			t.Fatalf("Trie.CountPrefix(%q) = %d, want %d", p, got, want)
		}
		if got, want := random.StartsWith(p), countsByPrefix[p] > 0; got != want {
			t.Fatalf("Trie.StartsWith(%q) = %v, want %v", p, got, want)
		}
	}
	okPrint("Trie 300 random words/got=prefix counts/want=prefix counts")

	cache := make(map[int][]string)
	wordsFor := func(n int) []string {
		if w, ok := cache[n]; ok {
			return w
		}
		w := randomWords(n, 8, int64(n))
		cache[n] = w
		return w
	}
	assertExponent(t, "O(total characters)", complexity.FitExponent(1000, 3, func(n int) {
		tree := NewTrie()
		for _, w := range wordsFor(n) {
			tree.Insert(w)
		}
	}), 0.7, 1.4)
}

func TestLCA(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}

	lca := NewLCA(root)
	cases := []struct {
		a, b, want int
	}{
		{5, 1, 3},
		{5, 4, 5},
		{6, 4, 5},
		{7, 4, 2},
		{0, 8, 1},
		{6, 6, 6},
	}
	for _, c := range cases {
		if got := lca.Query(c.a, c.b); got != c.want {
			t.Errorf("LCA.Query(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		} else {
			okPrint("LCA Query(%d, %d)/%d/%d", c.a, c.b, got, c.want)
		}
	}

	for _, n := range []int{1, 7, 50, 200} {
		tree := randomBinaryTree(n, int64(n))
		lca := NewLCA(tree)
		rng := rand.New(rand.NewSource(int64(n)))
		for q := 0; q < 200; q++ {
			a, b := rng.Intn(n), rng.Intn(n)
			got, want := lca.Query(a, b), naiveLCA(tree, a, b)
			if got != want {
				t.Fatalf("LCA(n=%d).Query(%d, %d) = %d, want %d", n, a, b, got, want)
			}
		}
	}
	okPrint("LCA 4 random trees, 200 queries each/got=naive parent walk/want=naive parent walk")

	cache := make(map[int]*TreeNode)
	build := func(n int) *TreeNode {
		if root, ok := cache[n]; ok {
			return root
		}
		root := completeTree(n)
		cache[n] = root
		return root
	}
	assertExponent(t, "O(n log n) build and queries", complexity.FitExponent(1000, 3, func(n int) {
		lca := NewLCA(build(n))
		for i := 0; i < n; i++ {
			lca.Query((i*7919)%n, (i*104729)%n)
		}
	}), 0.9, 1.6)
}
