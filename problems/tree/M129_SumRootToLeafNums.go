package tree

func sumNumbers(root *TreeNode) int {
	ans := 0
	var helper func(*TreeNode, int)
	helper = func(node *TreeNode, sum int) {
		sum += node.Val
		if node.Left == nil && node.Right == nil {
			ans += sum
		}
		if node.Left != nil {
			helper(node.Left, sum*10)
		}
		if node.Right != nil {
			helper(node.Right, sum*10)
		}
	}
	helper(root, 0)
	return ans
}
