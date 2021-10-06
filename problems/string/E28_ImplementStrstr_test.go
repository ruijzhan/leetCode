package string

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_strStr(t *testing.T) {
	fs := []func(string, string) int{strStr}
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				haystack: "hello",
				needle:   "ll",
			},
			want: 2,
		},
		{
			args: args{
				haystack: "aaaaa",
				needle:   "bba",
			},
			want: -1,
		},
		{
			args: args{
				haystack: "",
				needle:   "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.haystack, tt.args.needle); got != tt.want {
					t.Errorf("%v(\"%v\", \"%v\") = %v, want %v",
						utils.GetFunctionName(f), tt.args.haystack, tt.args.needle, got, tt.want)
				}
			})
		}
	}
}
