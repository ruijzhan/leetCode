package offer

import (
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_isSymmetric(t *testing.T) {
	type args struct {
		root []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: []int{1, 2, 2, 3, 4, 4, 3},
			},
			want: true,
		},
		{
			args: args{
				root: []int{1, 2, 2, null, 3, null, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric((*TreeNode)(tree.New(tt.args.root))); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
