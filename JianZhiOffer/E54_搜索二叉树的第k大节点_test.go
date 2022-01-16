package offer

import (
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_kthLargest(t *testing.T) {
	type args struct {
		root []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				root: []int{3, 1, 4, null, 2},
				k:    1,
			},
			want: 4,
		},
		{
			args: args{
				root: []int{5, 3, 6, 2, 4, null, null, 1},
				k:    3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthLargest((*TreeNode)(tree.New(tt.args.root)), tt.args.k); got != tt.want {
				t.Errorf("kthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
