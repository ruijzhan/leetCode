package binary_search

import "testing"

func Test_findMin2(t *testing.T) {
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
				nums: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			args: args{
				nums: []int{2, 2, 2, 0, 1},
			},
			want: 0,
		},
		{
			args: args{
				nums: []int{4, 4, 4, 4, 4, 4, 4, 4, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin2(tt.args.nums); got != tt.want {
				t.Errorf("findMin2() = %v, want %v", got, tt.want)
			}
		})
	}
}
