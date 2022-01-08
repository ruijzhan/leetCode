package offer

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		head []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				head: []int{4, 5, 1, 9},
				val:  1,
			},
			want: []int{4, 5, 9},
		},
		{
			args: args{
				head: []int{4, 5, 1, 9},
				val:  5,
			},
			want: []int{4, 1, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(NewListNode(tt.args.head...), tt.args.val); !reflect.DeepEqual(got, NewListNode(tt.want...)) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
