package dp

func climbStairs(n int) int {
	// 递归 超时
	// if n == 1 || n == 2 {
	// return n
	// }
	// return climbStairs(n-1) + climbStairs(n-2)

	// 参考动画： https://leetcode-cn.com/problems/climbing-stairs/solution/pa-lou-ti-by-leetcode-solution/
	p, q, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r

}
