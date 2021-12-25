package tree

import "math"

func isEvenOddTree(root *TreeNode) bool {
	level := 0
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}

		var prev int
		odd := level%2 == 0
		if odd {
			prev = math.MinInt64
		} else {
			prev = math.MaxInt64
		}

		for _, node := range nodes {
			// 检查 奇偶
			if (odd && node.Val%2 == 0) || (!odd && node.Val%2 != 0) {
				return false
			}

			// 检查递增/递减
			if odd {
				if node.Val > prev {
					prev = node.Val
				} else {
					return false
				}
			} else {
				if node.Val < prev {
					prev = node.Val
				} else {
					return false
				}
			}

			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}

		level++
		nodes = newNodes
	}
	return true
}
