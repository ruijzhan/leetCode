package tree

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	middle := func(left, right *ListNode) *ListNode {
		slow, fast := left, left
		for fast != right && fast.Next != right {
			slow = slow.Next
			fast = fast.Next.Next
		}
		return slow
	}

	var bstBuilder func(*ListNode, *ListNode) *TreeNode
	bstBuilder = func(left, right *ListNode) *TreeNode {
		if left == right {
			return nil
		}
		mid := middle(left, right)
		root := &TreeNode{Val: mid.Val}
		root.Left = bstBuilder(left, mid)
		root.Right = bstBuilder(mid.Next, right)
		return root
	}

	return bstBuilder(head, nil)
}
