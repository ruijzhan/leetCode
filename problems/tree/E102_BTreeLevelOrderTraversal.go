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
