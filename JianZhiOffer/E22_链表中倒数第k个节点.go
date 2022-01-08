package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
