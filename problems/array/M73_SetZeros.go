package array

func setZeroes(matrix [][]int) {
	mr := make([]bool, len(matrix))
	mc := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				mr[i] = true
				mc[j] = true
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if mr[i] || mc[j] {
				matrix[i][j] = 0
			}
		}
	}
}
