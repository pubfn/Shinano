package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_60103(msg *[]byte) ([]byte, int) {
	sc_60103 := &protobuf.SC_60103{
		UserInfo: &protobuf.USER_GUILD_INFO{
			DonateCount:    proto.Uint32(1),
			DonateTasks:    []uint32{},
			BenefitTime:    proto.Uint32(0),
			TechId:         []uint32{},
			WeeklyTaskFlag: proto.Uint32(0),
			ExtraDonate:    proto.Uint32(0),
			ExtraOperation: proto.Uint32(0),
		},
	}

	data, err := proto.Marshal(sc_60103)
	if err != nil {
		log.Fatal(err)
	}

	return data, 60103
}
