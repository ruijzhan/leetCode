package array

import (
	"reflect"
	"testing"
)

func Test_testReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			args: args{
				s: []byte{'1', '2', '3'},
			},
			want: []byte{'3', '2', '1'},
		},
		{
			args: args{
				s: []byte{'1', '2', '3', '4'},
			},
			want: []byte{'4', '3', '2', '1'},
		},
		{
			args: args{
				s: []byte{'1'},
			},
			want: []byte{'1'},
		},
		{
			args: args{
				s: []byte{},
			},
			want: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := testReverseString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
