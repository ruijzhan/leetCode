package array

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
}

func testReverseString(s []byte) []byte {
	reverseString(s)
	return s
}
