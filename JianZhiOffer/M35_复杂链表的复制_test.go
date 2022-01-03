package offer

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			args: args{
				head: [][]int{{3, null}, {3, 0}, {3, null}},
			},
			want: [][]int{{3, null}, {3, 0}, {3, null}},
		},
		{
			args: args{
				head: [][]int{},
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(NodeList(tt.args.head)); !reflect.DeepEqual(got, NodeList(tt.want)) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
