package string

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			args: args{
				s: "race a car",
			},
			want: false,
		},
		{
			args: args{
				s: "aa",
			},
			want: true,
		},
		{
			args: args{
				s: "ab",
			},
			want: false,
		},
		{
			args: args{
				s: "0P",
			},
			want: false,
		},
		{
			args: args{
				s: "ab2a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
