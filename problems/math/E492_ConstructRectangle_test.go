package math

import (
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				area: 4,
			},
			want: []int{2, 2},
		},
		{
			args: args{
				area: 6,
			},
			want: []int{3, 2},
		},
		{
			args: args{
				area: 8,
			},
			want: []int{4, 2},
		},
		{
			args: args{
				area: 10,
			},
			want: []int{5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructRectangle(tt.args.area); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
