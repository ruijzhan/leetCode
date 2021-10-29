package linkedlist

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		left, right := head, head.Next
		if left.Val > right.Val {
			left.Val, right.Val = right.Val, left.Val
		}
		return head
	}

	merge := func(l, r *ListNode) *ListNode {
		dummy := new(ListNode)
		curr := dummy
		for l != nil || r != nil {
			if l != nil && r != nil {
				if l.Val < r.Val {
					curr.Next = l
					l = l.Next
				} else {
					curr.Next = r
					r = r.Next
				}
			} else if l != nil {
				curr.Next = l
				l = l.Next
			} else {
				curr.Next = r
				r = r.Next
			}
			curr = curr.Next
		}
		return dummy.Next
	}

	cut := func(head *ListNode) (*ListNode, *ListNode) {
		slow, fast := head, head
		var prev *ListNode
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}
		prev.Next = nil
		return head, slow
	}

	left, right := cut(head)
	return merge(sortList(left), sortList(right))
}
