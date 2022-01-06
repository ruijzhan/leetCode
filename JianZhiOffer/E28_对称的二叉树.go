package offer

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}

		return compare((*TreeNode)(l.Left), (*TreeNode)(r.Right)) &&
			compare((*TreeNode)(l.Right), (*TreeNode)(r.Left))

	}
	return compare((*TreeNode)(root.Left), (*TreeNode)(root.Right))
}
