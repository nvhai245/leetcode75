package keysandrooms

import "testing"

func Test_canVisitAllRooms(t *testing.T) {
	type args struct {
		rooms [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				rooms: [][]int{
					{1},
					{2},
					{3},
					{},
				},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				rooms: [][]int{
					{0, 1},
					{3, 2},
					{2},
					{4},
					{5},
					{6},
					{7},
					{7},
				},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				rooms: [][]int{
					{1, 3},
					{3, 0, 1},
					{2},
					{0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canVisitAllRooms(tt.args.rooms); got != tt.want {
				t.Errorf("canVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
