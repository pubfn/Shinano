package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

var keys uint32 = 10

func Build_40002(msg *[]byte) ([]byte, int) {
	keys++
	sc := &protobuf.SC_40002{
		Key:             proto.Uint32(keys),
		Result:          proto.Uint32(0),
		DropPerformance: []*protobuf.DROPPERFORMANCE{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 40002
}
