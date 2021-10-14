package binary_search

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_peakIndexInMountainArray(t *testing.T) {
	fs := []func([]int) int{peakIndexInMountainArray, peakIndexInMountainArray2}
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			args: args{
				arr: []int{0, 1, 2, 3, 4, 5, 6, 7, 6, 5, 4},
			},
			want: 7,
		},
		{
			args: args{
				arr: []int{8, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.arr); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.arr, got, tt.want)
				}
			})
		}
	}
}
