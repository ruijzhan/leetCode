package array

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			args: args{
				digits: []int{0},
			},
			want: []int{1},
		},
		{
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
		{
			args: args{
				digits: []int{9, 9, 9},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			args: args{
				digits: []int{7, 9, 9},
			},
			want: []int{8, 0, 0},
		},
	}
	for _, tt := range tests {
		cp := make([]int, len(tt.args.digits))
		for _, f := range []func([]int) []int{plusOne, plusOne2} {
			copy(cp, tt.args.digits)
			t.Run(tt.name, func(t *testing.T) {
				if got := f(cp); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.digits, got, tt.want)
				}
			})
		}
	}
}
