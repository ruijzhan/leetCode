package linkedlist

import (
	"reflect"
	"testing"
)

func Test_mergeInBetween(t *testing.T) {
	type args struct {
		list1 *ListNode
		a     int
		b     int
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				list1: newLinkedList(0, 1, 2, 3, 4, 5),
				a:     3,
				b:     4,
				list2: newLinkedList(1000000, 1000001, 1000002),
			},
			want: newLinkedList(0, 1, 2, 1000000, 1000001, 1000002, 5),
		},
		{
			args: args{
				list1: newLinkedList(0, 1, 2, 3, 4, 5, 6),
				a:     2,
				b:     5,
				list2: newLinkedList(1000000, 1000001, 1000002, 1000003, 1000004),
			},
			want: newLinkedList(0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6),
		},
		{
			args: args{
				list1: newLinkedList(0, 1, 2),
				a:     1,
				b:     1,
				list2: newLinkedList(1000000, 1000001, 1000002, 1000003),
			},
			want: newLinkedList(0, 1000000, 1000001, 1000002, 1000003, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeInBetween(tt.args.list1, tt.args.a, tt.args.b, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeInBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
