package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/fleet"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12101(msg *[]byte) ([]byte, int) {
	fleetInfo, err := fleet.GetFleet(b.DB)
	if err != nil {
		log.Fatal(err)
	}

	sc := &protobuf.SC_12101{
		GroupList: fleetInfo,
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12101
}
