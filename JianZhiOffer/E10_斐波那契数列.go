package offer

func fib(n int) int {
	const mod = 1e9 + 7
	a, b, sum := 0, 1, 0
	for i := 0; i < n; i++ {
		// 状态转移
		sum = (a + b) % mod
		a = b
		b = sum
	}
	return a
}
