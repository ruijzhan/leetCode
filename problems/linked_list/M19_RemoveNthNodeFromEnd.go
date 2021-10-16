package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	var prev *ListNode
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}

	if prev != nil {
		prev.Next = slow.Next
	} else {
		if slow.Next == nil {
			return nil
		}
		slow.Val = slow.Next.Val
		slow.Next = slow.Next.Next
	}

	return head
}
