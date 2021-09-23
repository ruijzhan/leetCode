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
				head: makeLinkedList([]int{}),
			},
			want: nil,
		},
		{
			args: args{
				head: makeLinkedList([]int{0}),
			},
			want: makeLinkedList([]int{0}),
		},
		{
			args: args{
				head: makeLinkedList([]int{0, 1}),
			},
			want: makeLinkedList([]int{1, 0}),
		},
		{
			args: args{
				head: makeLinkedList([]int{0, 1, 2}),
			},
			want: makeLinkedList([]int{2, 1, 0}),
		},
		{
			args: args{
				head: makeLinkedList([]int{0, 1, 2, 3}),
			},
			want: makeLinkedList([]int{3, 2, 1, 0}),
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
