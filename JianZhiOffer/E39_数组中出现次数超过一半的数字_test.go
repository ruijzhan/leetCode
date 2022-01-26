package offer

import "testing"

func Test_majorityElement(t *testing.T) {
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
				nums: []int{1, 2, 3, 2, 2, 2, 5, 4, 2},
			},
			want: 2,
		},
	}
	fs := []func([]int) int{majorityElement, majorityElement2, majorityElement3}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); got != tt.want {
					t.Errorf("majorityElement() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}
