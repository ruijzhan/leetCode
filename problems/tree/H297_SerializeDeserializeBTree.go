package tree

import (
	"math"
	"strconv"
	"strings"
)

type Codec struct {
	null int
}

func Constructor() Codec {
	return Codec{null: math.MinInt64}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	nums := make([]int, 0)
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node == nil {
				nums = append(nums, c.null)
			} else {
				nums = append(nums, node.Val)
				newNodes = append(newNodes, node.Left, node.Right)
			}
		}
		nodes = newNodes
	}

	if nums[0] == c.null {
		nums = make([]int, 0)
	}

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] != c.null {
			nums = nums[:i+1]
			break
		}
	}
	builder := strings.Builder{}
	builder.WriteRune('[')
	numStrs := make([]string, len(nums))
	for i, n := range nums {
		if n == c.null {
			numStrs[i] = "null"
		} else {
			numStrs[i] = strconv.Itoa(n)
		}
	}
	builder.WriteString(strings.Join(numStrs, ","))
	builder.WriteRune(']')
	return builder.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	data = strings.Replace(data, "[", "", -1)
	data = strings.Replace(data, "]", "", -1)
	numsStr := strings.Split(data, ",")
	if numsStr[0] == "" {
		return nil
	}
	nums := make([]int, len(numsStr))
	for i, ns := range numsStr {
		if ns == "null" {
			nums[i] = c.null
		} else {
			n, err := strconv.Atoi(ns)
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}
	}
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	nodes := []*TreeNode{root}
	nums = nums[1:]
outer:
	for len(nodes) > 0 {
		newNodes := []*TreeNode{}
		for _, node := range nodes {
			if len(nums) > 0 {
				if nums[0] != c.null {
					node.Left = &TreeNode{Val: nums[0]}
					newNodes = append(newNodes, node.Left)
				}
				nums = nums[1:]
			} else {
				break outer
			}
			if len(nums) > 0 {
				if nums[0] != c.null {
					node.Right = &TreeNode{Val: nums[0]}
					newNodes = append(newNodes, node.Right)
				}
				nums = nums[1:]
			} else {
				break outer
			}
		}
		nodes = newNodes
	}

	return root
}

// TODO: 官方解法
