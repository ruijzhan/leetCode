package linkedlist

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	curr := dummy
	for curr.Next != nil && curr.Next.Next != nil {
		node1 := curr.Next
		node2 := curr.Next.Next
		curr.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		curr = node1
	}
	return dummy.Next
}

//TODO: 递归
