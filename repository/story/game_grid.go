package story

import (
	"math/rand/v2"

	"Shinano/cfg"
	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

type point struct {
	y int
	x int
}

func getCellFlag(bossPos point) *protobuf.CELLFLAG {
	return &protobuf.CELLFLAG{
		Pos: &protobuf.CHAPTERCELLPOS_P13{
			Row:    proto.Uint32(uint32(bossPos.y)),
			Column: proto.Uint32(uint32(bossPos.x)),
		},
		FlagList: []uint32{1},
	}
}

func getMainGroupList(startPost point, shipId []uint32, fleetId uint32) *protobuf.GROUPINCHAPTER_P13 {
	shipList := []*protobuf.SHIPINCHAPTER_P13{}

	for _, id := range shipId {
		buf := &protobuf.SHIPINCHAPTER_P13{
			Id:     &id,
			HpRant: proto.Uint32(10000),
		}

		shipList = append(shipList, buf)
	}

	return &protobuf.GROUPINCHAPTER_P13{
		Id:       &fleetId,
		ShipList: shipList,
		Pos: &protobuf.CHAPTERCELLPOS_P13{
			Row:    proto.Uint32(uint32(startPost.y)),
			Column: proto.Uint32(uint32(startPost.x)),
		},
		StepCount:        proto.Uint32(0),
		BoxStrategyList:  []*protobuf.STRATEGYINFO_P13{},
		ShipStrategyList: []*protobuf.STRATEGYINFO_P13{},
		StrategyIds:      []uint32{},
		StartPos: &protobuf.CHAPTERCELLPOS_P13{
			Row:    proto.Uint32(0),
			Column: proto.Uint32(0),
		},
		Bullet:        proto.Uint32(5),
		CommanderList: []*protobuf.COMMANDERSINFO{},
		MoveStepDown:  proto.Uint32(0),
		KillCount:     proto.Uint32(0),
		FleetId:       &fleetId,
		VisionLv:      proto.Uint32(0),
	}
}

func fetchCellList(id uint32) ([]*protobuf.CHAPTERCELLINFO_P13, point, point) {
	cellList := []*protobuf.CHAPTERCELLINFO_P13{}
	shipSpawn := point{}
	flagPos := point{}

	for _, cell := range cfg.ChapterData[id].Grid {
		if !cell.Active {
			continue
		}

		buf := &protobuf.CHAPTERCELLINFO_P13{
			Pos: &protobuf.CHAPTERCELLPOS_P13{
				Row:    proto.Uint32(uint32(cell.Y)),
				Column: proto.Uint32(uint32(cell.X)),
			},
			ItemId:   proto.Uint32(0),
			ItemFlag: proto.Uint32(0),
			ItemData: proto.Uint32(0),
			ExtraId:  []uint32{},
		}

		switch cell.GridType {
		case 1:
			shipSpawn = point{
				x: cell.X,
				y: cell.Y,
			}
			buf.ItemType = proto.Uint32(0)

		case 8:
			flagPos = point{
				x: cell.X,
				y: cell.Y,
			}
			buf.ItemType = proto.Uint32(0)

		case 6:
			enemy := rollEnemy(cfg.ChapterData[id].Enemy)
			buf.ItemId = proto.Uint32(uint32(enemy))
			buf.ItemType = proto.Uint32(6)

		default:
			buf.ItemType = proto.Uint32(uint32(cell.GridType))
		}

		cellList = append(cellList, buf)
	}

	return cellList, shipSpawn, flagPos
}

func rollEnemy(enemies []cfg.Enemy) int {
	var totalWeight int = 0

	for _, enemy := range enemies {
		totalWeight += enemy.Weight
	}

	num := rand.IntN(totalWeight)

	var buf int
	for i := range len(enemies) {
		buf += enemies[i].Weight

		if buf >= num {
			return enemies[i].Id
		}
	}
	return -1
}
