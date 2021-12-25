package iteration_recursion

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			args: args{
				nums: []int{0},
			},
			want: [][]int{{}, {0}},
		},
	}

	deepEqual := func(s1, s2 [][]int) bool {
	outter:
		for _, ss1 := range s1 {
			for _, ss2 := range s2 {
				if reflect.DeepEqual(ss1, ss2) {
					continue outter
				}
			}
			return false
		}
	outter1:
		for _, ss2 := range s2 {
			for _, ss1 := range s1 {
				if reflect.DeepEqual(ss1, ss2) {
					continue outter1
				}
			}
			return false
		}

		return true
	}
	for _, tt := range tests {
		for _, f := range []func([]int) [][]int{subsets, subsets2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); !deepEqual(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.nums, got, tt.want)
				}
			})
		}
	}
}
