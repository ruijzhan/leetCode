package array

import (
	"reflect"
	"testing"
)

func Test_singleNumber6(t *testing.T) {
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
				nums: []int{1, 2, 3, 1, 4, 2},
			},
			want: []int{3, 4},
		},
		{
			args: args{
				nums: []int{3, 4},
			},
			want: []int{3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber6(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber6() = %v, want %v", got, tt.want)
			}
		})
	}
}
