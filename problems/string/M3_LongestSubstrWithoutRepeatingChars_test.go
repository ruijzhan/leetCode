package string

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			args: args{
				s: "bbbbbbbbb",
			},
			want: 1,
		},
		{
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		for _, f := range []func(string) int{lengthOfLongestSubstring, lengthOfLongestSubstring2} {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s); got != tt.want {
					t.Errorf("%v(%v) = %v, want %v", utils.GetFunctionName(f), tt.args.s, got, tt.want)
				}
			})
		}
	}
}
