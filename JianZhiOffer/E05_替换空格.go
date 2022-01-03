package offer

func replaceSpace(s string) string {
	ans := make([]byte, len(s)*3)
	i := 0
	for _, c := range s {
		if c == ' ' {
			ans[i] = '%'
			ans[i+1] = '2'
			ans[i+2] = '0'
			i += 3
		} else {
			ans[i] = byte(c)
			i++
		}
	}
	return string(ans[:i])
}
