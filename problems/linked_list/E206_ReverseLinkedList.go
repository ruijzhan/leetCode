package linkedlist

/*
	迭代法: O(N)
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		// 赋值转一个圈
		next := curr.Next // 用 next 保存下一个节点的指针
		curr.Next = prev  // 把当前节点的 Next 指针指向 "上一个节点"
		prev = curr       // 把上一个节点换成当前节点
		curr = next       // 把当前节点换成下一个节点
	}

	return prev
}

//TODO: 用递归方法解题
