package binary_search

func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for v, c := range m {
		if c > 1 {
			return v
		}
	}
	return 0
}

//TODO: 二分查找，快慢指针
