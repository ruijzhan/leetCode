package array

func longestConsecutive(nums []int) int {
	maxLen := 0
	numSet := make(map[int]bool)

	for _, num := range nums {
		numSet[num] = true
	}

	for k := range numSet {
		if !numSet[k-1] {
			currentLen := 1
			for numSet[k+currentLen] {
				currentLen++
			}
			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}

	return maxLen
}
