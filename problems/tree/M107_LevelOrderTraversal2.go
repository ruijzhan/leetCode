package tree

func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		vals := make([]int, 0, len(nodes))
		newNodes := make([]*TreeNode, 0, len(nodes))
		for _, node := range nodes {
			if node != nil {
				vals = append(vals, node.Val)
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}
		nodes = newNodes
		if len(vals) > 0 {
			ans = append(ans, vals)
		}
	}
	ansL := len(ans)
	for i := 0; i < ansL/2; i++ {
		ans[i], ans[ansL-1-i] = ans[ansL-1-i], ans[i]
	}

	return ans
}
