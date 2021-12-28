package tree

type BSTIterator struct {
	root *TreeNode
	vals []int
}

func Constructor2(root *TreeNode) BSTIterator {
	t := BSTIterator{
		root: root,
		vals: make([]int, 0),
	}
	t.dfs(root)
	return t
}

func (t *BSTIterator) dfs(node *TreeNode) {
	if node == nil {
		return
	}
	t.dfs(node.Left)
	t.vals = append(t.vals, node.Val)
	t.dfs(node.Right)
}

func (t *BSTIterator) Next() int {
	val := t.vals[0]
	t.vals = t.vals[1:]
	return val
}

func (t *BSTIterator) HasNext() bool {
	return len(t.vals) != 0
}
