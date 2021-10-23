package tree

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_increasingBST(t *testing.T) {
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
				root: New([]int{5, 3, 6, 2, 4, null, 8, 1, null, null, null, 7, 9}),
			},
			want: New([]int{1, null, 2, null, 3, null, 4, null, 5, null, 6, null, 7, null, 8, null, 9}),
		},
		{
			args: args{
				root: New([]int{5, 1, 7}),
			},
			want: New([]int{1, null, 5, null, 7}),
		},
	}
	for _, tt := range tests {
		for _, f := range []func(*TreeNode) *TreeNode{increasingBST, increasingBST2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
