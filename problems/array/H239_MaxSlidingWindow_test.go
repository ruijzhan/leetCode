package array

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_maxSlidingWindow(t *testing.T) {
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
				nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
				k:    3,
			},
			want: []int{3, 3, 5, 5, 6, 7},
		},
		{
			args: args{
				nums: []int{1, -1},
				k:    1,
			},
			want: []int{1, -1},
		},
		{
			args: args{
				nums: []int{4, -2},
				k:    2,
			},
			want: []int{4},
		},
		{
			args: args{
				nums: []int{9, 11},
				k:    2,
			},
			want: []int{11},
		},
		{
			args: args{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		for _, f := range []func([]int, int) []int{maxSlidingWindow} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
					if len(got) < 10000 {
						t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
					}
				}
			})
		}
	}
}
