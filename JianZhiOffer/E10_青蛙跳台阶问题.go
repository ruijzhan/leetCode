package offer

func numWays(n int) int {
	const mod = 1e9 + 7
	a, b, sum := 1, 1, 0
	for i := 0; i < n; i++ {
		sum = (a + b) % mod
		a = b
		b = sum
	}
	return a
}
