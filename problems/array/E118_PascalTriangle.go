package array

func generate(numRows int) [][]int {
	res := [][]int{}
	var lastRow []int
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		if i == 0 {
			row[0] = 1
		} else {
			lastRow = res[i-1]
			row[0], row[len(row)-1] = 1, 1
			for j := 1; j < len(row)-1; j++ {
				row[j] = lastRow[j-1] + lastRow[j]
			}
		}
		res = append(res, row)
	}
	return res
}
