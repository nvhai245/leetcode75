package mergestringsalternatively

import "testing"

func Test_mergeAlternatively(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should success",
			args: args{
				word1: "abc",
				word2: "def",
			},
			want: "adbecf",
		},
		{
			name: "should success",
			args: args{
				word1: "abcd",
				word2: "efg",
			},
			want: "aebfcgd",
		},
		{
			name: "should success",
			args: args{
				word1: "abc",
				word2: "defg",
			},
			want: "adbecfg",
		},
		{
			name: "should success",
			args: args{
				word1: "a",
				word2: "bcde",
			},
			want: "abcde",
		},
		{
			name: "should success",
			args: args{
				word1: "abcdef",
				word2: "g",
			},
			want: "agbcdef",
		},
		{
			name: "should success",
			args: args{
				word1: "",
				word2: "abcdef",
			},
			want: "abcdef",
		},
		{
			name: "should success",
			args: args{
				word1: "abcdefgh",
				word2: "",
			},
			want: "abcdefgh",
		},
		{
			name: "should success",
			args: args{
				word1: "",
				word2: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAlternatively(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("mergeAlternatively() = %v, want %v", got, tt.want)
			}
		})
	}
}
