package offer

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 []int
		l2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				l1: []int{1, 2, 4},
				l2: []int{1, 3, 4},
			},
			want: []int{1, 1, 2, 3, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(NewListNode(tt.args.l1...), NewListNode(tt.args.l2...)); !reflect.DeepEqual(got, NewListNode(tt.want...)) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
