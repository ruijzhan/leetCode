package tree

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	bst2Slice := func(root *TreeNode) []int {
		res := []int{}
		var helper func(*TreeNode)
		helper = func(node *TreeNode) {
			if node == nil {
				return
			}
			helper(node.Left)
			res = append(res, node.Val)
			helper(node.Right)
		}
		helper(root)
		return res
	}
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
				nums: []int{-10, -3, 0, 5, 9},
			},
			want: New([]int{0, -3, 9, -10, null, 5}),
		},
		{
			args: args{
				nums: []int{1, 3},
			},
			want: New([]int{3, 1}),
		},
		{
			args: args{
				nums: []int{},
			},
			want: New([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(bst2Slice(got), bst2Slice(tt.want)) {
				t.Errorf("sortedArrayToBST() = %v, want %v", bst2Slice(got), bst2Slice(tt.want))
			}
		})
	}
}
