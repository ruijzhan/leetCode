package array

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if len(mat)*len(mat[0]) != r*c {
		return mat
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	for i, row := range mat {
		for j, n := range row {
			nth := (i * len(mat[0])) + j
			k, l := nth/c, nth%c
			res[k][l] = n
		}
	}
	return res
}
