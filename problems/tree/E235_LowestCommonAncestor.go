package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	path := func(root, node *TreeNode) []*TreeNode {
		res := []*TreeNode{}
		curr := root
		for curr != nil {
			res = append(res, curr)
			if curr.Val > node.Val {
				curr = curr.Left
			} else if curr.Val < node.Val {
				curr = curr.Right
			} else {
				break
			}
		}
		return res
	}
	pathP := path(root, p)
	pathQ := path(root, q)

	var ancestor *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ); i++ {
		if pathP[i] == pathQ[i] {
			ancestor = pathP[i]
		} else {
			break
		}
	}

	return ancestor
}
