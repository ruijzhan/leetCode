package offer

func kthLargest(root *TreeNode, k int) int {
	ans := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 倒序中序遍历二叉树, 先右子树再左子树
		dfs((*TreeNode)(node.Right))
		if k == 1 {
			ans = node.Val
		}
		k--
		dfs((*TreeNode)(node.Left))

	}

	dfs(root)
	return ans
}
