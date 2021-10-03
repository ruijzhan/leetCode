package linkedlist

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				headA: newLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9),
				headB: newLinkedList(1, 3, 4, 56, 6),
			},
			want: nil,
		},
		{
			args: args{
				headA: newLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9),
				headB: newLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9),
			},
			want: nil,
		},
		{
			args: args{
				headA: newLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9),
				headB: newLinkedList(),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
