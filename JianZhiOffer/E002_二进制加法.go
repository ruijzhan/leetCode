package offer

import "strconv"

func addBinary(a, b string) string {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	la := len(a)
	lb := len(b)

	l := max(la, lb)
	ans := ""
	carry := 0
	for i := 0; i < l; i++ {
		if i < la {
			carry += int(a[la-i-1] - '0')
		}
		if i < lb {
			carry += int(b[lb-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		return "1" + ans
	}

	return ans
}
