package offer

func reverseLeftWords(s string, k int) string {
	if k == 0 {
		return s
	}
	reverse := func(sb []byte) {
		l := len(sb)
		for i := 0; i < l/2; i++ {
			sb[i], sb[l-i-1] = sb[l-i-1], sb[i]
		}
	}
	k = k % len(s)
	sb := []byte(s)
	l := len(sb)
	reverse(sb)
	reverse(sb[:l-k])
	reverse(sb[l-k:])
	return string(sb)
}
