package offer

func mirrorTree(root *TreeNode) *TreeNode {
	var mirror func(*TreeNode)
	mirror = func(node *TreeNode) {
		if node == nil {
			return
		}
		mirror((*TreeNode)(node.Left))
		mirror((*TreeNode)(node.Right))
		node.Left, node.Right = node.Right, node.Left
	}

	mirror(root)
	return root
}
