package bt

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				s: "25525511135",
			},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			args: args{
				s: "0000",
			},
			want: []string{"0.0.0.0"},
		},
		{
			args: args{
				s: "1111",
			},
			want: []string{"1.1.1.1"},
		},
		{
			args: args{
				s: "010010",
			},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			args: args{
				s: "101023",
			},
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
