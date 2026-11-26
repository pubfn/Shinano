package story

import "Shinano/cfg"

var playerPos point
var (
	gameMap          [][]cell
	xOffset, yOffset int
)

func initMap(id uint32) {
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

		if cfg.ChapterData[id].Grid[i].GridType == 6 ||
			cfg.ChapterData[id].Grid[i].GridType == 8 {
			x := cfg.ChapterData[id].Grid[i].X - xOff
			y := cfg.ChapterData[id].Grid[i].Y - yOff

			gMap[y][x] = hostile
		}
	}

	gameMap = gMap
	xOffset = xOff
	yOffset = yOff
}
