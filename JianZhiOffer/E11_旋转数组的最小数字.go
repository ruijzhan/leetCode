package offer

func minArray(numbers []int) int {
	low, high := 0, len(numbers)-1
	mid := -1
	for low <= high {
		mid = low + (high-low)>>1
		if numbers[low] > numbers[mid] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else if numbers[low] == numbers[mid] || numbers[low] == numbers[high] {
			high--
		} else {
			high = mid - 1
		}
	}
	return numbers[mid]
}
