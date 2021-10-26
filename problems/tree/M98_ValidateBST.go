package tree

import "math"

func isValidBST(root *TreeNode) bool {
	nums := []int{}
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		nums = append(nums, node.Val)
		helper(node.Right)

	}

	helper(root)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			return false
		}
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	var valid func(*TreeNode, int, int) bool
	valid = func(root *TreeNode, min, max int) bool {
		if root == nil {
			return true
		}
		if root.Val >= max || root.Val <= min {
			return false
		}
		return valid(root.Left, min, root.Val) && valid(root.Right, root.Val, max)
	}

	return valid(root, math.MinInt64, math.MaxInt64)
}
