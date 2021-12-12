package string

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				s: "Hello",
			},
			want: "hello",
		},
		{
			args: args{
				s: "here",
			},
			want: "here",
		},
		{
			args: args{
				s: "LOVELY",
			},
			want: "lovely",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.s); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
