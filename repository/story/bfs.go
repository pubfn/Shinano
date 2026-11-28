package story

import (
	"slices"
)

func (s *Chapter) bfs(finish point) []point {
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	y := len(s.Map)
	x := len(s.Map[0])
	visited := makeArr[bool](y, x)
	parent := makeArr[point](y, x)

	queue := []point{}

	queue = append(queue, s.Player)
	visited[s.Player.y][s.Player.x] = true

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

			if s.isBlocked(ny, nx, x, y, finish) {
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

	for !(cur.x == s.Player.x && cur.y == s.Player.y) {
		path = append(path, cur)

		cur = parent[cur.y][cur.x]
	}

	slices.Reverse(path)

	return path
}

func (s *Chapter) isBlocked(ny, nx, x, y int, finish point) bool {
	if nx < 0 || ny < 0 ||
		nx >= x || ny >= y {
		return true
	}

	if s.Map[ny][nx] == wall {
		return true
	}

	for _, e := range s.Enemies {
		if e.Pos.x == nx && e.Pos.y == ny &&
			!(nx == finish.x && ny == finish.y) {
			return true
		}
	}

	return false
}

func makeArr[T any](height int, width int) [][]T {
	arr := make([][]T, height)

	for i := range arr {
		arr[i] = make([]T, width)
	}

	return arr
}
