package offer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	nRows, nColumns := len(matrix[0]), len(matrix)
	r, c := 0, nColumns-1
	for r < nRows && c >= 0 {
		n := matrix[r][c]
		if n == target {
			return true
		} else if n > target {
			c--
		} else {
			r++
		}
	}
	return false
}
