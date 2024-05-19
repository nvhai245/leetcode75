package nearestexit

import "fmt"

type Position struct {
	coordinate []int
	step       int
}

type Queue struct {
	queue []*Position
}

func (q *Queue) Push(item *Position) {
	q.queue = append(q.queue, item)
}

func (q *Queue) Pop() *Position {
	if len(q.queue) == 0 {
		return nil
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func isExit(maze [][]byte, position *Position) bool {
	coordinate := position.coordinate
	if position.step != 0 &&
		(coordinate[0] == 0 ||
			coordinate[0] == len(maze)-1 ||
			coordinate[1] == 0 ||
			coordinate[1] == len(maze[0])-1) {
		return true
	}

	return false
}

func getNeighbors(maze [][]byte, position *Position) []*Position {
	var neighbors []*Position
	// Up
	if row := position.coordinate[0] - 1; row >= 0 && maze[row][position.coordinate[1]] == '.' {
		neighbors = append(neighbors, &Position{
			coordinate: []int{row, position.coordinate[1]},
			step:       position.step + 1,
		})
	}
	// Down
	if row := position.coordinate[0] + 1; row < len(maze) &&
		maze[row][position.coordinate[1]] == '.' {
		neighbors = append(neighbors, &Position{
			coordinate: []int{row, position.coordinate[1]},
			step:       position.step + 1,
		})
	}
	// Left
	if col := position.coordinate[1] - 1; col >= 0 && maze[position.coordinate[0]][col] == '.' {
		neighbors = append(neighbors, &Position{
			coordinate: []int{position.coordinate[0], col},
			step:       position.step + 1,
		})
	}
	// Right
	if col := position.coordinate[1] + 1; col < len(maze[0]) &&
		maze[position.coordinate[0]][col] == '.' {
		neighbors = append(neighbors, &Position{
			coordinate: []int{position.coordinate[0], col},
			step:       position.step + 1,
		})
	}
	return neighbors
}

func nearestExit(maze [][]byte, entrance []int) int {
	entrancePos := &Position{
		coordinate: entrance,
		step:       0,
	}
	queue := &Queue{
		queue: []*Position{
			entrancePos,
		},
	}
	visited := make(map[string]*Position)
	visited[fmt.Sprintf("%d,%d", entrance[0], entrance[1])] = entrancePos
	for {
		curr := queue.Pop()
		if curr == nil {
			break
		}
		if isExit(maze, curr) {
			return curr.step
		}

		for _, n := range getNeighbors(maze, curr) {
			if _, ok := visited[fmt.Sprintf("%d,%d", n.coordinate[0], n.coordinate[1])]; !ok {
				queue.Push(n)
				visited[fmt.Sprintf("%d,%d", n.coordinate[0], n.coordinate[1])] = n
			}
		}

	}
	return -1
}
