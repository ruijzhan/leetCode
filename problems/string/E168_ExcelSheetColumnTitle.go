package string

func convertToTitle(columnNumber int) string {
	res := ""
	for columnNumber > 0 {
		columnNumber--
		res = string(rune('A'+columnNumber%26)) + res
		columnNumber = columnNumber / 26
	}
	return res
}
