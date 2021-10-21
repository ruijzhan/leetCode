package tree

func invertTree(root *TreeNode) *TreeNode {
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		helper(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	helper(root)
	return root
}
