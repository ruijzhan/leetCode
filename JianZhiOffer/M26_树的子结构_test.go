package offer

import (
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				A: []int{1, 2, 3},
				B: []int{3, 1},
			},
			want: false,
		},
		{
			args: args{
				A: []int{3, 4, 5, 1, 2},
				B: []int{4, 1},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubStructure((*TreeNode)(tree.New(tt.args.A)), (*TreeNode)(tree.New(tt.args.B))); got != tt.want {
				t.Errorf("isSubStructure() = %v, want %v", got, tt.want)
			}
		})
	}
}
