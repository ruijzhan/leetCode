package tree

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		max := func(x, y int) int {
			if x > y {
				return x
			}
			return y
		}
		return 1 + max(height(node.Left), height(node.Right))
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	if abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	}
	return true && isBalanced(root.Left) && isBalanced(root.Right)
}

func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
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

		hL := height(node.Left)
		hR := height(node.Right)
		if hL == -1 || hR == -1 || abs(hL-hR) > 1 {
			return -1
		}
		return 1 + max(hL, hR)
	}

	return height(root) != -1
}
