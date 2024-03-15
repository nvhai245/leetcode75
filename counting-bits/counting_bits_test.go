package countingbits

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{2}, []int{0, 1, 1}},
		{"2", args{5}, []int{0, 1, 1, 2, 1, 2}},
		{"3", args{10}, []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2}},
		{"4", args{0}, []int{0}},
		{"5", args{1}, []int{0, 1}},
		{"6", args{3}, []int{0, 1, 1, 2}},
		{"7", args{6}, []int{0, 1, 1, 2, 1, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBits(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
