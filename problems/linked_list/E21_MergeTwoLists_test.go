package linkedlist

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_mergeTwoLists(t *testing.T) {
	fs := []func(*ListNode, *ListNode) *ListNode{mergeTwoLists, mergeTwoLists2}
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			args: args{
				l1: New(1, 3, 5, 7, 9),
				l2: nil,
			},
			want: New(1, 3, 5, 7, 9),
		},
		{
			args: args{
				l1: New(1, 3, 5, 7, 9),
				l2: New(2, 4, 6, 8, 10),
			},
			want: New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		},
		{
			args: args{
				l1: New(1, 1, 2, 3, 3),
				l2: New(1, 1, 2, 2, 2),
			},
			want: New(1, 1, 1, 1, 2, 2, 2, 2, 3, 3),
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.l1.DeepCopy(), tt.args.l2.DeepCopy()); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%v, %v) = %v, want %v",
						utils.GetFunctionName(f), tt.args.l1, tt.args.l2, got, tt.want)
				}
			})
		}
	}
}
