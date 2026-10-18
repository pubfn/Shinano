package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13000(msg *[]byte) ([]byte, int) {
	sc_13000 := &protobuf.SC_13000{
		DailyRepairCount: proto.Uint32(0),
		CurrentChapter:   nil,
	}

	data, err := proto.Marshal(sc_13000)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13000
}
