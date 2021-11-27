package tree

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if val > node.Val {
			if node.Right == nil {
				node.Right = &TreeNode{Val: val}
				return
			}
			dfs(node.Right)
		} else {
			if node.Left == nil {
				node.Left = &TreeNode{Val: val}
				return
			}
			dfs(node.Left)
		}
	}

	dfs(root)
	return root
}
