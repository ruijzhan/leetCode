package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	dummyOdd, dummyEven := new(ListNode), new(ListNode)
	currOdd, currEven := dummyOdd, dummyEven
	curr := head
	for i := 0; curr != nil; i++ {
		if i%2 == 0 { //odd
			currOdd.Next = curr
			currOdd = currOdd.Next
		} else {
			currEven.Next = curr
			currEven = currEven.Next
		}
		curr = curr.Next
	}
	currEven.Next = nil
	currOdd.Next = dummyEven.Next
	return dummyOdd.Next
}
