package finddifferenceoftwoarrays

import (
	"slices"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[]int{1, 2, 3}, []int{2, 4, 6}}, [][]int{{3, 1}, {4, 6}}},
		{"2", args{[]int{1, 2, 3, 3}, []int{1, 1, 2, 2}}, [][]int{{3}, {}}},
		{"3", args{[]int{1, 2, 3}, []int{1, 2, 3}}, [][]int{{}, {}}},
		{"4", args{[]int{1, 2, 3}, []int{1, 1}}, [][]int{{2, 3}, {}}},
		{"5", args{[]int{1, 1, 3, 3}, []int{1, 2, 3, 3}}, [][]int{{}, {2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDifference(tt.args.nums1, tt.args.nums2)
			if len(got) != len(tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
			for i, arr := range got {
				if len(arr) != len(tt.want[i]) {
					t.Errorf("findDifference() = %v, want %v", got, tt.want)
				}
				for _, v := range arr {
					if !slices.Contains(tt.want[i], v) {
						t.Errorf("findDifference() = %v, want %v", got, tt.want)
					}
				}
			}
		})
	}
}
