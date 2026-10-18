package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_63100(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_63100{
		BlueprintList:            []*protobuf.BLUPRINTINFO{},
		ColdTime:                 proto.Uint32(0),
		DailyCatchupStrengthen:   proto.Uint32(0),
		DailyCatchupStrengthenUr: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 63100
}
