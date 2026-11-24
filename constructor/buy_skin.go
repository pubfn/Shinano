package constructor

import (
	"log"

	"Shinano/cfg"
	"Shinano/protobuf"
	"Shinano/repository/ship"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_16002(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_16001{}

	proto.Unmarshal(*msg, cs)

	skinID := cfg.SkinPool[cs.GetId()].Id

	if err := ship.BuySkin(b.DB, skinID); err != nil {
		log.Fatal(err)
	}

	sc := &protobuf.SC_16002{
		Result:   proto.Uint32(0),
		DropList: []*protobuf.DROPINFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 16002
}
