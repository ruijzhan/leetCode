package tree

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		nums := []int{}
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			nums = append(nums, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		res = append(res, nums)
		nodes = newNodes
	}

	return res
}

// 递归
func levelOrder2(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var helper func([]*TreeNode)
	helper = func(nodes []*TreeNode) {
		if len(nodes) == 0 {
			return
		}
		newNodes := []*TreeNode{}
		nums := []int{}
		for _, node := range nodes {
			nums = append(nums, node.Val)
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		res = append(res, nums)
		helper(newNodes)
	}

	helper([]*TreeNode{root})
	return res
}
