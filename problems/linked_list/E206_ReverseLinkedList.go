package linkedlist

/*
	迭代法: O(N)
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

//TODO: 用递归方法解题
