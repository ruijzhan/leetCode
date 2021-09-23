package linkedlist

func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head
	for fast != nil && fast.Next != nil { // 要检查两个指针的有效性
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
