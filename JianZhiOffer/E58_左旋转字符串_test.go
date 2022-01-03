package offer

import "testing"

func Test_reverseLeftWords(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "cdefgab",
		},
		{
			args: args{
				s: "lrloseumgh",
				k: 6,
			},
			want: "umghlrlose",
		},
		{
			args: args{
				s: "",
				k: 0,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
