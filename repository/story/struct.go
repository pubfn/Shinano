package story

type point struct {
	x int
	y int
}

type cell uint8

const (
	tile cell = iota
	wall
)

type EnemyStatus uint8

const (
	EnemyPending EnemyStatus = iota
	EnemyActive
	EnemyDefeated
)

type Enemy struct {
	Id     uint32
	Pos    point
	Status EnemyStatus
	Weight int
}

type Chapter struct {
	Id uint32

	BossRefresh int

	EnemyRefresh []int

	Map [][]cell

	Player point

	EnemySpawn []point

	Enemies []Enemy

	Boss []Enemy

	Offset point
}
