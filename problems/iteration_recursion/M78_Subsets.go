package iteration_recursion

// 枚举和迭代
func subsets(nums []int) [][]int {
	ans := [][]int{}
	for mask := 0; mask < 1<<len(nums); mask++ { // 注意此处退出循环条件
		subset := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 { // 注意此处条件
				subset = append(subset, v)
			}
		}
		ans = append(ans, subset)
	}
	return ans
}

// TODO: 递归
func subsets2(nums []int) [][]int {
	//ans := [][]int{}
	ans := subsets(nums)
	return ans
}
