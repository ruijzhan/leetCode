package tree

func increasingBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}

	inorder(root)
	newRoot := new(TreeNode)
	curr := newRoot
	for _, n := range nums {
		curr.Right = &TreeNode{Val: n}
		curr = curr.Right
	}
	return newRoot.Right
}

// 中序遍历。将当前节点的左孩子设置为 nil，然后设置为上一个节点的右孩子
func increasingBST2(root *TreeNode) *TreeNode {
	newRoot := new(TreeNode)
	prevNode := newRoot
	var inorder func(*TreeNode)

	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		node.Left = nil
		prevNode.Right = node
		prevNode = node

		inorder(node.Right)
	}

	inorder(root)
	return newRoot.Right
}
