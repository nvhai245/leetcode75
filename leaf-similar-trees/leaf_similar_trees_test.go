package leafsimilartrees

import (
	"reflect"
	"testing"
)

func Test_leafSimilar(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				&TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}},
				&TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 5}},
			},
			false,
		},
		{
			"2",
			args{
				&TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6}},
					Right: &TreeNode{Val: 1},
				},
				&TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}}},
					Right: &TreeNode{Val: 1},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSimilar(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leafSequence(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{&TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
			[]int{5, 1},
		},
		{
			"2",
			args{
				&TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6}},
					Right: &TreeNode{Val: 1},
				},
			},
			[]int{6, 1},
		},
		{
			"3",
			args{
				&TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 5, Left: &TreeNode{Val: 6, Left: &TreeNode{Val: 7}}},
					Right: &TreeNode{Val: 1},
				},
			},
			[]int{7, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leafSequence(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("leafSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
