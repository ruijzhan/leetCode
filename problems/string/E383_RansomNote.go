package string

func canConstruct(ransomNote string, magazine string) bool {
	letters := [26]int{}
	for _, c := range magazine {
		letters[byte(c)-'a'] += 1
	}
	for _, c := range ransomNote {
		if letters[byte(c)-'a'] == 0 {
			return false
		}
		letters[byte(c)-'a'] -= 1
	}
	return true
}
