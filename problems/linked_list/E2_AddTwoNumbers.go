package linkedlist

// 空间复杂度 O(1)
// 时间复杂度 O(N)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	add := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = new(ListNode)
		} else if l2 == nil {
			l2 = new(ListNode)
		}

		l1.Val = l1.Val + l2.Val + add
		add = 0
		if l1.Val > 9 {
			add = 1
			l1.Val %= 10
		}

		curr.Next = l1
		curr = curr.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	if add > 0 {
		curr.Next = &ListNode{Val: add}
	}

	return head.Next
}

// 递归
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l1.Val += l2.Val
	if l1.Val >= 10 {
		l1.Val -= 10
		if l1.Next == nil {
			l1.Next = &ListNode{Val: 1}
		} else {
			l1.Next.Val++
			if l2.Next == nil {
				l2.Next = new(ListNode)
			}
		}
	}

	l1.Next = addTwoNumbers2(l1.Next, l2.Next)

	return l1
}
