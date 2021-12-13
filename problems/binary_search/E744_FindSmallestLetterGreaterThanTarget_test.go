package binary_search

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'a',
			},
			want: 'c',
		},
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'c',
			},
			want: 'f',
		},
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'd',
			},
			want: 'f',
		},
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'g',
			},
			want: 'j',
		},
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'j',
			},
			want: 'c',
		},
		{
			args: args{
				letters: []byte{'c', 'f', 'j'},
				target:  'k',
			},
			want: 'c',
		},
		{
			args: args{
				letters: []byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
				target:  'e',
			},
			want: 'n',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = '%c', want '%c'", got, tt.want)
			}
		})
	}
}
