package linkedlist

/*
	迭代法: O(N)
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// 递归方法:
// f(head) = f(head.Next) + 反转操作
// 反转操作:  1,head 的下一个节点的 next 指向 head
//           2,head 的 next 指向 nil
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
