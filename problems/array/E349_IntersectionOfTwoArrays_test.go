package array

import (
	"reflect"
	"sort"
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_intersection(t *testing.T) {
	fs := []func([]int, []int) []int{intersection, intersection2}

	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums1: []int{1, 2, 2, 1},
				nums2: []int{2, 2},
			},
			want: []int{2},
		},
		{
			args: args{
				nums1: []int{4, 9, 5},
				nums2: []int{9, 4, 9, 8, 4},
			},
			want: []int{9, 4},
		},
	}
	for _, tt := range tests {
		for _, f := range fs {

			t.Run(tt.name, func(t *testing.T) {
				got := f(tt.args.nums1, tt.args.nums2)
				sort.Ints(got)
				sort.Ints(tt.want)
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%v(%v, %v) = %v, want %v",
						utils.GetFunctionName(f), tt.args.nums1, tt.args.nums2, got, tt.want)
				}
			})
		}
	}
}
