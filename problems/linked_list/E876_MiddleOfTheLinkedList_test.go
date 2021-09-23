package linkedlist

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
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
				head: makeLinkedList([]int{1}),
			},
			want: makeLinkedList([]int{1}),
		},
		{
			args: args{
				head: makeLinkedList([]int{1, 2, 3, 4, 5}),
			},
			want: makeLinkedList([]int{3, 4, 5}),
		},
		{
			args: args{
				head: makeLinkedList([]int{1, 2, 3, 4, 5, 6}),
			},
			want: makeLinkedList([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
