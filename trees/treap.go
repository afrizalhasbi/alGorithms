package trees

// [medium] Treap is a binary search tree that also keeps a heap order on
// a fixed priority per value, which keeps it balanced in expectation. It
// supports insert, erase, and inorder enumeration.
//
// Background. A treap combines a search tree and a heap. Values obey the
// binary search tree property; priorities obey the heap property, with
// larger priorities closer to the root. If the priorities are random and
// independent of the values, the resulting shape is the same as a binary
// search tree built from a random insertion order, so the expected height
// is logarithmic. Inserting places the new value as a leaf and then
// rotates it upward while its priority is larger than its parent's.
// Erasing a node joins its two subtrees by repeatedly taking the subtree
// with the larger root priority. Priorities are derived from the value
// with the provided hash, so the shape is deterministic for a given set
// of values.
//
// Definitions:
//   - priority: priority(v) is the provided 64-bit hash of v. Larger means
//     closer to the root.
//   - heap property: every node's priority is at least the priority of
//     each of its children.
//   - binary search tree property: every value in a left subtree is
//     smaller than the node's value, and every value in a right subtree
//     is larger.
//   - rotation up: a left rotation or right rotation that moves a node
//     toward the root while preserving the search tree property.
//   - merge of two subtrees: joining a left subtree and a right subtree
//     whose values are all ordered correctly, by making the root with the
//     larger priority the parent and merging the remaining parts
//     recursively.
//   - duplicate: a value already present. Insert ignores duplicates.
//
// Steps:
//  1. Insert(v): perform an ordinary binary search tree insertion with
//     priority(v). Ignore v when it is already present. On the way back
//     up, if the new node or subtree root has a larger priority than its
//     parent, rotate it above the parent. Continue until the heap
//     property holds or the node is the root.
//  2. Erase(v): find the node holding v. If it does not exist, do
//     nothing. Otherwise replace the node with the merge of its left and
//     right subtrees. To merge two subtrees, if one is empty return the
//     other; otherwise make the root with the larger priority the new
//     root, and merge the remaining subtree into the appropriate child.
//  3. InOrder(): traverse the left subtree, the node, then the right
//     subtree.
//
// Notes:
//   - The zero value is an empty tree.
//   - Duplicates are ignored.
//   - The priority of every node must equal treapPriority(node.Val).
//   - Expected height is O(log n); the worst case is O(n) but has
//     negligible probability for the fixed hash on distinct values.
//
// Expected: O(log n) expected per operation, O(n) for InOrder; O(n) space
// Input: mutated in place
type Treap struct {
	Root *TreapNode
}

// TreapNode is one node of a treap.
type TreapNode struct {
	Val      int
	Priority uint64
	Left     *TreapNode
	Right    *TreapNode
}

// Insert adds v unless it is already present, and restores the heap
// property.
func (t *Treap) Insert(v int) {
	// TODO: implement
}

// Erase removes v if it is present, and restores the heap property.
func (t *Treap) Erase(v int) {
	// TODO: implement
}

// InOrder returns the values in increasing order.
func (t *Treap) InOrder() []int {
	// TODO: implement
	return nil
}

// treapPriority is the fixed priority of a value: a SplitMix64 hash of
// the value. It is provided so that the tree shape is deterministic.
func treapPriority(v int) uint64 {
	z := uint64(v) + 0x9e3779b97f4a7c15
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
