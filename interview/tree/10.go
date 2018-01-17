package tree

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func NewTreeNode(val int, l, r *TreeNode) *TreeNode {
	return &TreeNode{val: val, left: l, right: r}
}

func PrintTree(root *TreeNode) (data [][]int) {
	return
}


