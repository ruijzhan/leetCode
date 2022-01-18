package offer

func isBalanced(root *TreeNode) bool {
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	var hight func(*TreeNode) int
	hight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		hL := hight((*TreeNode)(node.Left))
		hR := hight((*TreeNode)(node.Right))

		if hL == -1 || hR == -1 || abs(hL-hR) > 1 {
			// 用高度 -1 来指代树为非平衡二叉树
			return -1
		}

		// 高度 += 1
		return 1 + max(hL, hR)
	}

	return hight(root) != -1
}
