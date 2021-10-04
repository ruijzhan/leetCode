package array

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_rotateTest(t *testing.T) {

	fs := []func([]int, int){rotate, rotate2, rotate3}

	testF := func(nums []int, k int, f func([]int, int)) []int {
		cp := make([]int, len(nums))
		copy(cp, nums)
		f(cp, k)
		return cp
	}
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7},
				k:    3,
			},
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			args: args{
				nums: []int{-1, -100, 3, 99},
				k:    2,
			},
			want: []int{3, 99, -1, -100},
		},
		{
			args: args{
				nums: []int{-1},
				k:    2,
			},
			want: []int{-1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, f := range fs {
				if got := testF(tt.args.nums, tt.args.k, f); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%v, %v) = %v, want %v", utils.GetFunctionName(f), tt.args.nums, tt.args.k, got, tt.want)
				}
			}
		})
	}
}
