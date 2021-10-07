package string

import (
	"testing"

	"github.com/ruijzhan/leetCode/pkg/utils"
)

func Test_countSegments(t *testing.T) {
	fs := []func(string) int{countSegments, countSegments2}
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
				s: "Hello, my name is John",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		for _, f := range fs {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.args.s); got != tt.want {
					t.Errorf("%v(\"%v\") = %v, want %v",
						utils.GetFunctionName(f), tt.args.s, got, tt.want)
				}
			})
		}
	}
}
