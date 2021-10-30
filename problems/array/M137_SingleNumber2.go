package array

func singleNumber5(nums []int) int {
	m := make(map[int]struct{})
	sum := 0
	for _, num := range nums {
		m[num] = struct{}{}
		sum += num
	}

	keySum := 0
	for k := range m {
		keySum += k
	}

	return (keySum*3 - sum) / 2
}

//TODO: 位运算
