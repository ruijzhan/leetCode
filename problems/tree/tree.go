package tree

import (
	"math"

	linkedlist "github.com/ruijzhan/leetCode/problems/linked_list"
)

const (
	null = math.MinInt64
)

type ListNode = linkedlist.ListNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	treeBuilder([]*TreeNode{root}, vals[1:])

	return root
}

func treeBuilder(nodes []*TreeNode, vals []int) {
	if len(nodes) == 0 {
		if len(vals) > 0 {
			panic("invalid input")
		}
		return
	}

	newNode := func(vals []int) (*TreeNode, []int) {
		if len(vals) == 0 {
			return nil, vals
		}
		val, vals := vals[0], vals[1:]
		if val == null {
			return nil, vals
		} else {
			return &TreeNode{Val: val}, vals
		}
	}

	newNodes := make([]*TreeNode, 0)
	for _, node := range nodes {
		if node == nil {
			continue
		}

		node.Left, vals = newNode(vals)
		newNodes = append(newNodes, node.Left)

		node.Right, vals = newNode(vals)
		newNodes = append(newNodes, node.Right)
	}

	treeBuilder(newNodes, vals)
}

func (root *TreeNode) levelOrder() []int {
	processing := []*TreeNode{root}
	vals := make([]int, 0)
	for ; len(processing) > 0; processing = processing[1:] {
		node := processing[0]
		if node == nil {
			vals = append(vals, null)
		} else {
			vals = append(vals, node.Val)
			processing = append(processing, node.Left, node.Right)
		}
	}

	if vals[0] == null {
		return []int{}
	}

	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] != null {
			return vals[:i+1]
		}
	}
	return vals
}

func (root *TreeNode) inOrder() []int {
	vals := make([]int, 0)
	var helper func(*TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		vals = append(vals, node.Val)
		helper(node.Right)
	}
	helper(root)
	return vals
}
