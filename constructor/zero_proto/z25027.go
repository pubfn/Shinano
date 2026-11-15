package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_25027(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_25027{
		Level: proto.Uint32(1),
		Exp:   proto.Uint32(0),
		Slots: []*protobuf.COMMANDERHOMESLOT{},
		Clean: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 25027
}
