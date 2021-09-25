package linkedlist

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	// 先把 fast 指针往前移动 n 词
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 一起移动 slow 和 fast，使 fast 到达末尾
	var prev *ListNode
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}

	//如果 slow 在链表中间，删除它即可
	if prev != nil {
		prev.Next = slow.Next
	} else { //如果 slow 在开头
		if slow.Next == nil {
			return prev
		}
		slow.Val = slow.Next.Val
		slow.Next = slow.Next.Next
	}

	return head
}
