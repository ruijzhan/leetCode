package linkedlist

/*
	既然无法得到上一个节点，就将下一个节点的值拷贝过来，然后把下一个节点删除
	复杂度为 O(1)
*/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
