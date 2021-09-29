package stack

import (
	"reflect"
	"testing"
)

func Test_genericStack_Empty(t *testing.T) {
	type fields struct {
		array []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			fields: fields{
				array: []interface{}{},
			},
			want: true,
		},
		{
			fields: fields{
				array: []interface{}{struct{}{}},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &genericStack{
				array: tt.fields.array,
			}
			if got := s.Empty(); got != tt.want {
				t.Errorf("genericStack.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_genericStack_Push(t *testing.T) {
	type fields struct {
		array []interface{}
	}
	type args struct {
		e interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantLen int
	}{
		{
			name: "",
			fields: fields{
				array: []interface{}{},
			},
			args: args{
				e: struct{}{},
			},
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &genericStack{
				array: tt.fields.array,
			}
			s.Push(tt.args.e)
			if s.Len() != tt.wantLen {
				t.Errorf("genericStack.Empty() = %v, want %v", s.Len(), tt.wantLen)
			}
		})
	}
}

func Test_genericStack_Pop(t *testing.T) {
	type fields struct {
		array []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
		wantLen int
	}{
		{
			fields: fields{
				array: []interface{}{struct{}{}},
			},
			want:    struct{}{},
			wantErr: false,
			wantLen: 0,
		},
		{
			fields: fields{
				array: []interface{}{},
			},
			want:    nil,
			wantErr: true,
			wantLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &genericStack{
				array: tt.fields.array,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("genericStack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genericStack.Pop() = %v, want %v", got, tt.want)
			}
			if s.Len() != tt.wantLen {
				t.Errorf("genericStack.Len() = %v, want %v", s.Len(), tt.wantLen)
			}
		})
	}
}
