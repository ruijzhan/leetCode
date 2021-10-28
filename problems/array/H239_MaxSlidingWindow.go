package array

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	queue := []int{}
	res := make([]int, len(nums)-k+1)

	for i, num := range nums {
		for len(queue) > 0 && num > queue[len(queue)-1] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, num)

		if i >= k && nums[i-k] == queue[0] {
			queue = queue[1:]
		}

		if i >= k-1 {
			res[i-k+1] = queue[0]
		}
	}

	return res
}
