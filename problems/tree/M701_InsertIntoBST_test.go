package tree

import (
	"reflect"
	"testing"
)

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				root: nil,
				val:  5,
			},
			want: &TreeNode{Val: 5},
		},
		{
			args: args{
				root: New([]int{4, 2, 7, 1, 3}),
				val:  5,
			},
			want: New([]int{4, 2, 7, 1, 3, 5}),
		},
		{
			args: args{
				root: New([]int{40, 20, 60, 10, 30, 50, 70}),
				val:  25,
			},
			want: New([]int{40, 20, 60, 10, 30, 50, 70, null, null, 25}),
		},
		{
			args: args{
				root: New([]int{4, 2, 7, 1, 3, null, null, null, null, null, null}),
				val:  5,
			},
			want: New([]int{4, 2, 7, 1, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
