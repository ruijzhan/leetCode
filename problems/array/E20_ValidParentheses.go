package array

func isValid(s string) bool {
	lefts := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	rights := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}

	for _, b := range []byte(s) {
		if _, ok := lefts[b]; ok {
			stack = append(stack, b)
		} else {
			expect := rights[b]
			l := len(stack)
			if l == 0 {
				return false
			}
			if stack[l-1] != expect {
				return false
			}
			stack = stack[:l-1]
		}
	}
	return len(stack) == 0
}
