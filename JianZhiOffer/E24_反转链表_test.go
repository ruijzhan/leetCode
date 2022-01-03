package offer

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				head: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{6, 5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(NewListNode(tt.args.head...)); !reflect.DeepEqual(got, NewListNode(tt.want...)) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
