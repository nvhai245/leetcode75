package numberofprovinces

import "github.com/nvhai245/leetcode75/datastructures"

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) == 0 || len(isConnected[0]) == 0 {
		return 0
	}
	visited := make(map[int]bool, len(isConnected))
	st := new(datastructures.Stack[int])
	provinces := 0

	for city := range len(isConnected) {
		if _, isVisited := visited[city]; isVisited {
			continue
		}
		provinces++
		st.Push(city)
		dfs(st, isConnected, visited)
	}
	return provinces
}

func dfs(st *datastructures.Stack[int], isConnected [][]int, visited map[int]bool) {
	if len(st.Stack) == 0 {
		return
	}
	city := st.Pop()
	visited[city] = true
	for i, connected := range isConnected[city] {
		if i == city {
			continue
		}
		if connected == 1 {
			if _, isVisited := visited[i]; isVisited {
				continue
			}
			st.Push(i)
		}
	}
	dfs(st, isConnected, visited)
}
