package offer

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currA, currB := headA, headB
	// 如果 currA 指针遍历到了链表结尾，则从 B 链表的头开始继续遍历
	// 对 currB 也一样。这样两个指针会在公共节点处相遇。或者同时到达
	// 各自的末尾。所以结果返回 currA 就行了
	for currA != currB {
		if currA == nil {
			currA = headB
		} else {
			currA = currA.Next
		}

		if currB == nil {
			currB = headA
		} else {
			currB = currB.Next
		}
	}

	return currA
}
