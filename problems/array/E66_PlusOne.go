package array

func plusOne(digits []int) []int {
	ln := len(digits)
	digits[ln-1] += 1

	for i := ln - 1; i > 0; i-- {
		if digits[i] < 10 {
			return digits
		} else {
			digits[i] -= 10
			digits[i-1] += 1
		}
	}

	if digits[0] > 9 {
		digits[0] -= 10
		digits = append([]int{1}, digits...)
	}

	return digits
}
