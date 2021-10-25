package linkedlist

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_addTwoNumbers(t *testing.T) {
	fs := []func(*ListNode, *ListNode) *ListNode{addTwoNumbers, addTwoNumbers2}
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
				l1: New(2, 4, 3),
				l2: New(5, 6, 4),
			},
			want: New(7, 0, 8),
		},
		{
			args: args{
				l1: New(9, 9, 9, 9, 9, 9, 9),
				l2: New(9, 9, 9, 9),
			},
			want: New(8, 9, 9, 9, 0, 0, 0, 1),
		},
		{
			args: args{
				l1: nil,
				l2: nil,
			},
			want: nil,
		},
		{
			args: args{
				l1: New(2, 4, 3),
				l2: nil,
			},
			want: New(2, 4, 3),
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
