package tree

import "strconv"

func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	var helper func(string, *TreeNode)
	helper = func(path string, node *TreeNode) {
		if path == "" {
			path += strconv.Itoa(node.Val)
		} else {
			path += "->" + strconv.Itoa(node.Val)
		}
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path)
			return
		}
		if node.Left != nil {
			helper(path, node.Left)
		}
		if node.Right != nil {
			helper(path, node.Right)
		}
	}

	helper("", root)
	return ans
}
