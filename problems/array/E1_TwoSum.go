package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		sub := target - n
		if j, ok := m[sub]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}
