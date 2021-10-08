package tree

// 层级遍历
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		depth++
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		nodes = newNodes
	}
	return depth
}

// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
