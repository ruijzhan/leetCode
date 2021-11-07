package tree

func diameterOfBinaryTree(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		res = max(res, l+r+1)
		return max(l, r) + 1
	}
	depth(root)
	return res - 1
}
