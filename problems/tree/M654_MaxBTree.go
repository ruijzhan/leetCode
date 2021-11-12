package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	max := func(nums []int) int {
		var iMax int
		currMax := nums[0]
		for i := 0; i < len(nums); i++ {
			if nums[i] >= currMax {
				iMax = i
				currMax = nums[i]
			}
		}
		return iMax
	}

	iMax := max(nums)
	root := &TreeNode{Val: nums[iMax]}
	root.Left = constructMaximumBinaryTree(nums[:iMax])
	root.Right = constructMaximumBinaryTree(nums[iMax+1:])

	return root
}
