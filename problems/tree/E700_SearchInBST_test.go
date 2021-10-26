package tree

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_searchBST(t *testing.T) {
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
				root: New([]int{4, 2, 7, 1, 3}),
				val:  2,
			},
			want: New([]int{2, 1, 3}),
		},
		{
			args: args{
				root: New([]int{}),
				val:  2,
			},
			want: New([]int{}),
		},
	}
	for _, tt := range tests {
		for _, f := range []func(*TreeNode, int) *TreeNode{searchBST, searchBST2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
