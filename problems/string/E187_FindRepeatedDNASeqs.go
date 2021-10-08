package string

func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0)
	L := 10
	m := make(map[string]int)
	for i := 0; i+L <= len(s); i++ {
		ss := s[i : i+L]
		m[ss]++
		if m[ss] == 2 {
			res = append(res, ss)
		}
	}
	return res
}
