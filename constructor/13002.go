package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_13002(msg *[]byte) ([]byte, int) {
	sc_13002 := &protobuf.SC_13002{
		CollectionList: []*protobuf.COLLECTIONINFO{},
		MaxTeam:        proto.Uint32(2),
	}

	data, err := proto.Marshal(sc_13002)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13002
}
