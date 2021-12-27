package array

// 归并排序
func sortArray(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	lnums := len(nums)
	low := sortArray(nums[:lnums/2])
	high := sortArray(nums[lnums/2:])
	sorted := make([]int, lnums)

	for i, j, k := len(low)-1, len(high)-1, lnums-1; i >= 0 || j >= 0; k-- {
		if i < 0 {
			sorted[k] = high[j]
			j--
		} else if j < 0 {
			sorted[k] = low[i]
			i--
		} else if low[i] > high[j] {
			sorted[k] = low[i]
			i--
		} else {
			sorted[k] = high[j]
			j--
		}
	}

	return sorted
}

//TODO: 其他排序方法
