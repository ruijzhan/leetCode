package linkedlist

import (
	"reflect"
	"testing"
)

func Test_makeLinkedList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				nums: []int{1, 2, 3},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
		{
			args: args{
				nums: []int{1},
			},
			want: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
		{
			args: args{
				nums: []int{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeLinkedList(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
