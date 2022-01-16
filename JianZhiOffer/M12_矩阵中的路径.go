package offer

func exist(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	var dfs func([][]byte, int, int, []byte, int) bool
	dfs = func(board [][]byte, i, j int, word []byte, k int) bool {
		// 如果越界了，查找失败
		if i >= row || j >= col || i < 0 || j < 0 {
			return false
		}
		// 如果一开始就不匹配，查找失败
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		// 找过的就设置为0
		board[i][j] = 0
		// 上下左右继续 dfs
		res := dfs(board, i+1, j, word, k+1) ||
			dfs(board, i-1, j, word, k+1) ||
			dfs(board, i, j+1, word, k+1) ||
			dfs(board, i, j-1, word, k+1)

		// 这一句最巧妙，用完了还原
		board[i][j] = word[k]

		return res
	}

	// 从任何可能的起点开始 dfs
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dfs(board, i, j, []byte(word), 0) {
				return true
			}
		}
	}

	return false
}
