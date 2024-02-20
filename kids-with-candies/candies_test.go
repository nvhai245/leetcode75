package kidswithcandies

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "should success",
			args: args{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: []bool{true, true, true, false, true},
		},
		{
			name: "should success",
			args: args{
				candies:      []int{1, 1, 2, 2, 5, 5},
				extraCandies: 3,
			},
			want: []bool{false, false, true, true, true, true},
		},
		{
			name: "should success",
			args: args{
				candies:      []int{1, 1, 1, 1, 1},
				extraCandies: 1,
			},
			want: []bool{true, true, true, true, true},
		},
		{
			name: "should success",
			args: args{
				candies:      []int{1, 1, 1, 1, 1},
				extraCandies: 0,
			},
			want: []bool{true, true, true, true, true},
		},
		{
			name: "should success",
			args: args{
				candies:      []int{1, 1, 1, 1, 1},
				extraCandies: 2,
			},
			want: []bool{true, true, true, true, true},
		},
		{
			name: "should success",
			args: args{
				candies:      []int{12, 1, 12},
				extraCandies: 10,
			},
			want: []bool{true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(
				got,
				tt.want,
			) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
