package tree

// 层级迭代
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	process := func(nodes []*TreeNode) []*TreeNode {
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node == nil {
				continue
			}
			newNodes = append(newNodes, node.Left, node.Right)
		}
		return newNodes
	}
	nodes := []*TreeNode{root}
	i := 0
	for {
		nodes = process(nodes)
		if len(nodes) > 0 {
			i++
		} else {
			break
		}
	}

	return i
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
