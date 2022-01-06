package offer

func levelOrder3(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	nodes := []*TreeNode{root}
	for level := 0; len(nodes) > 0; level++ {
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
			if level%2 != 0 {
				for i := 0; i < len(row)/2; i++ {
					row[i], row[len(row)-1-i] = row[len(row)-1-i], row[i]
				}
			}
			ans = append(ans, row)
		}
	}
	return
}
