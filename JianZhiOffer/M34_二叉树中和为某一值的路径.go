package offer

func pathSum(root *TreeNode, target int) (ans [][]int) {
	path := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, remain int) {
		if node == nil {
			return
		}
		remain -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		if node.Left == nil && node.Right == nil && remain == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}
		dfs((*TreeNode)(node.Left), remain)
		dfs((*TreeNode)(node.Right), remain)
	}

	dfs(root, target)
	return
}
