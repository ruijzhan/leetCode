package linkedlist

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_removeDuplicateNodes(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: New(1, 2, 3, 3, 2, 1),
			},
			want: New(1, 2, 3),
		},
		{
			args: args{
				head: New(1, 1, 1, 1, 1, 1, 1, 2),
			},
			want: New(1, 2),
		},
	}
	for _, tt := range tests {
		for _, f := range []func(*ListNode) *ListNode{removeDuplicateNodes, removeDuplicateNodes2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.head.DeepCopy()); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}
