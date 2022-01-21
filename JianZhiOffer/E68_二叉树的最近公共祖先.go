package offer

// 如果节点 x 是 节点 p，q 的最近（深度最大）的公共祖先，那么需要满足以下条件：
// 1，x 的左右两个子树，分别包含 p 或者 q 节点。或者
// 2，x 的值等于 p 或者 q，那么只要求其左或者右子树，包含 p 或者 q 节点。
// 要想找到的 x 节点是满足条件的最大深度的节点，采用后序遍历，自底向上遍历二叉树

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 定义出口
	if root == nil {
		return nil
	}
	// 遍历过程中如果遇到 p 或者 q，则公共祖先一定在其上方，没有必要再往下继续遍历
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor((*TreeNode)(root.Left), p, q)
	right := lowestCommonAncestor((*TreeNode)(root.Right), p, q)

	// 满足了第一个条件，左右子树中分别都找到了 p 或者 q
	if left != nil && right != nil {
		return root
	}

	// 因为在开始递归以前就已经检查了节点的值是否等于 p 或者 q，
	// 所以不可能出现左右子树都找不到 p，q的情况
	if left == nil {
		return right
	}

	return left
}
