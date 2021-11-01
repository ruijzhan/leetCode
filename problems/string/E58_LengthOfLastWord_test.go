package string

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
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
				s: " ",
			},
			want: 0,
		},
		{
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			args: args{
				s: " c",
			},
			want: 1,
		},
		{
			args: args{
				s: "c",
			},
			want: 1,
		},
		{
			args: args{
				s: "aaa bbb c",
			},
			want: 1,
		},
		{
			args: args{
				s: "aaa as  bbb     ccccc",
			},
			want: 5,
		},
		{
			args: args{
				s: "   fly me   to   the moon  ",
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord(\"%v\") = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
