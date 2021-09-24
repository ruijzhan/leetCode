package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkedList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	prev := head
	for _, n := range nums[1:] {
		prev.Next = &ListNode{Val: n}
		prev = prev.Next
	}
	return head
}
