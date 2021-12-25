package tree

import "math"

func isEvenOddTree(root *TreeNode) bool {
	level := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0, len(nodes)*2)

		var prev int
		odd := level%2 == 0
		if odd {
			prev = math.MinInt64
		} else {
			prev = math.MaxInt64
		}

		for _, node := range nodes {
			if node != nil {
				if odd && node.Val%2 == 0 || !odd && node.Val%2 != 0 || // 检查奇偶
					odd && node.Val <= prev || !odd && node.Val >= prev { // 检查递增/减
					return false
				}
				prev = node.Val
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}

		level++
		nodes = newNodes
	}
	return true
}
