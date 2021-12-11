package binary_search

import (
	"reflect"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type args struct {
		arr []int
		k   int
		x   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				arr: []int{0, 1, 2, 2, 2, 3, 6, 8, 8, 9},
				k:   5,
				x:   9,
			},
			want: []int{3, 6, 8, 8, 9},
		},
		{
			args: args{
				arr: []int{0, 2, 2, 3, 4, 5, 6, 8, 9, 9},
				k:   4,
				x:   5,
			},
			want: []int{3, 4, 5, 6},
		},
		{
			args: args{
				arr: []int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8},
				k:   2,
				x:   2,
			},
			want: []int{1, 3},
		},
		{
			args: args{
				arr: []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8},
				k:   3,
				x:   5,
			},
			want: []int{3, 3, 4},
		},
		{
			args: args{
				arr: []int{1, 1, 1, 10, 10, 10},
				k:   1,
				x:   9,
			},
			want: []int{10},
		},
		{
			args: args{
				arr: []int{1},
				k:   1,
				x:   1,
			},
			want: []int{1},
		},
		{
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   -1,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosestElements(tt.args.arr, tt.args.k, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
