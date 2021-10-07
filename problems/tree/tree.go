package tree

import "math"

const (
	null = math.MinInt64
)

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

func (root *TreeNode) slice() []int {
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

	for i := len(vals) - 1; i > 0; i-- {
		if vals[i] != null {
			return vals[:i+1]
		}
	}
	return vals
}
