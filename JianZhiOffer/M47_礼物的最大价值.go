package offer

func maxValue(grid [][]int) int {
	width, height := len(grid[0]), len(grid)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// 当前状态 grid[i][j] 表示在此处能得到的最大值

			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				// 如果当前状态仍然在第一行，那肯定只有从左边那一格的状态转换而来
				grid[i][j] += grid[i][j-1]
			} else if j == 0 {
				// 如果在第一列，那也只能从上面一格状态转换过来
				grid[i][j] += grid[i-1][j]
			} else {
				// 当在棋盘中间时，当前 i,j 格的值是不变的，而又只能从上方和左边的状态转换，
				// 那么就选择两个可能的之前状态中较大的那个，加上当前格子中的值得到当前状态。
				grid[i][j] += func(a, b int) int {
					if a > b {
						return a
					}
					return b
				}(grid[i-1][j], grid[i][j-1])
			}
		}
	}
	return grid[height-1][width-1]
}
