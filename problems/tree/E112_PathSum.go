package tree

// 递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

//TODO: 层序遍历
//func hasPathSum2(root *TreeNode, targetSum int) bool {
//	return false
//}
