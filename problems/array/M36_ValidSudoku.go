package array

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		mRow := make(map[byte]bool)
		for j := 0; j < len(board[0]); j++ {
			if i%3 == 0 && j%3 == 0 {
				m9 := make([]bool, 9)
				for k := i; k < i+3; k++ {
					for l := j; l < j+3; l++ {
						if board[k][l] == '.' {
							continue
						}
						if m9[board[k][l]-'1'] {
							return false
						} else {
							m9[board[k][l]-'1'] = true
						}
					}
				}
			}

			//列
			if i == 0 {
				mCol := make([]bool, 9)
				for k := 0; k < len(board); k++ {
					if board[k][j] == '.' {
						continue
					}
					if mCol[board[k][j]-'1'] {
						return false
					} else {
						mCol[board[k][j]-'1'] = true
					}
				}
			}
			//行
			if board[i][j] == '.' {
				continue
			}
			if mRow[board[i][j]] {
				return false
			} else {
				mRow[board[i][j]] = true
			}

		}
	}
	return true
}
