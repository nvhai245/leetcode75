package keysandrooms

type Queue struct {
	queue []int
}

func (q *Queue) Push(room int) {
	q.queue = append(q.queue, room)
}

func (q *Queue) Pop() int {
	if len(q.queue) == 0 {
		return -1
	}
	room := q.queue[0]
	q.queue = q.queue[1:]
	return room
}

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}
	visited := make(map[int]bool, len(rooms))
	keys := make(map[int]bool, len(rooms))
	q := new(Queue)
	q.Push(0)
	keys[0] = true

	for len(q.queue) > 0 {
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
