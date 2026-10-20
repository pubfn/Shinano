package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_13001{
		ChapterList: []*protobuf.CHAPTERINFO{},
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
