package binary_search

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want map[int]bool
	}{
		{
			args: args{
				nums: []int{1},
			},
			want: map[int]bool{0: true},
		},
		{
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: map[int]bool{2: true},
		},
		{
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: map[int]bool{1: true, 5: true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); !tt.want[got] {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
