package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11213(msg *[]byte) ([]byte, int) {
	sc_11213 := &protobuf.SC_11213{
		ActivityReTimerId: proto.Uint32(1),
		ActivityReId:      proto.Uint32(7),
		FinishReId:        []uint32{7},
	}

	data, err := proto.Marshal(sc_11213)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11213
}
