package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_25001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_25001{
		Commanders: []*protobuf.COMMANDERINFO{},
		Box:        []*protobuf.COMMANDERBOXINFO{},
		UsageCount: proto.Uint32(0),
		Presets:    []*protobuf.PRESETFLEET{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 25001
}
