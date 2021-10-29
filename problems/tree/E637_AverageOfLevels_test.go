package tree

import (
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			args: args{
				root: New([]int{3, 9, 20, null, null, 15, 7}),
			},
			want: []float64{3, 14.5, 11},
		},
		{
			args: args{
				root: New([]int{3}),
			},
			want: []float64{3},
		},
		{
			args: args{
				root: nil,
			},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageOfLevels(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("averageOfLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
