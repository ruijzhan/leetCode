package linkedlist

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: newLinkedList(1),
				n:    1,
			},
			want: nil,
		},
		{
			args: args{
				head: newLinkedList(1, 2, 3, 4, 5),
				n:    2,
			},
			want: newLinkedList(1, 2, 3, 5),
		},
		{
			args: args{
				head: newLinkedList(1, 2),
				n:    1,
			},
			want: newLinkedList(1),
		},
		{
			args: args{
				head: newLinkedList(1, 2),
				n:    2,
			},
			want: newLinkedList(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
