package tree

import "testing"

func Test_isEvenOddTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				root: New([]int{1, 10, 4, 3, null, 7, 9, 12, 8, 6, null, null, 2}),
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{5, 4, 2, 3, 3, 7}),
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{5, 9, 1, 3, 5, 7}),
			},
			want: false,
		},
		{
			args: args{
				root: New([]int{1}),
			},
			want: true,
		},
		{
			args: args{
				root: New([]int{11, 8, 6, 1, 3, 9, 11, 30, 20, 18, 16, 12, 10, 4, 2, 17}),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEvenOddTree(tt.args.root); got != tt.want {
				t.Errorf("isEvenOddTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
