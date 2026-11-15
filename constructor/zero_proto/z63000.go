package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_63000(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_63000{
		RefreshList: []*protobuf.TECHNOLOGYREFRESH{},
		RefreshFlag: proto.Uint32(0),
		Catchup: &protobuf.TECHNOLOGYCATCHUP{
			Version:   proto.Uint32(0),
			Target:    proto.Uint32(0),
			Pursuings: []*protobuf.TECHPURSUING{},
		},
		Queue: []*protobuf.TECHNOLOGYINFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 63000
}
