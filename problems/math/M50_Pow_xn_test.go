package math

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				x: 2.0,
				n: 10,
			},
			want: 1024.0,
		},
		{
			args: args{
				x: 2.1,
				n: 3,
			},
			want: 9.261000000000001,
		},
		{
			args: args{
				x: 2,
				n: -2,
			},
			want: 0.250,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
