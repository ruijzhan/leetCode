package linkedlist

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
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
				head: newRing(0, 1, 2),
			},
			want: newRing(0, 1, 2),
		},
		{
			args: args{
				head: newRing(-1, 0),
			},
			want: nil,
		},
		{
			args: args{
				head: newRing(1, 3, 2, 0, -4),
			},
			want: newRing(0, 2, 0, -4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.args.head)
			if got == nil || tt.want == nil {
				if got != tt.want {
					t.Errorf("detectCycle() error")
				}
				return
			}
			if !reflect.DeepEqual(got.Val, tt.want.Val) {
				t.Errorf("detectCycle() = %v, want %v", got.Val, tt.want.Val)
			}
		})
	}
}
