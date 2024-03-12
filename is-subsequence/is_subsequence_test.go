package issubsequence

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abc", "ahbgdc"}, true},
		{"2", args{"axc", "ahbgdc"}, false},
		{"3", args{"ace", "abcde"}, true},
		{"4", args{"aec", "abcde"}, false},
		{"5", args{"aaa", "a"}, false},
		{"6", args{"", "a"}, true},
		{"7", args{"", ""}, true},
		{"8", args{"a", "a"}, true},
		{"9", args{"aaa", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
