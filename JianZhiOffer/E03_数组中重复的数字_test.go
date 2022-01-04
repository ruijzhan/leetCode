package offer

import "testing"

func Test_findRepeatNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums: []int{2, 3, 1, 0, 2, 5, 3},
			},
			want: []int{2, 3},
		},
	}

	in := func(n int, ns []int) bool {
		for _, i := range ns {
			if i == n {
				return true
			}
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumber(tt.args.nums); !in(got, tt.want) {
				t.Errorf("findRepeatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
