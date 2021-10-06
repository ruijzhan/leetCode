package array

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_thirdMax(t *testing.T) {
	fs := []func([]int) int{thirdMax, thirdMax2}
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
				nums: []int{3, 2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{1, 2},
			},
			want: 2,
		},
		{
			args: args{
				nums: []int{2, 2, 3, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {

			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v",
						utils.GetFunctionName(f), tt.args.nums, got, tt.want)
				}
			})
		}
	}
}
