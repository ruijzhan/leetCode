package linkedlist

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: nil},
			want: nil,
		},
		{
			args: args{head: &ListNode{0, nil}},
			want: &ListNode{0, nil},
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2, 2, 3, 3),
			},
			want: newLinkedList(1, 2, 3),
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2),
			},
			want: newLinkedList(1, 2),
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2, 2, 2, 2, 3, 3, 4, 4),
			},
			want: newLinkedList(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicates2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: nil},
			want: nil,
		},
		{
			args: args{head: &ListNode{0, nil}},
			want: &ListNode{0, nil},
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2, 2, 3, 3),
			},
			want: newLinkedList(1, 2, 3),
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2),
			},
			want: newLinkedList(1, 2),
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2, 2, 2, 2, 3, 3, 4, 4),
			},
			want: newLinkedList(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicates3(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{head: nil},
			want: nil,
		},
		{
			args: args{head: &ListNode{0, nil}},
			want: &ListNode{0, nil},
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2, 2, 3, 3),
			},
			want: newLinkedList(1, 2, 3),
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2),
			},
			want: newLinkedList(1, 2),
		},
		{
			args: args{
				head: newLinkedList(1, 1, 2, 2, 2, 2, 3, 3, 4, 4),
			},
			want: newLinkedList(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates3() = %v, want %v", got, tt.want)
			}
		})
	}
}
