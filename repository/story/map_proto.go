package story

import (
	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func (s *Chapter) getCellList() []*protobuf.CHAPTERCELLINFO_P13 {
	cellList := []*protobuf.CHAPTERCELLINFO_P13{}
	activeEnemy := []*Enemy{}

	for _, a := range s.Enemies {
		if a.Status == EnemyActive {
			activeEnemy = append(activeEnemy, &a)
		}
	}

	for y, row := range s.Map {
		for x, cell := range row {
			if cell == tile {
				buf := &protobuf.CHAPTERCELLINFO_P13{
					Pos: &protobuf.CHAPTERCELLPOS_P13{
						Row:    proto.Uint32(uint32(y + s.Offset.y)),
						Column: proto.Uint32(uint32(x + s.Offset.x)),
					},
					ItemType: proto.Uint32(0),
					ItemId:   proto.Uint32(0),
					ItemFlag: proto.Uint32(0),
					ItemData: proto.Uint32(0),
					ExtraId:  []uint32{},
				}

				for _, e := range activeEnemy {
					if e.Pos.x == x && e.Pos.y == y {
						buf.ItemType = proto.Uint32(6)
						buf.ItemId = &e.Id
					}
				}

				cellList = append(cellList, buf)
			}
		}
	}

	return cellList
}

func (c *Chapter) getBossList() []*protobuf.CELLFLAG {
	bossList := []*protobuf.CELLFLAG{}

	for _, b := range c.Boss {
		bossList = append(bossList, &protobuf.CELLFLAG{
			Pos: &protobuf.CHAPTERCELLPOS_P13{
				Row:    proto.Uint32(uint32(b.Pos.y + c.Offset.y)),
				Column: proto.Uint32(uint32(b.Pos.x + c.Offset.x)),
			},
			FlagList: []uint32{1},
		})
	}

	return bossList
}

func (s *Chapter) getMainGroupList(shipId []uint32, fleetId uint32) *protobuf.GROUPINCHAPTER_P13 {
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
			Row:    proto.Uint32(uint32(s.Player.y + s.Offset.y)),
			Column: proto.Uint32(uint32(s.Player.x + s.Offset.x)),
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
