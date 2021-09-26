package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
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
				l1: makeLinkedList(1, 3, 5, 7, 9),
				l2: nil,
			},
			want: makeLinkedList(1, 3, 5, 7, 9),
		},
		{
			args: args{
				l1: makeLinkedList(1, 3, 5, 7, 9),
				l2: makeLinkedList(2, 4, 6, 8, 10),
			},
			want: makeLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		},
		{
			args: args{
				l1: makeLinkedList(1, 1, 2, 3, 3),
				l2: makeLinkedList(1, 1, 2, 2, 2),
			},
			want: makeLinkedList(1, 1, 1, 1, 2, 2, 2, 2, 3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
