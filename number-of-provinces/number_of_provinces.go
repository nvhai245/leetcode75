package numberofprovinces

import "github.com/nvhai245/leetcode75/datastructures"

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 || len(isConnected[0]) == 0 {
		return 0
	}
	visited := make(map[int]bool, len(isConnected))
	st := new(datastructures.Stack[int])
	st.Push(0)
	for len(st.Stack) > 0 {
		city := st.Pop()
		for i, connected := range isConnected[city] {

		}
	}
}
