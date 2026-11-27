package trees

// [hard] AVL is a self-balancing binary search tree. It supports insert,
// delete, inorder enumeration, and height.
//
// Background. A plain binary search tree degenerates into a path when
// values arrive in sorted order, which makes every operation linear. An
// AVL tree prevents this by keeping the two subtree heights of every node
// within one of each other. After an insertion or deletion, the height of
// each node on the path back to the root is recomputed. If a node's
// balance factor leaves the allowed range, one or two rotations restore
// it. A single rotation handles a left-left or right-right imbalance; a
// double rotation handles a left-right or right-left imbalance. The
// height bound guarantees that the tree height stays logarithmic.
//
// Definitions:
//   - height: the number of nodes on the longest path from a node to a
//     leaf in its subtree. A nil child has height zero and a leaf has
//     height one.
//   - balance factor: the height of the left subtree minus the height of
//     the right subtree.
//   - AVL property: every node has a balance factor of -1, 0, or 1.
//   - right rotation at x: the left child y of x becomes the new root of
//     the subtree; y's right subtree becomes x's left subtree; x becomes
//     y's right child.
//   - left rotation at x: the right child y of x becomes the new root of
//     the subtree; y's left subtree becomes x's right subtree; x becomes
//     y's left child.
//   - left-right case: x has balance factor 2 and its left child has
//     balance factor -1. Rotate the left child left, then rotate x right.
//   - right-left case: x has balance factor -2 and its right child has
//     balance factor 1. Rotate the right child right, then rotate x left.
//   - inorder successor: the smallest value in a node's right subtree.
//
// Steps:
//  1. Insert(v): perform an ordinary binary search tree insertion. Ignore
//     v when it is already present. On the way back up, recompute the
//     height of each visited node from its children, then repair it: if
//     its balance factor is 2, use a right rotation when its left child's
//     balance factor is non-negative, otherwise a left-right rotation. If
//     its balance factor is -2, use a left rotation when its right
//     child's balance factor is non-positive, otherwise a right-left
//     rotation. Return the root of the repaired subtree.
//  2. Delete(v): perform an ordinary binary search tree deletion. For a
//     node with two children, copy the inorder successor's value into the
//     node and delete the successor from the right subtree. On the way
//     back up, recompute heights and repair imbalances exactly as in
//     insertion.
//  3. InOrder(): traverse the left subtree, the node, then the right
//     subtree.
//  4. Height(): return the height stored at the root, or zero for an
//     empty tree.
//
// Notes:
//   - The zero value is an empty tree.
//   - Duplicates are ignored.
//   - The Height field of every node must equal one plus the larger of its
//     children's heights, with nil children counting as zero.
//   - The height of an AVL tree with n nodes is at most about
//     1.44*log2(n+2).
//
// Expected: O(log n) per operation, O(n) for InOrder; O(n) space
// Input: mutated in place
type AVL struct {
	Root *AVLNode
}

// AVLNode is one node of an AVL tree.
type AVLNode struct {
	Val    int
	Height int
	Left   *AVLNode
	Right  *AVLNode
}

// Insert adds v unless it is already present, and restores the AVL
// property.
func (a *AVL) Insert(v int) {
	// TODO: implement
}

// Delete removes v if it is present, and restores the AVL property.
func (a *AVL) Delete(v int) {
	// TODO: implement
}

// InOrder returns the values in increasing order.
func (a *AVL) InOrder() []int {
	// TODO: implement
	return nil
}

// Height returns the number of nodes on the longest root-to-leaf path, or
// zero for an empty tree.
func (a *AVL) Height() int {
	// TODO: implement
	return 0
}
