package linkedlist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func newLinkedList(nums ...int) *ListNode {
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

func newRing(n int, nums ...int) *ListNode {
	l := newLinkedList(nums...)
	var nthNode, tail *ListNode
	nth := 0
	curr := l

	for curr != nil {
		if nth == n {
			nthNode = curr
		}
		nth++
		tail = curr
		curr = curr.Next
	}
	tail.Next = nthNode

	return l
}

func (head *ListNode) String() string {
	buf := strings.Builder{}
	buf.WriteString("[")
	for head != nil {
		buf.WriteString(strconv.Itoa(head.Val))
		if head.Next != nil {
			buf.WriteString(",")
		}
		head = head.Next
	}
	buf.WriteString("]")
	return buf.String()
}
