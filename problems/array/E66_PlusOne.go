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

// 递归
func plusOne2(digits []int) []int {
	var helper func([]int) int
	helper = func(digits []int) int {
		if len(digits) == 0 {
			return 0
		}

		digits[0] += helper(digits[1:])
		if digits[0] > 9 {
			digits[0] -= 10
			return 1
		}
		return 0
	}

	digits[len(digits)-1] += 1

	if helper(digits) == 1 {
		newDigits := make([]int, 1, len(digits)+1)
		newDigits[0] = 1
		return append(newDigits, digits...)
	}
	return digits
}
