package linkedlist

import (
	"reflect"
	"testing"
)

func Test_deleteMiddle(t *testing.T) {
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
				head: New(1, 3, 4, 7, 1, 2, 6),
			},
			want: New(1, 3, 4, 1, 2, 6),
		},
		{
			args: args{
				head: New(1, 3, 4, 7),
			},
			want: New(1, 3, 7),
		},
		{
			args: args{
				head: New(1, 3),
			},
			want: New(1),
		},
		{
			args: args{
				head: New(1),
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteMiddle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}
