package offer

func levelOrder(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node != nil {
				ans = append(ans, node.Val)
				newNodes = append(newNodes, (*TreeNode)(node.Left), (*TreeNode)(node.Right))
			}
		}
		nodes = newNodes
	}

	return
}
