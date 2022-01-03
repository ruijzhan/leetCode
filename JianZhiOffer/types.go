package offer

import "math"

var (
	null = math.MinInt64
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func NodeList(nodes [][]int) *Node {
	cache := make(map[int]*Node)
	dummy := new(Node)
	prev := dummy
	for i, tuple := range nodes {
		prev.Next = &Node{
			Val: tuple[0],
		}
		prev = prev.Next
		cache[i] = prev
	}

	for i, curr := 0, dummy.Next; curr != nil; i, curr = i+1, curr.Next {
		if nodes[i][1] != null {
			curr.Random = cache[nodes[i][1]]
		}
	}

	return dummy.Next
}
