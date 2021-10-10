package binary_search

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_arrangeCoins(t *testing.T) {
	fs := []func(int) int{arrangeCoins, arrangeCoins2}
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 5,
			},
			want: 2,
		},
		{
			args: args{
				n: 8,
			},
			want: 3,
		},
		{
			args: args{
				n: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.n); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.n, got, tt.want)
				}
			})
		}
	}
}
