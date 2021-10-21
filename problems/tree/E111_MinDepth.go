package tree

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	minD := math.MaxInt64
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}

	return 1 + minD
}
