package array

func plusOne(digits []int) []int {
	ln := len(digits)
	digits[ln-1] += 1
	if digits[ln-1] != 10 {
		return digits
	}

	for i := ln - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1] += 1
		}
	}

	if digits[0] == 10 {
		digits = append([]int{1, 0}, digits[1:]...)
	}

	return digits
}
