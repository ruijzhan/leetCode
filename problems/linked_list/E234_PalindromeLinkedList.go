package linkedlist

/*
	时间复杂度 O(N)
	空间复杂度 O(1)
*/

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	var prev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	if fast != nil {
		slow = slow.Next
	}

	for prev != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}
