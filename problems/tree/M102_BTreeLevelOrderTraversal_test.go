package tree

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_levelOrder(t *testing.T) {
	fs := []func(*TreeNode) [][]int{levelOrder, levelOrder2}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root: New([]int{3, 9, 20, null, null, 15, 7}),
			},
			want: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			args: args{
				root: New([]int{}),
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
