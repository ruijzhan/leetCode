package linkedlist

func middleNode(head *ListNode) *ListNode {

	mid, fast := head, head

	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}

	return mid

}
