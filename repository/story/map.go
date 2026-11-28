package story

import (
	"fmt"
	"math/rand/v2"

	"Shinano/cfg"
)

func (s *Chapter) initMap(id uint32) {
	xOff := cfg.ChapterData[id].Grid[len(cfg.ChapterData[id].Grid)-1].X
	yOff := cfg.ChapterData[id].Grid[len(cfg.ChapterData[id].Grid)-1].Y

	width := (cfg.ChapterData[id].Grid[0].X - xOff) + 1
	height := (cfg.ChapterData[id].Grid[0].Y - yOff) + 1

	gMap := makeArr[cell](height, width)

	for i := len(cfg.ChapterData[id].Grid) - 1; i >= 0; i-- {
		if cfg.ChapterData[id].Grid[i].Active == false {
			x := cfg.ChapterData[id].Grid[i].X - xOff
			y := cfg.ChapterData[id].Grid[i].Y - yOff

			gMap[y][x] = wall
		}
	}

	s.Map = gMap
	s.Offset.x = xOff
	s.Offset.y = yOff
	s.Id = id
}

func (s *Chapter) initEntities() {
	for _, cell := range cfg.ChapterData[s.Id].Grid {
		if !cell.Active || cell.ItemType == 0 {
			continue
		}

		fmt.Println(cell.ItemType)

		switch cell.ItemType {
		case 1:
			s.Player = point{
				x: cell.X - s.Offset.x,
				y: cell.Y - s.Offset.y,
			}

		case 8:
			pos := point{
				x: cell.X - s.Offset.x,
				y: cell.Y - s.Offset.y,
			}

			for _, v := range cfg.ChapterData[s.Id].BossId {
				s.Boss = append(s.Boss, Enemy{
					Id:     uint32(v),
					Status: EnemyPending,
					Pos:    pos,
					Weight: 0,
				})
			}

		case 6:
			for _, v := range cfg.ChapterData[s.Id].Enemy {
				s.Enemies = append(s.Enemies, Enemy{
					Id:     uint32(v.Id),
					Status: EnemyPending,
					Pos:    point{},
					Weight: v.Weight,
				})
			}
			s.EnemySpawn = append(s.EnemySpawn, point{
				x: cell.X - s.Offset.x,
				y: cell.Y - s.Offset.y,
			})
		}
	}
	s.EnemyRefresh = append([]int(nil), cfg.ChapterData[s.Id].EnemyRefesh...)

	s.BossRefresh = cfg.ChapterData[s.Id].BossSpawnKillCount
}

func (s *Chapter) rollEnemy() {
	var totalWeight int = 0

	for _, enemy := range s.Enemies {
		if enemy.Status != EnemyPending {
			continue
		}

		totalWeight += enemy.Weight
	}

	num := rand.IntN(totalWeight)

	var buf int
	for i := range len(s.Enemies) {
		buf += s.Enemies[i].Weight

		if buf >= num {
			s.Enemies[i].Status = EnemyActive
			spawnPoint := rand.IntN(len(s.EnemySpawn))
			s.Enemies[i].Pos = s.EnemySpawn[spawnPoint]
			s.RemoveEnemySpawn(spawnPoint)
			return
		}
	}
}

func (s *Chapter) mapUpdate() {
	if len(s.EnemyRefresh) == 0 {
		return
	}

	for range s.EnemyRefresh[0] {
		s.rollEnemy()
	}

	s.EnemyRefresh = s.EnemyRefresh[1:]
}

func (s *Chapter) RemoveEnemySpawn(index int) {
	last := len(s.EnemySpawn) - 1

	s.EnemySpawn[index] = s.EnemySpawn[last]
	s.EnemySpawn = s.EnemySpawn[:last]
}
