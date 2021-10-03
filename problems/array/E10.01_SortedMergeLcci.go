package array

import (
	"sort"
)

func merge(A []int, m int, B []int, n int) {
	for ia, ib, i := m-1, n-1, m+n-1; i >= 0; i-- {
		if ia < 0 {
			A[i] = B[ib]
			ib--
		} else if ib < 0 {
			A[i] = A[ia]
			ia--
		} else if A[ia] >= B[ib] {
			A[i] = A[ia]
			ia--
		} else {
			A[i] = B[ib]
			ib--
		}
	}
}

func merge2(A []int, m int, B []int, n int) {
	sorted := make([]int, m+n)

	for ia, ib, i := 0, 0, 0; i < m+n; i++ {
		if ia == m {
			sorted[i] = B[ib]
			ib++
		} else if ib == n {
			sorted[i] = A[ia]
			ia++
		} else if A[ia] <= B[ib] {
			sorted[i] = A[ia]
			ia++
		} else {
			sorted[i] = B[ib]
			ib++
		}
	}

	copy(A, sorted)
}

func merge3(A []int, m int, B []int, n int) {
	copy(A[m:], B)
	sort.Ints(A)
}
