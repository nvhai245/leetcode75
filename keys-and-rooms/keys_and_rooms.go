package keysandrooms

import (
	"github.com/nvhai245/leetcode75/datastructures"
)

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	visited := make(map[int]bool, len(rooms))
	keys := make(map[int]bool, len(rooms))
	q := new(datastructures.Queue[int])
	q.Push(0)
	keys[0] = true

	for len(q.Queue) > 0 {
		r := q.Pop()
		if _, ok := visited[r]; ok {
			continue
		} else if _, ok := keys[r]; ok {
			visited[r] = true
			for _, k := range rooms[r] {
				keys[k] = true
				if _, ok := visited[k]; ok {
					continue
				}
				q.Push(k)
			}
		} else {
			q.Push(r)
		}
	}

	return len(visited) == len(rooms)
}
