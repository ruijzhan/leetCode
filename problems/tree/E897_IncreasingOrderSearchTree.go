package tree

func increasingBST3(root *TreeNode) *TreeNode {
	dummy := new(TreeNode)
	prev := dummy

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		node.Left = nil
		prev.Right = node
		prev = node
		inorder(node.Right)
	}

	inorder(root)

	return dummy.Right
}
