package hash

func isIsomorphic(s string, t string) bool {
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		cs := s[i]
		ct := t[i]

		if (ms[cs] > 0 && ms[cs] != ct) || (mt[ct] > 0 && mt[ct] != cs) {
			return false
		}
		ms[cs] = ct
		mt[ct] = cs

	}
	return true
}
