package tree

import "testing"

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: New([]int{5, 3, 6, 2, 4, null, 7}),
				k:    9,
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{5, 3, 6, 2, 4, null, 7}),
				k:    28,
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{2, 1, 3}),
				k:    4,
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{2, 1, 3}),
				k:    1,
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{2, 1, 3}),
				k:    3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
