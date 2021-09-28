package linkedlist

func mergeKLists(lists []*ListNode) *ListNode {

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

// TODO: 分治

// TODO: 有序队列
