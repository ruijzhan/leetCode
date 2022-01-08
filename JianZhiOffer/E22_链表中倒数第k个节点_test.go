package offer

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				head: []int{1, 2, 3, 4, 5},
				k:    2,
			},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(NewListNode(tt.args.head...), tt.args.k); !reflect.DeepEqual(got, NewListNode(tt.want...)) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
