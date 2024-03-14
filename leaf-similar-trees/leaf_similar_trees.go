package leafsimilartrees

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	s1, s2 := leafSequence(root1), leafSequence(root2)
	if len(s1) != len(s2) {
		return false
	}
	for i, leaf := range s1 {
		if leaf != s2[i] {
			return false
		}
	}
	return true
}

func leafSequence(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	return slices.Concat(leafSequence(root.Left), leafSequence(root.Right))
}
