package string

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_isAnagram(t *testing.T) {
	fs := []func(string, string) bool{isAnagram, isAnagram2}
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			args: args{
				s: "ratxxx",
				t: "car",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s, tt.args.t); got != tt.want {
					t.Errorf("%v(%v, %v) = %v, want %v",
						utils.GetFunctionName(f), tt.args.s, tt.args.t, got, tt.want)
				}
			})
		}
	}
}
