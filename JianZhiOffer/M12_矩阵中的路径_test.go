package offer

import "testing"

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				board: [][]byte{
					{'A', 'B', 'C', 'E'},
					{'S', 'F', 'C', 'S'},
					{'A', 'D', 'E', 'E'},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				word: "abcd",
			},
			want: false,
		},
		{
			args: args{
				board: [][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
