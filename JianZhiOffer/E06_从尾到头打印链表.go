package offer

func reversePrint(head *ListNode) []int {
	ans := []int{}
	for ; head != nil; head = head.Next {
		ans = append(ans, head.Val)
	}
	l := len(ans)
	for i := 0; i < l/2; i++ {
		ans[i], ans[l-i-1] = ans[l-i-1], ans[i]
	}
	return ans
}
