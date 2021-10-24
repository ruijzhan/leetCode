package tree

func balanceBST(root *TreeNode) *TreeNode {
	nums := []int{}
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	var makeBST func(int, int) *TreeNode
	makeBST = func(low, high int) *TreeNode {
		if low > high {
			return nil
		}
		mid := low + (high-low)>>1
		root := &TreeNode{Val: nums[mid]}
		root.Left = makeBST(low, mid-1)
		root.Right = makeBST(mid+1, high)
		return root
	}

	return makeBST(0, len(nums)-1)
}

//TODO: 官方解法：贪心构造
