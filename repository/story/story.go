package story

import (
	"database/sql"
	"fmt"
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

func StoryAct(actId uint32, y uint32, x uint32) []*protobuf.CHAPTERCELLPOS_P13 {
	movePath := []*protobuf.CHAPTERCELLPOS_P13{}
	fmt.Println(actId)
	switch actId {
	case 1:
		{
			path := bfs(point{
				x: int(x) - xOffset,

				y: int(y) - yOffset,
			})
			if len(path) == 0 {
				fmt.Println("len is 0")
				return movePath
			}

			for _, p := range path {
				buf := &protobuf.CHAPTERCELLPOS_P13{
					Row:    proto.Uint32(uint32(p.y + yOffset)),
					Column: proto.Uint32(uint32(p.x + xOffset)),
				}
				fmt.Println("x:", p.x+xOffset)
				fmt.Println("y:", p.y+yOffset)
				fmt.Println()
				movePath = append(movePath, buf)
			}

			playerPos = path[len(path)-1]
		}
	}

	return movePath
}

func StartChapter(id uint32, shipId []uint32, fleetId *uint32) *protobuf.CURRENTCHAPTERINFO {
	cellList, startPos, bossPos := fetchCellList(id)

	initMap(id)

	playerPos = point{x: startPos.x - xOffset, y: startPos.y - yOffset}

	mainGroup := getMainGroupList(startPos, shipId, *fleetId)

	cellFlag := getCellFlag(bossPos)

	chapterInfo := &protobuf.CURRENTCHAPTERINFO{
		Id:                  &id,
		CellList:            cellList,
		MainGroupList:       []*protobuf.GROUPINCHAPTER_P13{mainGroup},
		CellFlagList:        []*protobuf.CELLFLAG{cellFlag},
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

	return chapterInfo
}
