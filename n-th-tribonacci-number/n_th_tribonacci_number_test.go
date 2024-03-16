package nthtribonaccinumber

import "testing"

func Test_tribonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{0}, 0},
		{"2", args{1}, 1},
		{"3", args{2}, 1},
		{"4", args{3}, 2},
		{"5", args{4}, 4},
		{"6", args{5}, 7},
		{"7", args{6}, 13},
		{"8", args{7}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tribonacci(tt.args.n); got != tt.want {
				t.Errorf("tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
