package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12203(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_12202{}

	if err := proto.Unmarshal(*msg, cs); err != nil {
		log.Fatal(err)
	}

	if err := repository.ApplySkin(b.DB, int(cs.GetShipId()), int(cs.GetSkinId())); err != nil {
		log.Fatal(err)
	}

	sc := &protobuf.SC_12203{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12203
}
