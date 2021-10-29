package linkedlist

import (
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
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
				head: New(4, 2, 1, 3),
			},
			want: New(1, 2, 3, 4),
		},
		{
			args: args{
				head: New(9, 8, 7, 6, 5, 4, 3, 2, 1, 0),
			},
			want: New(0, 1, 2, 3, 4, 5, 6, 7, 8, 9),
		},
		{
			args: args{
				head: New(1, 1, 1, 1, 2, 2, 2, 2),
			},
			want: New(1, 1, 1, 1, 2, 2, 2, 2),
		},
		{
			args: args{
				head: New(1),
			},
			want: New(1),
		},
		{
			args: args{
				head: New(),
			},
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
