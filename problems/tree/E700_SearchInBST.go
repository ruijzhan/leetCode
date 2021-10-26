package tree

// 递归
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val > root.Val {
		return searchBST(root.Right, val)
	} else {
		return searchBST(root.Left, val)
	}
}

// 迭代
func searchBST2(root *TreeNode, val int) *TreeNode {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		if node != nil {
			if node.Val == val {
				return node
			}
			queue = append(queue, node.Left, node.Right)
		}
		queue = queue[1:]
	}
	return nil
}
