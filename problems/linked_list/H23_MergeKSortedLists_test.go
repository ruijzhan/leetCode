package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				lists: []*ListNode{
					New(1, 4, 5),
					New(1, 3, 4),
					New(2, 6),
				},
			},
			want: New(1, 1, 2, 3, 4, 4, 5, 6),
		},
		{
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			args: args{
				lists: []*ListNode{nil},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
