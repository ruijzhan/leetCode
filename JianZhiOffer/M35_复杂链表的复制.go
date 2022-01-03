package offer

func copyRandomList(head *Node) *Node {
	cache := make(map[*Node]*Node)
	var deepCopy func(*Node) *Node
	deepCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, has := cache[node]; has {
			return n
		}
		newNode := &Node{Val: node.Val}
		cache[node] = newNode
		newNode.Next = deepCopy(node.Next)
		newNode.Random = deepCopy(node.Random)
		return newNode
	}

	return deepCopy(head)
}
