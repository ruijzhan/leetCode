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
			name: "Test with abcabcbb",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Test with bbbbbbbbb",
			args: args{
				s: "bbbbbbbbb",
			},
			want: 1,
		},
		{
			name: "Test with pwwkew",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Test with empty string",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "Test with a",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "Test with au",
			args: args{
				s: "au",
			},
			want: 2,
		},
		{
			name: "Test with dvdf",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "Test with anviaj",
			args: args{
				s: "anviaj",
			},
			want: 5,
		},
		{
			name: "Test with tmmzuxt",
			args: args{
				s: "tmmzuxt",
			},
			want: 5,
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
