package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/story"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_13001(msg *[]byte) ([]byte, int) {
	chapterList, err := story.GetChapterList(b.DB)
	if err != nil {
		log.Fatal(err)
	}

	sc := &protobuf.SC_13001{
		ChapterList: chapterList,
		ReactChapter: &protobuf.REACTCHAPTER_INFO{
			Count:           proto.Uint32(0),
			ActiveTimestamp: proto.Uint32(0),
			ActiveId:        proto.Uint32(0),
			DailyCount:      proto.Uint32(0),
		},
		FleetList:    []*protobuf.FLEET_INFO{},
		Oil:          proto.Uint32(0),
		TimeAcc:      proto.Uint32(0),
		ExtraTimeMax: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13001
}
