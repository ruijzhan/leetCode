package array

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_singleNumber(t *testing.T) {
	fs := []func([]int) int{singleNumber, singleNumber2, singleNumber3, singleNumber4}
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			args: args{
				nums: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 7, 7, 8, 8, 9, 9},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.nums, got, tt.want)
				}
			})
		}
	}
}
