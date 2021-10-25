package tree

import (
	"reflect"
	"testing"

	linkedlist "github.com/ruijzhan/leetCode/problems/linked_list"
)

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				head: linkedlist.New(-10, -3, 0, 5, 9),
			},
			want: New([]int{0, -3, 9, -10, null, 5}),
		},
		{
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got.inOrder(), tt.want.inOrder()) {
				t.Errorf("sortedListToBST() = %v, want %v", got.inOrder(), tt.want.inOrder())
			}
		})
	}
}
