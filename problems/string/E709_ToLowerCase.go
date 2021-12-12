package string

import "strings"

func toLowerCase(s string) string {
	sb := strings.Builder{}
	sb.Grow(len(s))
	for _, c := range s {
		if c >= 65 && c <= 90 {
			c |= 32
		}
		sb.WriteRune(c)
	}
	return sb.String()
}
