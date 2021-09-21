package array

import "testing"

func Test_removeDuplicates(t *testing.T) {
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
				nums: []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9, 9},
			},
			want: 9,
		},
		{
			args: args{
				nums: []int{1, 1, 2, 3, 3, 4, 4, 4, 5, 6, 7, 7, 7, 7, 8, 8, 8, 9, 9, 9, 9, 9, 9, 10, 11, 12, 12, 12, 13, 13, 14, 15},
			},
			want: 15,
		},
		{
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
