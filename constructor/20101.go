package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_20101(msg *[]byte) ([]byte, int) {
	sc_20101 := &protobuf.SC_20101{
		Info: &protobuf.WEEKLY_INFO{
			Task: []*protobuf.WEEKLY_TASK_P20{
				{Id: proto.Uint32(10141), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10111), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10091), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10081), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10061), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10051), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10041), Progress: proto.Uint32(3)},
				{Id: proto.Uint32(10011), Progress: proto.Uint32(0)},
				{Id: proto.Uint32(10001), Progress: proto.Uint32(0)},
			},
			Pt:       proto.Uint32(0),
			RewardLv: proto.Uint32(0),
		},
	}

	data, err := proto.Marshal(sc_20101)
	if err != nil {
		log.Fatal(err)
	}

	return data, 20101
}
