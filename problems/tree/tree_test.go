package tree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{3, 9, 20, null, null, 15, 7},
			},
		},
		{
			args: args{
				nums: []int{5, 1, 4, null, null, 3, 6},
			},
		},
		{
			args: args{
				nums: []int{2, 1, 3},
			},
		},
		{
			args: args{
				nums: []int{1, 2, 2, 3, 4, 4, 3},
			},
		},
		{
			args: args{
				nums: []int{0, -10, 5, null, -3, null, 9},
			},
		},
		{
			args: args{
				nums: []int{0, -3, 9, -10, null, 5},
			},
		},
		{
			args: args{
				nums: []int{21, 7, 14, 1, 1, 2, 2, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := New(tt.args.nums)
			nums := root.slice()
			if len(tt.args.nums) != len(nums) {
				t.Errorf("NewTree() = %v, want %v", nums, tt.args.nums)
			}
			for i, n := range tt.args.nums {
				if n != nums[i] {
					t.Errorf("NewTree() = %v, want %v", nums, tt.args.nums)
				}
			}
		})
	}
}
