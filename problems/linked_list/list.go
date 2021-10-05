package linkedlist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
	Prev *ListNode
}

// 递归
func newLinkedList(nums ...int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	return &ListNode{Val: nums[0], Next: newLinkedList(nums[1:]...)}
}

func newLinkedList2(nums ...int) *ListNode {
	head := new(ListNode)
	curr := head
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}
	return head.Next
}

func newDoublyLinkedList(nums ...int) *ListNode {
	return doublyHelper(nil, nums)
}

func doublyHelper(prev *ListNode, nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0], Prev: prev}
	head.Next = doublyHelper(head, nums[1:])

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
	met := make(map[*ListNode]struct{})
	buf := strings.Builder{}
	buf.WriteString("[")
	for head != nil {
		if _, ok := met[head]; ok {
			buf.WriteString("ERROR: ring detected")
			break
		}
		met[head] = struct{}{}
		buf.WriteString(strconv.Itoa(head.Val))
		if head.Next != nil {
			buf.WriteString(",")
		}
		head = head.Next
	}
	buf.WriteString("]")
	return buf.String()
}

func (head *ListNode) DeepCopy() *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	prev := newHead
	curr := head
	for curr != nil {
		node := &ListNode{Val: curr.Val}
		prev.Next = node
		prev = prev.Next
		curr = curr.Next
	}
	return newHead.Next
}
