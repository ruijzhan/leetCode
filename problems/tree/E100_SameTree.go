package tree

// 递归
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if (p == nil || q == nil) || p.Val != q.Val {
		return false
	} else {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
}

// 层序
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		if l != nil && r != nil {
			if l.Val != r.Val {
				return false
			}
			queue = append(queue, l.Left, r.Left, l.Right, r.Right)
		} else if l != nil || r != nil {
			return false
		}
		queue = queue[2:]
	}
	return true
}
