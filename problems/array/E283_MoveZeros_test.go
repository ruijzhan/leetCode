package array

import (
	"reflect"
	"testing"
)

func moveZerosTest(nums []int, f func([]int)) []int {
	f(nums)
	return nums
}

func Test_moveZerosTest(t *testing.T) {
	type args struct {
		nums []int
		f    func([]int)
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{0, 1, 0, 3, 12},
				f:    moveZeroes,
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			args: args{
				nums: []int{0, 1, 0, 3, 12},
				f:    moveZeroes2,
			},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveZerosTest(tt.args.nums, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveZerosTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
