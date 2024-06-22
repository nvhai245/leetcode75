package stringcompression

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want2 []byte
	}{
		{
			"1",
			args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}},
			6,
			[]byte{'a', '2', 'b', '2', 'c', '3'},
		},
		{
			"2",
			args{[]byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}},
			4,
			[]byte{'a', 'b', '1', '2'},
		},
		{
			"3",
			args{[]byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'}},
			6,
			[]byte{'a', '3', 'b', '2', 'a', '2'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chars := tt.args.chars
			if got := compress(chars); got != tt.want {
				t.Errorf(
					"compress() = %v, got %s, want %v",
					got,
					chars,
					tt.want,
				)
			}
		})
	}
}

func Test_repeatNumToBytes(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"1", args{123}, []byte{'1', '2', '3'}},
		{"2", args{12345}, []byte{'1', '2', '3', '4', '5'}},
		{"3", args{123456789}, []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatNumToBytes(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repeatNumToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
