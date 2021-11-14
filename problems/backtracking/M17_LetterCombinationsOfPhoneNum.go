package bt

func letterCombinations(digits string) []string {
	ans := []string{}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var helper func(string, string)
	helper = func(nums, letters string) {
		if len(letters) == len(digits) {
			if len(letters) > 0 {
				ans = append(ans, letters)
			}
			return
		}
		for _, letter := range m[nums[0]] {
			helper(nums[1:], letters+letter)
		}
	}

	helper(digits, "")
	return ans
}
