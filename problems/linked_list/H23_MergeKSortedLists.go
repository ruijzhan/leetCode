package linkedlist

func mergeKLists2(lists []*ListNode) *ListNode {

	mergeTwoLists := func(l1 *ListNode, l2 *ListNode) *ListNode {
		var head *ListNode = &ListNode{}

		curr := head
		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				curr.Next = l1
				l1 = l1.Next
			} else {
				curr.Next = l2
				l2 = l2.Next
			}
			curr = curr.Next
		}

		if l1 != nil {
			curr.Next = l1
		} else {
			curr.Next = l2
		}

		return head.Next
	}

	var ans *ListNode

	for _, l := range lists {
		ans = mergeTwoLists(ans, l)
	}

	return ans
}

// 分治
func mergeKLists(lists []*ListNode) *ListNode {
	merge2Lists := func(l1, l2 *ListNode) *ListNode {
		dummy := new(ListNode)
		curr := dummy
		for l1 != nil || l2 != nil {
			if l1 != nil && l2 != nil {
				if l1.Val < l2.Val {
					curr.Next = l1
					l1 = l1.Next
				} else {
					curr.Next = l2
					l2 = l2.Next
				}
			} else if l1 == nil {
				curr.Next = l2
				l2 = l2.Next
			} else {
				curr.Next = l1
				l1 = l1.Next
			}
			curr = curr.Next
		}
		return dummy.Next
	}

	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return merge2Lists(lists[0], lists[1])
	}

	mid := len(lists) >> 1
	left := lists[:mid]
	right := lists[mid:]

	return merge2Lists(mergeKLists(left), mergeKLists(right))
}

// TODO: 有序队列
