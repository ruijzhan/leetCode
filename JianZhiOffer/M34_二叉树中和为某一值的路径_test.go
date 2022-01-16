package offer

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/problems/tree"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				root:   []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1},
				target: 22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			args: args{
				root:   []int{1, 2, 3},
				target: 4,
			},
			want: [][]int{{1, 3}},
		},
		{
			args: args{
				root:   []int{1, 2},
				target: 3,
			},
			want: [][]int{{1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum((*TreeNode)(tree.New(tt.args.root)), tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
