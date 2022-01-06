package offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	var compare func(*TreeNode, *TreeNode) bool
	compare = func(a, b *TreeNode) bool {
		//判断b是否为a的子树，如果b为空
		if b == nil {
			return true
		}
		// 如果a为空或者值不相等
		if a == nil || a.Val != b.Val {
			return false
		}
		// 继续递归比较 a b 的左右子树
		return compare((*TreeNode)(a.Left), (*TreeNode)(b.Left)) && compare((*TreeNode)(a.Right), (*TreeNode)(b.Right))
	}

	return compare(A, B) || isSubStructure((*TreeNode)(A.Left), B) || isSubStructure((*TreeNode)(A.Right), B)
}
