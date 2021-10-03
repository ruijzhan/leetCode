package array

import (
	"reflect"
	"testing"
)

func mergeTest(A []int, m int, B []int, n int, f func([]int, int, []int, int)) []int {
	f(A, m, B, n)
	return A
}
func Test_mergeTest(t *testing.T) {
	type args struct {
		A []int
		m int
		B []int
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				A: []int{1, 2, 3, 0, 0, 0},
				m: 3,
				B: []int{2, 5, 6},
				n: 3,
			},
			want: []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copyA := make([]int, len(tt.args.A))
			copyB := make([]int, len(tt.args.B))

			copy(copyA, tt.args.A)
			copy(copyB, tt.args.B)
			if got := mergeTest(copyA, tt.args.m, copyB, tt.args.n, merge); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTest() = %v, want %v", got, tt.want)
			}

			copy(copyA, tt.args.A)
			copy(copyB, tt.args.B)
			if got := mergeTest(copyA, tt.args.m, copyB, tt.args.n, merge2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTest2() = %v, want %v", got, tt.want)
			}

			copy(copyA, tt.args.A)
			copy(copyB, tt.args.B)
			if got := mergeTest(copyA, tt.args.m, copyB, tt.args.n, merge3); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTest3() = %v, want %v", got, tt.want)
			}
		})
	}
}
