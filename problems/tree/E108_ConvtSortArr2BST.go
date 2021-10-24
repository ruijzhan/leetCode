package tree

// 递归
func sortedArrayToBST(nums []int) *TreeNode {
	var helper func(low, high int) *TreeNode
	helper = func(low, high int) *TreeNode {
		if low > high {
			return nil
		}
		mid := low + (high-low)>>1
		root := &TreeNode{Val: nums[mid]}
		root.Left = helper(low, mid-1)
		root.Right = helper(mid+1, high)
		return root
	}
	return helper(0, len(nums)-1)
}
