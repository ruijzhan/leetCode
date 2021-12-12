package binary_search

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		need := target - numbers[i]
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := low + (high-low)>>1
			n := numbers[mid]
			if n == need {
				return []int{i + 1, mid + 1}
			}
			if n < need {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}
	return []int{-1, -1}
}
