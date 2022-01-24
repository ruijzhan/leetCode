package offer

func hammingWeight(num uint32) int {
	ones := 0
	for i := 0; i < 32; i++ {
		// 1<<i 只有第 i 位上是 1
		if 1<<i&num > 0 {
			ones++
		}
	}
	return ones
}
