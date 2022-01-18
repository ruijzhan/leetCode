package offer

import "testing"

func Test_minNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				nums: []int{10, 2},
			},
			want: "102",
		},
		{
			args: args{
				nums: []int{3, 30, 34, 5, 9},
			},
			want: "3033459",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumber(tt.args.nums); got != tt.want {
				t.Errorf("minNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
