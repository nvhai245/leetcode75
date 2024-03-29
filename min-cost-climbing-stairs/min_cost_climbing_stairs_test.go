package mincostclimbingstairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{10, 15, 20}}, 15},
		{"2", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
		{"3", args{[]int{0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
