package array

import (
	"reflect"
	"testing"
)

func Test_testSetZeroes(t *testing.T) {
	testSetZeroes := func(m [][]int) [][]int {
		setZeroes(m)
		return m
	}
	type args struct {
		m [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				m: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			},
			want: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testSetZeroes(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testSetZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
