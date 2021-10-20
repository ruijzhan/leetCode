package tree

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		helper(node.Left)
		helper(node.Right)

	}
	helper(root)
	return res
}
