package offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr, curr1, curr2 := dummy, l1, l2
	for curr1 != nil || curr2 != nil {
		if curr1 == nil {
			curr.Next = curr2
			curr2 = nil
		} else if curr2 == nil {
			curr.Next = curr1
			curr1 = nil
		} else if curr1.Val < curr2.Val {
			curr.Next = curr1
			curr = curr.Next
			curr1 = curr1.Next
		} else {
			curr.Next = curr2
			curr = curr.Next
			curr2 = curr2.Next
		}
	}
	return dummy.Next
}
