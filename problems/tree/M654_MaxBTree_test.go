package tree

import (
	"reflect"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				nums: []int{3, 2, 1, 6, 0, 5},
			},
			want: New([]int{6, 3, 5, null, 2, 0, null, null, 1}),
		},
		{
			args: args{
				nums: []int{3, 2, 1},
			},
			want: New([]int{3, null, 2, null, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMaximumBinaryTree(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMaximumBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
