package string

import "strings"

func findWords(words []string) []string {
	mKeyRow := func() map[byte]int {
		rows := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
		m := make(map[byte]int, 26*2)
		for i, row := range rows {
			for _, key := range []byte(row + strings.ToUpper(row)) {
				m[key] = i
			}
		}
		return m
	}()

	n := 0
nextWord:
	for _, word := range words {
		row := mKeyRow[word[0]]
		for i := 1; i < len(word); i++ {
			if mKeyRow[word[i]] != row {
				continue nextWord
			}
		}
		words[n] = word
		n++
	}

	return words[:n]
}
