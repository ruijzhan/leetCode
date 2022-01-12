package offer

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return 1 + max(maxDepth((*TreeNode)(root.Left)), maxDepth((*TreeNode)(root.Right)))
}
