package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_62101(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_62101{
		Technologys: []*protobuf.GUILD_TECHNOLOGY{
			{
				Id:       proto.Uint32(6004),
				State:    proto.Uint32(0),
				Progress: proto.Uint32(40),
			},
			{
				Id:       proto.Uint32(5003),
				State:    proto.Uint32(0),
				Progress: proto.Uint32(50),
			},
			{
				Id:       proto.Uint32(4003),
				State:    proto.Uint32(0),
				Progress: proto.Uint32(50),
			},
			{
				Id:       proto.Uint32(3008),
				State:    proto.Uint32(0),
				Progress: proto.Uint32(24),
			},
			{
				Id:       proto.Uint32(2008),
				State:    proto.Uint32(0),
				Progress: proto.Uint32(144),
			},
		},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 62101
}
