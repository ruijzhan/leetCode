package tree

func flatten(root *TreeNode) {
	var helper func(*TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		left := helper(node.Left)
		right := helper(node.Right)
		node.Right = left
		node.Left = nil
		curr := node
		for curr != nil && curr.Right != nil {
			curr = curr.Right
		}
		curr.Right = right
		return node
	}

	helper(root)
}
