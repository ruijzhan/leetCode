package array

import (
	"reflect"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{3, 2, 3},
			},
			want: []int{3},
		},
		{
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			args: args{
				nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		for _, f := range []func([]int) []int{majorityElement, majorityElement2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.nums, got, tt.want)
				}
			})
		}
	}
}
