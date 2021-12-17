package linkedlist

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
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
				head: New(1, 2, 3, 4, 5),
			},
			want: New(1, 3, 5, 2, 4),
		},
		{
			args: args{
				head: New(2, 1, 3, 5, 6, 4, 7),
			},
			want: New(2, 3, 6, 7, 1, 5, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
