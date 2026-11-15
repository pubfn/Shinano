package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_17001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_17001{
		FinishList:    []uint32{},
		ProgressList:  []*protobuf.ACHIEVEMENT_INFO{},
		ShipInfoList:  []*protobuf.SHIP_STATISTICS_INFO{},
		ShipAwardList: []*protobuf.SHIP_STATISTICS_AWARD{},
		DailyDiscuss:  proto.Uint32(0),
		TransformList: []uint32{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 17001
}
