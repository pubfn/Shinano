package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_20101(msg *[]byte) ([]byte, int) {
	sc_20101 := &protobuf.SC_20101{
		Info: &protobuf.WEEKLY_INFO{
			Task:     []*protobuf.WEEKLY_TASK_P20{},
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
