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
				head: newRing(1, 3, 2, 0, -4),
			},
			want: true,
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
