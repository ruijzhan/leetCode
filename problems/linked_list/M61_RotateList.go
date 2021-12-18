package linkedlist

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		n++
	}
	k = k % n
	if k == 0 {
		return head
	}
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	var prev *ListNode
	for fast != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
	}
	prev.Next = nil
	curr = slow
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = head
	return slow
}
