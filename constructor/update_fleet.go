package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/fleet"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12103(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_12102{}

	if err := proto.Unmarshal(*msg, cs); err != nil {
		log.Fatal(err)
	}

	err := fleet.UpdateFleet(b.DB, int(cs.GetId()), cs.GetShipList())
	if err != nil {
		log.Fatal(err)
	}

	sc := &protobuf.SC_12103{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12103
}
