package binary_search

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums: []int{2, 3, 4, 5, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{2, 1},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{3, 4, 5, 6, 7, 1, 2},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
