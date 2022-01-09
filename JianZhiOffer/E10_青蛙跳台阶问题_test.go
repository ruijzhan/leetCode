package offer

import "testing"

func Test_numWays(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			args: args{
				n: 7,
			},
			want: 21,
		},
		{
			args: args{
				n: 0,
			},
			want: 1,
		},
		{
			args: args{
				n: 92,
			},
			want: 720754435,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWays(tt.args.n); got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
