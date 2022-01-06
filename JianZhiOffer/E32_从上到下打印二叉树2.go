package offer

func levelOrder2(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		row := []int{}
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if node != nil {
				row = append(row, node.Val)
				newNodes = append(newNodes, (*TreeNode)(node.Left), (*TreeNode)(node.Right))
			}
		}
		nodes = newNodes
		if len(row) > 0 {
			ans = append(ans, row)
		}
	}
	return
}
