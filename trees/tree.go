package trees

// TreeNode is a binary tree node shared by the tree exercises. Left and
// Right may be nil. Values are not assumed to be unique unless an
// exercise says so.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
