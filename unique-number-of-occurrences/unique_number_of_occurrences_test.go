package uniquenumberofoccurrences

import "testing"

func Test_uniqueOccurrences(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[]int{1, 2, 2, 1, 1, 3}}, true},
		{"2", args{[]int{1, 2}}, false},
		{"3", args{[]int{1, 2, 3}}, false},
		{"4", args{[]int{1, 2, 3, 4}}, false},
		{"5", args{[]int{1, 1, 1, 1, 1, 1}}, true},
		{"6", args{[]int{1, 2, 3, 2, 1}}, false},
		{"7", args{[]int{}}, true},
		{"8", args{[]int{1}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueOccurrences(tt.args.arr); got != tt.want {
				t.Errorf("uniqueOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}
