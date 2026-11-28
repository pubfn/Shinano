package story

import (
	"database/sql"
	"time"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func GetChapterList(b *sql.DB) ([]*protobuf.CHAPTERINFO, error) {
	chapterList := []*protobuf.CHAPTERINFO{}

	rows, err := b.Query("select * from story")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id, progress, killboss, killenemy, takeboks, defeat, pass uint32

		err := rows.Scan(&id, &progress, &killboss, &killenemy, &takeboks, &defeat, &pass)
		if err != nil {
			return nil, err
		}

		chapterList = append(chapterList, &protobuf.CHAPTERINFO{
			Id:               &id,
			Progress:         &progress,
			KillBossCount:    &killboss,
			KillEnemyCount:   &killenemy,
			TakeBoxCount:     &takeboks,
			TodayDefeatCount: proto.Uint32(0),
			DefeatCount:      &defeat,
			PassCount:        &pass,
		})
	}

	return chapterList, nil
}

func GetChapterInfo(id int) {
}

func InitChapter(id uint32) *Chapter {
	var chap Chapter

	chap.initMap(id)
	chap.initEntities()
	chap.mapUpdate()

	return &chap
}

func (s *Chapter) StoryAct(actId uint32, y uint32, x uint32) []*protobuf.CHAPTERCELLPOS_P13 {
	movePath := []*protobuf.CHAPTERCELLPOS_P13{}

	switch actId {
	case 1:
		{
			path := s.bfs(point{
				x: int(x) - s.Offset.x,
				y: int(y) - s.Offset.y,
			})

			if len(path) == 0 {
				return movePath
			}

			for _, p := range path {
				buf := &protobuf.CHAPTERCELLPOS_P13{
					Row:    proto.Uint32(uint32(p.y + s.Offset.y)),
					Column: proto.Uint32(uint32(p.x + s.Offset.x)),
				}

				movePath = append(movePath, buf)
			}

			s.Player = path[len(path)-1]
		}
	}

	return movePath
}

func (c *Chapter) StartChapter(shipId []uint32, fleetId *uint32) *protobuf.CURRENTCHAPTERINFO {
	mainGroup := c.getMainGroupList(shipId, *fleetId)
	bossLi := c.getBossList()

	return &protobuf.CURRENTCHAPTERINFO{
		Id:                  &c.Id,
		CellList:            c.getCellList(),
		MainGroupList:       []*protobuf.GROUPINCHAPTER_P13{mainGroup},
		CellFlagList:        bossLi,
		ChapterStrategyList: []*protobuf.STRATEGYINFO_P13{},
		FleetDuties: []*protobuf.FLEETDUTYKEYVALUEPAIR{
			{
				Key:   proto.Uint32(1),
				Value: proto.Uint32(3),
			},
		},
		Time:                  proto.Uint32(uint32(time.Now().Unix() + 43200)),
		AiList:                []*protobuf.CHAPTERCELLINFO_P13{},
		EscortList:            []*protobuf.CHAPTERCELLINFO_P13{},
		Round:                 proto.Uint32(0),
		LoopFlag:              proto.Uint32(1),
		ExtraFlagList:         []uint32{},
		IsSubmarineAutoAttack: proto.Uint32(1),
		OperationBuff:         []uint32{},
		ModelActCount:         proto.Uint32(0),
		BuffList:              []uint32{},
		ChapterHp:             proto.Uint32(0),
		KillCount:             proto.Uint32(0),
		InitShipCount:         proto.Uint32(6),
		ContinuousKillCount:   proto.Uint32(0),
		BattleStatistics:      []*protobuf.STRATEGYINFO_P13{},
		MoveStepCount:         proto.Uint32(0),
		StartTime:             proto.Uint32(uint32(time.Now().Unix())),
		SubmarineGroupList:    []*protobuf.GROUPINCHAPTER_P13{},
		SupportGroupList:      []*protobuf.GROUPINCHAPTER_P13{},
	}
}
