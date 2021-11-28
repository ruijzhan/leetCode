package tree

func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool)
	nums := []int{}

	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		nums = append(nums, tn.Val)
		inorder(tn.Right)
	}

	inorder(root)

	for _, num := range nums {
		need := k - num
		if m[need] {
			return true
		}
		m[num] = true
	}
	return false
}
