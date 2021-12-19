package bfsdfs

// 深度优先
func numIslands(grid [][]byte) int {
	nr, nc := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(r, c int) {
		grid[r][c] = '0'
		if r-1 >= 0 && grid[r-1][c] == '1' {
			dfs(r-1, c)
		}
		if r+1 < nr && grid[r+1][c] == '1' {
			dfs(r+1, c)
		}
		if c-1 >= 0 && grid[r][c-1] == '1' {
			dfs(r, c-1)
		}
		if c+1 < nc && grid[r][c+1] == '1' {
			dfs(r, c+1)
		}
	}

	ans := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				ans++
				dfs(r, c)
			}
		}
	}

	return ans
}

// 广度优先
func numIslands2(grid [][]byte) int {
	nr, nc := len(grid), len(grid[0])
	ans := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				ans++
				grid[r][c] = '0'
				queue := [][2]int{{r, c}}
				for len(queue) > 0 {
					r, c := queue[0][0], queue[0][1]
					queue = queue[1:]
					if r-1 >= 0 && grid[r-1][c] == '1' {
						queue = append(queue, [2]int{r - 1, c})
						grid[r-1][c] = '0'
					}
					if r+1 < nr && grid[r+1][c] == '1' {
						queue = append(queue, [2]int{r + 1, c})
						grid[r+1][c] = '0'
					}
					if c-1 >= 0 && grid[r][c-1] == '1' {
						queue = append(queue, [2]int{r, c - 1})
						grid[r][c-1] = '0'
					}
					if c+1 < nc && grid[r][c+1] == '1' {
						queue = append(queue, [2]int{r, c + 1})
						grid[r][c+1] = '0'
					}
				}
			}
		}
	}
	return ans
}
