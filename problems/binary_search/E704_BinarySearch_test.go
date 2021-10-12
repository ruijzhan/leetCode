package binary_search

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_search(t *testing.T) {
	fs := []func([]int, int) int{search, search2, search3}
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
		{
			args: args{
				nums:   []int{-1, 0, 5},
				target: -1,
			},
			want: 0,
		},
		{
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 13,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums, tt.args.target); got != tt.want {
					t.Errorf("%v() = %v, want %v", utils.GetFunctionName(f), got, tt.want)
				}
			})
		}
	}
}

func Test_search4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8},
				target: 4,
			},
			want: 3,
		},
		{
			args: args{
				nums:   []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8},
				target: 5,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search5(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8},
				target: 4,
			},
			want: 8,
		},
		{
			args: args{
				nums:   []int{1, 2, 3, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8},
				target: 5,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search5(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search6(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{3, 4, 6, 7, 10},
				target: 5,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search6(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_search7(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				nums:   []int{3, 4, 6, 7, 10},
				target: 5,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search7(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search7() = %v, want %v", got, tt.want)
			}
		})
	}
}
