package offer

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_mirrorTree(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				root: []int{4, 2, 7, 1, 3, 6, 9},
			},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mirrorTree((*TreeNode)(tree.New(tt.args.root))); !reflect.DeepEqual((*tree.TreeNode)(got), tree.New(tt.want)) {
				t.Errorf("mirrorTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
