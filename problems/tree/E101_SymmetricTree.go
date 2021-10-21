package tree

func isSymmetric(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		for i := 0; i < l/2; i++ {
			left, right := nodes[i], nodes[l-1-i]
			if left == right {
				continue
			} else if left == nil || right == nil {
				return false
			} else if nodes[i].Val != nodes[l-1-i].Val {
				return false
			}
		}

		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node != nil {
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}
		nodes = newNodes
	}
	return true
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(*TreeNode, *TreeNode) bool
	helper = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val == right.Val {
			return helper(left.Left, right.Right) && helper(left.Right, right.Left)
		} else {
			return false
		}
	}

	return helper(root.Left, root.Right)
}
