package array

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "empty",
			nums: nil,
			want: 0,
		},
		{
			name: "single",
			nums: []int{7},
			want: 1,
		},
		{
			name: "example1",
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			name: "example2",
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			name: "example3",
			nums: []int{1, 0, 1, 2},
			want: 3,
		},
		{
			name: "with_negatives",
			nums: []int{-2, -3, -1, 5, 4, 3},
			want: 3,
		},
		{
			name: "disconnected_sequences",
			nums: []int{10, 5, 12, 3, 55, 11, 4},
			want: 3,
		},
		{
			name: "all_duplicates",
			nums: []int{2, 2, 2},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Fatalf("longestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
