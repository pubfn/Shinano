package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12023(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_12022{}

	if err := proto.Unmarshal(*msg, cs); err != nil {
		log.Fatal(err)
	}

	isLocked := cs.GetIsLocked()

	for _, v := range cs.GetShipIdList() {
		if err := repository.LockShip(b.DB, v, int(isLocked)); err != nil {
			log.Fatal(err)
		}
	}

	sc := &protobuf.SC_12023{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12023
}
