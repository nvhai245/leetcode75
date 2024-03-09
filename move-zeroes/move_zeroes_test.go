package movezeroes

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "test 2",
			args: args{
				nums: []int{0},
			},
			want: []int{0},
		},
		{
			name: "test 3",
			args: args{
				nums: []int{0, 0},
			},
			want: []int{0, 0},
		},
		{
			name: "test 4",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: []int{0, 0, 0},
		},
		{
			name: "test 5",
			args: args{
				nums: []int{0, 0, 1},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "test 6",
			args: args{
				nums: []int{0, 1, 0},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "test 7",
			args: args{
				nums: []int{1, 0, 0},
			},
			want: []int{1, 0, 0},
		},
		{
			name: "test 8",
			args: args{
				nums: []int{1, 0, 0, 0},
			},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "test 9",
			args: args{
				nums: []int{1, 0, 0, 1, 0},
			},
			want: []int{1, 1, 0, 0, 0},
		},
		{
			name: "test 10",
			args: args{
				nums: []int{1, 0, 0, 1, 0, 1},
			},
			want: []int{1, 1, 1, 0, 0, 0},
		},
		{
			name: "test 11",
			args: args{
				nums: []int{0, 0, 0, 1, 0, 1, 0},
			},
			want: []int{1, 1, 0, 0, 0, 0, 0},
		},
		{
			name: "test 12",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 1},
			},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "test 13",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 0},
			},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			name: "test 14",
			args: args{
				nums: []int{0, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			name: "test 15",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			name: "test 16",
			args: args{
				nums: []int{0, 0, 0, 0, 0, 0, 0, 0},
			},
			want: []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
