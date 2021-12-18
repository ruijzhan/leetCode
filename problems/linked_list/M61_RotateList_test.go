package linkedlist

import (
	"math"
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				head: New(1, 2),
				k:    2,
			},
			want: New(1, 2),
		},
		{
			args: args{
				head: New(),
				k:    1,
			},
			want: New(),
		},
		{
			args: args{
				head: New(),
				k:    0,
			},
			want: New(),
		},
		{
			args: args{
				head: New(1, 2, 3, 4, 5),
				k:    2,
			},
			want: New(4, 5, 1, 2, 3),
		},
		{
			args: args{
				head: New(0, 1, 2),
				k:    4,
			},
			want: New(2, 0, 1),
		},
		{
			args: args{
				head: New(1, 2, 3, 4, 5),
				k:    math.MaxInt64,
			},
			want: New(4, 5, 1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
