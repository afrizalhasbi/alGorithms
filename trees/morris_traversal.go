package trees

// [medium] MorrisInOrder returns the inorder sequence of a binary tree
// using constant extra space and then restores the tree.
//
// Background. Ordinary inorder traversal needs a stack or recursion to
// remember the path back up the tree, which costs space proportional to
// the height. Morris traversal removes that cost by temporarily
// redirecting the right link of each node's inorder predecessor to the
// node itself. The link acts as the return path. When the traversal
// reaches the predecessor a second time, the node's left subtree is
// finished, so the link is removed and the node is visited. Every
// temporary link is removed before the traversal ends, leaving the tree
// exactly as it was.
//
// Definitions:
//   - inorder predecessor: for a node with a left child, the rightmost
//     node in its left subtree. It is the node visited immediately before
//     the node in an inorder traversal.
//   - thread: a temporary right link from a predecessor back to the node
//     whose left subtree is being traversed.
//   - current: the node the traversal is examining.
//
// Steps:
//  1. Set current to the root and create an empty output sequence.
//  2. While current is not nil, if current has no left child, visit
//     current and move current to its right child.
//  3. Otherwise, find the rightmost node in current's left subtree. If
//     that node's right child is nil, set its right child to current and
//     move current to its left child. If that node's right child is
//     current, the left subtree has been finished: set its right child
//     back to nil, visit current, and move current to its right child.
//  4. Return the output sequence.
//
// Notes:
//   - The tree is temporarily modified but is fully restored before the
//     function returns.
//   - The function is not safe to run concurrently with any other reader
//     of the same tree.
//   - Every edge is walked a constant number of times.
//
// Expected: O(n) time, O(1) space
// Input: mutated temporarily and restored
func MorrisInOrder(root *TreeNode) []int {
	// TODO: implement
	return nil
}
