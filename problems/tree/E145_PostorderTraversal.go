package tree

func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		helper(node.Right)
		res = append(res, node.Val)
	}

	helper(root)
	return res
}
