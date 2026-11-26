package story

import (
	"slices"
)

type cell int

const (
	empty cell = iota
	wall
	hostile
)

func bfs(finish point) []point {
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	y := len(gameMap)
	x := len(gameMap[0])
	visited := makeArr[bool](y, x)
	parent := makeArr[point](y, x)
	queue := []point{}

	queue = append(queue, playerPos)
	visited[playerPos.y][playerPos.x] = true

	found := false

	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:]

		if current.x == finish.x &&
			current.y == finish.y {
			found = true
			break
		}

		for i := range 4 {
			nx := current.x + dx[i]
			ny := current.y + dy[i]

			if nx < 0 || ny < 0 ||
				nx >= x || ny >= y {
				continue
			}

			if gameMap[ny][nx] == wall {
				continue
			}

			if gameMap[ny][nx] == hostile &&
				!(nx == finish.x && ny == finish.y) {
				continue
			}

			if visited[ny][nx] {
				continue
			}

			visited[ny][nx] = true

			parent[ny][nx] = current

			queue = append(queue, point{x: nx, y: ny})
		}
	}

	if !found {
		return []point{}
	}

	path := []point{}
	cur := finish

	for !(cur.x == playerPos.x && cur.y == playerPos.y) {
		path = append(path, cur)

		cur = parent[cur.y][cur.x]
	}

	path = append(path, playerPos)

	slices.Reverse(path)

	return path[1:]
}

func makeArr[T any](height int, width int) [][]T {
	arr := make([][]T, height)

	for i := range arr {
		arr[i] = make([]T, width)
	}

	return arr
}
