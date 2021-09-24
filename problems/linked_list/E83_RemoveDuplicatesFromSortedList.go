package linkedlist

// 自己的解答，用的双指针
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next

	for fast != nil {
		next := fast.Next
		if slow.Val != fast.Val {
			slow = fast
		} else {
			slow.Next = next
		}
		fast = next
	}

	return head
}

//官方答案，只用curr指针
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
