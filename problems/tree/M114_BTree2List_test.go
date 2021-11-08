package tree

import (
	"reflect"
	"testing"
)

func Test_flattenTest(t *testing.T) {
	flattenTest := func(root *TreeNode) *TreeNode {
		flatten(root)
		return root
	}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: New([]int{1, 2, 3}),
			},
			want: New([]int{1, null, 2, null, 3}),
		},
		{
			args: args{
				root: New([]int{1, 2, 5, 3, 4, null, 6}),
			},
			want: New([]int{1, null, 2, null, 3, null, 4, null, 5, null, 6}),
		},
		{
			args: args{
				root: New([]int{0}),
			},
			want: New([]int{0}),
		},
		{
			args: args{
				root: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flattenTest(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flattenTest() = %v, want %v", got.levelOrder(), tt.want.levelOrder())
			}
		})
	}
}
