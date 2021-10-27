package linkedlist

import "math"

func insertionSortList(head *ListNode) *ListNode {
	dummy := &ListNode{Val: math.MinInt64}
	curr := head
	for curr != nil {
		next := curr.Next
		currSorted := dummy
		for currSorted != nil {
			if curr.Val >= currSorted.Val && (currSorted.Next == nil || curr.Val < currSorted.Next.Val) {
				curr.Next = currSorted.Next
				currSorted.Next = curr
				break
			}
			currSorted = currSorted.Next
		}
		curr = next
	}

	return dummy.Next
}
