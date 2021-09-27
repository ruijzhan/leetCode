package linkedlist

// 空间复杂度 O(1)
// 时间复杂度 O(N)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	curr := head

	add := 0
	for l1 != nil && l2 != nil {
		l1.Val = l1.Val + l2.Val + add
		if l1.Val >= 10 {
			add = 1
			l1.Val -= 10
		} else {
			add = 0
		}
		curr.Next = l1
		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next

	}
	var rest *ListNode
	if l1 != nil {
		rest = l1
	} else {
		rest = l2
	}

	for rest != nil {
		rest.Val = rest.Val + add
		if rest.Val >= 10 {
			add = 1
			rest.Val -= 10
		} else {
			add = 0
		}
		curr.Next = rest
		curr = curr.Next
		rest = rest.Next
	}

	if add == 1 {
		curr.Next = &ListNode{Val: 1}
	}

	return head.Next
}
