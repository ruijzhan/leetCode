package linkedlist

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	ath, bth := list1, list1

	for i := 0; i < b; i++ {
		bth = bth.Next
		if i >= b-a {
			ath = ath.Next
		}
	}

	l2tail := list2
	for l2tail.Next != nil {
		l2tail = l2tail.Next
	}
	l2tail.Next = bth.Next

	ath.Val = list2.Val
	ath.Next = list2.Next

	return list1
}
