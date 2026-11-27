package trees

// [easy] Traversals returns the inorder, preorder, postorder, and
// level-order sequences of a binary tree.
//
// Background. The three depth-first orders differ only in when a node is
// recorded relative to its two subtrees. Inorder records the left
// subtree, then the node, then the right subtree; on a search tree this
// yields sorted values. Preorder records the node before both subtrees;
// it is the order in which a tree is copied or serialized. Postorder
// records both subtrees before the node; it is the order in which a tree
// is freed or evaluated bottom-up. Level order records nodes by depth,
// left to right within a depth, which needs a queue rather than
// recursion.
//
// Definitions:
//   - subtree: a node together with all of its descendants. The subtree of
//     a nil child is empty.
//   - visit: recording a node's value in the output sequence.
//   - inorder: visit the left subtree, then the node, then the right
//     subtree.
//   - preorder: visit the node, then the left subtree, then the right
//     subtree.
//   - postorder: visit the left subtree, then the right subtree, then the
//     node.
//   - level order: visit nodes in increasing depth, and within one depth
//     from left to right.
//   - depth: the number of edges from the root to a node. The root has
//     depth zero.
//
// Steps:
//  1. Inorder: if the node is nil, do nothing. Otherwise traverse the
//     left child, visit the node, then traverse the right child.
//  2. Preorder: if the node is nil, do nothing. Otherwise visit the node,
//     traverse the left child, then traverse the right child.
//  3. Postorder: if the node is nil, do nothing. Otherwise traverse the
//     left child, traverse the right child, then visit the node.
//  4. Level order: if the root is nil, return an empty sequence.
//     Otherwise put the root in a queue. Remove the front node, visit it,
//     and append its left child and then its right child when they are
//     not nil. Repeat until the queue is empty.
//  5. Return the four sequences.
//
// Notes:
//   - An empty tree produces four empty sequences.
//   - The depth-first orders may be produced recursively or with an
//     explicit stack. Recursion depth can reach the tree height.
//
// Expected: O(n) time, O(n) space
// Input: not mutated
func Traversals(root *TreeNode) (inorder, preorder, postorder, levelorder []int) {
	// TODO: implement
	return nil, nil, nil, nil
}
