package string

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				words: []string{"Hello", "Alaska", "Dad", "Peace"},
			},
			want: []string{"Alaska", "Dad"},
		},
		{
			args: args{
				words: []string{"omk"},
			},
			want: []string{},
		},
		{
			args: args{
				words: []string{"adsdf", "sfd"},
			},
			want: []string{"adsdf", "sfd"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
