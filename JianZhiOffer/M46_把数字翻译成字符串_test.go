package offer

import "testing"

func Test_translateNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				num: 12258,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := translateNum(tt.args.num); got != tt.want {
				t.Errorf("translateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
