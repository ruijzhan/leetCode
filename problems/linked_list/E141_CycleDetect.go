package linkedlist

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		if slow.Val == fast.Val {
			return true
		}
		fast = fast.Next.Next
	}

	return false
}
