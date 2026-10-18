package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12024(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12024{
		WorklistCount: proto.Uint32(2),
		WorklistList:  []*protobuf.BUILDINFO{},
		DrawCount_1:   proto.Uint32(0),
		DrawCount_10:  proto.Uint32(0),
		ExchangeCount: proto.Uint32(7),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12024
}
