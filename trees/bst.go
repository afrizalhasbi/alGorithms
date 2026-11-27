package trees

// [medium] BST is a binary search tree over integers. It supports insert,
// search, delete, and inorder enumeration.
//
// Background. A binary search tree stores values so that every value in a
// left subtree is smaller than the node's value and every value in a
// right subtree is larger. Searching compares the target with the current
// node and moves left or right, so the number of comparisons is the
// height of the tree. Inserting follows the same path and attaches the
// new value as a leaf. Deleting is the only operation with cases: a leaf
// is removed directly, a node with one child is replaced by that child,
// and a node with two children is replaced by its inorder successor, the
// smallest value in its right subtree. The successor is then removed
// from the right subtree.
//
// Definitions:
//   - binary search tree property: for every node, all values in its left
//     subtree are smaller than its value, and all values in its right
//     subtree are larger.
//   - height: the number of edges on the longest path from the root to a
//     leaf. An empty tree has height -1 and a single node has height 0.
//   - inorder successor: the smallest value in a node's right subtree.
//   - duplicate: a value already present in the tree. Insert ignores
//     duplicates.
//
// Steps:
//  1. Insert(v): if the tree is empty, make v the root. Otherwise compare
//     v with the current node. If v is smaller, go to the left child; if
//     larger, go to the right child; if equal, stop without changing the
//     tree. When the chosen child is nil, attach a new node with value v
//     there.
//  2. Search(v): compare v with the current node. If equal, return true.
//     If smaller, continue at the left child; if larger, continue at the
//     right child. If the child is nil, return false.
//  3. Delete(v): find the node holding v. If it does not exist, do
//     nothing. If it is a leaf, remove it. If it has exactly one child,
//     replace it with that child. If it has two children, find its
//     inorder successor, copy the successor's value into the node, and
//     delete the successor from the right subtree. The successor has no
//     left child, so its own deletion is one of the first two cases.
//  4. InOrder(): return the values by traversing the left subtree, then
//     the node, then the right subtree.
//
// Notes:
//   - Insert, Search, and Delete take O(h) time, where h is the height.
//     The height is O(n) in the worst case and O(log n) only when the
//     tree is balanced. This type does not balance itself.
//   - InOrder takes O(n) time.
//   - The zero value is an empty tree.
//
// Expected: O(h) per operation, O(n) for InOrder; O(n) space
// Input: mutated in place
type BST struct {
	Root *TreeNode
}

// Insert adds v to the tree unless it is already present.
func (b *BST) Insert(v int) {
	// TODO: implement
}

// Search reports whether v is present.
func (b *BST) Search(v int) bool {
	// TODO: implement
	return false
}

// Delete removes v from the tree if it is present.
func (b *BST) Delete(v int) {
	// TODO: implement
}

// InOrder returns the values in increasing order.
func (b *BST) InOrder() []int {
	// TODO: implement
	return nil
}
