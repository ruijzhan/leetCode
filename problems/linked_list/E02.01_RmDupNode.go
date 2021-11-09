package linkedlist

func removeDuplicateNodes(head *ListNode) *ListNode {
	m := make(map[int]struct{})
	curr := head
	for curr != nil && curr.Next != nil {
		m[curr.Val] = struct{}{}
		if _, ok := m[curr.Next.Val]; ok {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

// 空间复杂度 O(1)
func removeDuplicateNodes2(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		curr2 := curr
		for curr2 != nil && curr2.Next != nil {
			if curr.Val == curr2.Next.Val {
				curr2.Next = curr2.Next.Next
			} else {
				curr2 = curr2.Next
			}
		}
		curr = curr.Next
	}
	return head
}
