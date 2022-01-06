package offer

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_levelOrder2(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root: []int{3, 9, 20, null, null, 15, 7},
			},
			want: [][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrder2((*TreeNode)(tree.New(tt.args.root))); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder2() = %v, want %v", got, tt.want)
			}
		})
	}
}
