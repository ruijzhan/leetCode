package linkedlist

import (
	"reflect"
	"testing"
)

func Test_newLinkedList(t *testing.T) {
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
						Val: 3,
					},
				},
			},
		},
		{
			args: args{
				nums: []int{1},
			},
			want: &ListNode{
				Val: 1,
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
			if got := newLinkedList(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeLinkedList() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := newLinkedList2(tt.args.nums...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeLinkedList2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDoublyLinkedList(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newDoublyLinkedList(tt.args.nums...)
			for l.Next != nil {
				if l.Next.Prev != l {
					t.Errorf("newDoublyLinkedList() = %v, want %v", l.Next.Prev, l)
				}
				l = l.Next
			}
		})
	}
}
