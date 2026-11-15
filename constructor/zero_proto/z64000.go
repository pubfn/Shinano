package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_64000(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_64000{
		TechList:    []*protobuf.FLEETTECH{},
		TechsetList: []*protobuf.TECHSET{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 64000
}
