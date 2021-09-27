package linkedlist

// 时间复杂度 O(N)
// 空间复杂度 O(N)
func detectCycle(head *ListNode) *ListNode {
	m := make(map[int]struct{})
	for head != nil {
		if _, ok := m[head.Val]; ok {
			return head
		}
		m[head.Val] = struct{}{}
		head = head.Next
	}
	return nil
}

//TODO: 用 双指针 方法完成
