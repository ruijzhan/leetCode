package offer

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA []int
		headB []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			args: args{
				headA: []int{1, 2, 3, 4},
				headB: []int{4, 5, 6, 7, 8},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(NewListNode(tt.args.headA...), NewListNode(tt.args.headB...)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
