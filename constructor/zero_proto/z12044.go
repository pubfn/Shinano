package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12044(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12044{
		InfoList: []*protobuf.BUILD_INFO{{
			Pos: proto.Uint32(0),
			Tid: proto.Uint32(0),
		}},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12044
}
