package array

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_singleNumber5(t *testing.T) {
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
				nums: []int{2, 2, 2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{0, 1, 0, 1, 0, 1, 99},
			},
			want: 99,
		},
	}
	for _, tt := range tests {
		for _, f := range []func([]int) int{singleNumber5} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
