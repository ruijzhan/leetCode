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
				head: New(),
			},
			want: nil,
		},
		{
			args: args{
				head: New(0),
			},
			want: New(0),
		},
		{
			args: args{
				head: New(0, 1),
			},
			want: New(1, 0),
		},
		{
			args: args{
				head: New(0, 1, 2),
			},
			want: New(2, 1, 0),
		},
		{
			args: args{
				head: New(0, 1, 2, 3),
			},
			want: New(3, 2, 1, 0),
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

func Test_reverseList2(t *testing.T) {
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
				head: New(),
			},
			want: nil,
		},
		{
			args: args{
				head: New(0),
			},
			want: New(0),
		},
		{
			args: args{
				head: New(0, 1),
			},
			want: New(1, 0),
		},
		{
			args: args{
				head: New(0, 1, 2),
			},
			want: New(2, 1, 0),
		},
		{
			args: args{
				head: New(0, 1, 2, 3),
			},
			want: New(3, 2, 1, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
