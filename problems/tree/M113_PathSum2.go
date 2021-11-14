package tree

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, target int, path []int) {
		path = append(path, node.Val)
		path2 := make([]int, len(path))
		copy(path2, path)
		target -= node.Val
		if target == 0 {
			if node.Left == nil && node.Right == nil {
				ans = append(ans, path2)
			}
		}
		if node.Left != nil {
			dfs(node.Left, target, path)
		}
		if node.Right != nil {
			dfs(node.Right, target, path)
		}
	}

	dfs(root, targetSum, []int{})
	return ans
}
