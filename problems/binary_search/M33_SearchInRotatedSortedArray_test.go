package binary_search

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_searchInRotatedSortedArray(t *testing.T) {
	fs := []func([]int, int) int{searchInRotatedSortedArray, searchInRotatedSortedArray2}
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			args: args{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{5, 6, 7, 8, 9, 10, 0, 1, 2, 3},
				target: 2,
			},
			want: 8,
		},
		{
			args: args{
				nums:   []int{5, 1, 3},
				target: 5,
			},
			want: 0,
		},
		{
			args: args{
				nums:   []int{5, 6, 3},
				target: 3,
			},
			want: 2,
		},
		{
			args: args{
				nums:   []int{5, 1, 3},
				target: 3,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums, tt.args.target); got != tt.want {
					t.Errorf("%v(%v, %v) = %v, want %v", utils.GetFunctionName(f), tt.args.nums, tt.args.target, got, tt.want)
				}
			})
		}
	}
}
