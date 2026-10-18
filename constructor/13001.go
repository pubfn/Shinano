package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_13001{
		ChapterList: []*protobuf.CHAPTERINFO{
			{
				Id:               proto.Uint32(101),
				Progress:         proto.Uint32(100),
				KillBossCount:    proto.Uint32(1),
				KillEnemyCount:   proto.Uint32(1),
				TakeBoxCount:     proto.Uint32(1),
				DefeatCount:      proto.Uint32(1),
				TodayDefeatCount: proto.Uint32(0),
				PassCount:        proto.Uint32(0),
			},
		},
		ReactChapter: &protobuf.REACTCHAPTER_INFO{
			Count:           proto.Uint32(0),
			ActiveTimestamp: proto.Uint32(0),
			ActiveId:        proto.Uint32(0),
			DailyCount:      proto.Uint32(0),
		},
		FleetList: []*protobuf.FLEET_INFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13001
}
