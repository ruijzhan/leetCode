package linkedlist

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				head: newRing(0, 1, 2),
			},
			want: true,
		},
		{
			args: args{
				head: newRing(-1, 0),
			},
			want: false,
		},
		{
			args: args{
				head: newRing(-1, -21, 10, 17, 8, 4, 26, 5, 35, 33, -7, -16, 27, -12, 6, 29, -12, 5, 9, 20, 14, 14, 2, 13, -24, 21, 23, -21, 5),
			},
			want: false,
		},
		{
			args: args{
				head: newRing(1, 3, 2, 0, -4),
			},
			want: true,
		},
		{
			args: args{
				head: newRing(-1, 1, 1, 1, 1),
			},
			want: false,
		},
		{
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			args: args{
				head: &ListNode{},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
