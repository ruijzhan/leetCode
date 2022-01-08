package offer

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head
	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
			break
		}
		prev = curr
		curr = curr.Next
	}
	return dummy.Next
}
