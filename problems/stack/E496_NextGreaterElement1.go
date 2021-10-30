package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	nextGreaterNums := make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		num2 := nums2[i]
		for len(stack) > 0 && num2 > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			nextGreaterNums[num2] = stack[len(stack)-1]
		} else {
			nextGreaterNums[num2] = -1
		}
		stack = append(stack, num2)
	}

	res := make([]int, len(nums1))
	for i, num1 := range nums1 {
		res[i] = nextGreaterNums[num1]
	}
	return res
}
