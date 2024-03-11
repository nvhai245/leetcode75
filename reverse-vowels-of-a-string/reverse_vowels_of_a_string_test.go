package reversevowelsofastring

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"hello"}, "holle"},
		{"2", args{"leetcode"}, "leotcede"},
		{"3", args{"aA"}, "Aa"},
		{"4", args{"aA"}, "Aa"},
		{"5", args{"aA"}, "Aa"},
		{"6", args{"look! an apple!"}, "leak! an opplo!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}
