package linkedlist

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
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
				head: makeLinkedList(),
			},
			want: nil,
		},
		{
			args: args{
				head: makeLinkedList(0),
			},
			want: makeLinkedList(0),
		},
		{
			args: args{
				head: makeLinkedList(0, 1),
			},
			want: makeLinkedList(1, 0),
		},
		{
			args: args{
				head: makeLinkedList(0, 1, 2),
			},
			want: makeLinkedList(2, 1, 0),
		},
		{
			args: args{
				head: makeLinkedList(0, 1, 2, 3),
			},
			want: makeLinkedList(3, 2, 1, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
