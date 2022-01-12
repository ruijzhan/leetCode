package offer

import (
	"strings"
)

func reverseWords(s string) string {
	subs := strings.Split(s, " ")
	words := make([]string, 0, len(subs))
	for _, sub := range subs {
		if sub != "" {
			words = append(words, sub)
		}
	}
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}
