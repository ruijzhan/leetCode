package offer

func numWays(n int) int {
	const mod = 1e9 + 7
	a, b := 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a % mod
}
