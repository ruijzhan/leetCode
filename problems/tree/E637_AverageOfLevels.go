package tree

func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}
	level := []*TreeNode{root}
	for len(level) != 0 {
		sum := 0
		newLevel := []*TreeNode{}
		for _, node := range level {
			sum += node.Val
			if node.Left != nil {
				newLevel = append(newLevel, node.Left)
			}
			if node.Right != nil {
				newLevel = append(newLevel, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(level)))
		level = newLevel
	}
	return res
}
