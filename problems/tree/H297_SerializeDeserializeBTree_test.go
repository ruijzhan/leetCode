package tree

import (
	"reflect"
	"testing"
)

func TestCodec_deserialize(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			args: args{
				data: "[1,2,3,null,null,4,5]",
			},
			want: New([]int{1, 2, 3, null, null, 4, 5}),
		},
		{
			args: args{
				data: "[]",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor()
			if got := c.deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Codec.deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				root: New([]int{1, 2, 3, null, null, 4, 5}),
			},
			want: "[1,2,3,null,null,4,5]",
		},
		{
			args: args{
				root: nil,
			},
			want: "[]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Constructor()
			if got := c.serialize(tt.args.root); got != tt.want {
				t.Errorf("Codec.serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
