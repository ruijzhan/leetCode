package tree

func findBottomLeftValue(root *TreeNode) int {
	nodes := []*TreeNode{root}
	var val int
	for len(nodes) > 0 {
		val = nodes[0].Val
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
	return val
}

//TODO: 递归
