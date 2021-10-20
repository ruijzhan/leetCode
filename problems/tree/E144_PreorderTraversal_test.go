package tree

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
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
			want: []int{1, 2, 3},
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
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
