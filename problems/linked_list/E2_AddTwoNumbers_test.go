package linkedlist

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
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
				l1: newLinkedList(2, 4, 3),
				l2: newLinkedList(5, 6, 4),
			},
			want: newLinkedList(7, 0, 8),
		},
		{
			args: args{
				l1: newLinkedList(9, 9, 9, 9, 9, 9, 9),
				l2: newLinkedList(9, 9, 9, 9),
			},
			want: newLinkedList(8, 9, 9, 9, 0, 0, 0, 1),
		},
		{
			args: args{
				l1: newLinkedList(0),
				l2: newLinkedList(0),
			},
			want: newLinkedList(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
