package guessnumber

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n    int
		pick int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				n:    10,
				pick: 6,
			},
			want: 6,
		},
		{
			name: "test",
			args: args{
				n:    1,
				pick: 1,
			},
			want: 1,
		},
		{
			name: "test",
			args: args{
				n:    2,
				pick: 1,
			},
			want: 1,
		},
		{
			name: "test",
			args: args{
				n:    2,
				pick: 2,
			},
			want: 2,
		},
		{
			name: "test",
			args: args{
				n:    3,
				pick: 2,
			},
			want: 2,
		},
		{
			name: "test",
			args: args{
				n:    3,
				pick: 3,
			},
			want: 3,
		},
		{
			name: "test",
			args: args{
				n:    3,
				pick: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pick = tt.args.pick
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
