package string

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_longestCommonPrefix(t *testing.T) {
	fs := []func([]string) string{longestCommonPrefix, longestCommonPrefix2}
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				strs: []string{"flower", "flow", "flight"},
			},
			want: "fl",
		},
		{
			args: args{
				strs: []string{"dog", "racecar", "car"},
			},
			want: "",
		},
		{
			args: args{
				strs: []string{"carrrrr", "car"},
			},
			want: "car",
		},
		{
			args: args{
				strs: []string{"car", "car"},
			},
			want: "car",
		},
		{
			args: args{
				strs: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		for _, f := range fs {

			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.strs); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v",
						utils.GetFunctionName(f), tt.args.strs, got, tt.want)
				}
			})
		}
	}
}
