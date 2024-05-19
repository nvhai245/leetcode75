package nearestexit

import "testing"

func Test_nearestExit(t *testing.T) {
	type args struct {
		maze     [][]byte
		entrance []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				maze: [][]byte{
					{'+', '.', '.', '+'},
					{'.', '.', '.', '.'},
					{'.', '.', '.', '+'},
				},
				entrance: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "test 2",
			args: args{
				maze: [][]byte{
					{'+', '+', '.', '+'},
					{'+', '.', '.', '.'},
					{'+', '+', '.', '+'},
				},
				entrance: []int{1, 1},
			},
			want: 2,
		},
		{
			name: "test 3",
			args: args{
				maze: [][]byte{
					{'+', '.', '+', '+'},
					{'.', '+', '.', '+'},
					{'.', '.', '+', '+'},
				},
				entrance: []int{1, 2},
			},
			want: -1,
		},
		{
			name: "test 4",
			args: args{
				maze: [][]byte{
					{'.', '+'},
				},
				entrance: []int{0, 0},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestExit(tt.args.maze, tt.args.entrance); got != tt.want {
				t.Errorf("nearestExit() = %v, want %v", got, tt.want)
			}
		})
	}
}
