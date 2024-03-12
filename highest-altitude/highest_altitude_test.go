package highestaltitude

import "testing"

func Test_largestAltitude(t *testing.T) {
	type args struct {
		gain []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				gain: []int{-5, 1, 5, 0, -7},
			},
			want: 1,
		},
		{
			name: "test",
			args: args{
				gain: []int{-4, -3, -2, -1, 4, 3, 2},
			},
			want: 0,
		},
		{
			name: "test",
			args: args{
				gain: []int{0},
			},
			want: 0,
		},
		{
			name: "test",
			args: args{
				gain: []int{1},
			},
			want: 1,
		},
		{
			name: "test",
			args: args{
				gain: []int{1, 2, 3, 4, 5},
			},
			want: 15,
		},
		{
			name: "test",
			args: args{
				gain: []int{-4, -3, -2, -1, 4, 3, 2},
			},
			want: 0,
		},
		{
			name: "test",
			args: args{
				gain: []int{-5, 1, 5, 0, -7},
			},
			want: 1,
		},
		{
			name: "test",
			args: args{
				gain: []int{-4, -3, -2, -1, 4, 3, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestAltitude(tt.args.gain); got != tt.want {
				t.Errorf("largestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
