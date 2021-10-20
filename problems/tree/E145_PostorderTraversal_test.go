package tree

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: New([]int{1, null, 2, 3}),
			},
			want: []int{3, 2, 1},
		},
		{
			args: args{
				root: New([]int{}),
			},
			want: []int{},
		},
		{
			args: args{
				root: New([]int{1}),
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
