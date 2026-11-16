package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

// gacha
func (b *Builder) Build_12026(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_12025{}
	proto.Unmarshal(*msg, cs)

	shipList := []*protobuf.SHIPINFO{}

	shipIds, err := repository.DrawBuild(b.DB, cs.GetPosList())
	if err != nil {
		log.Fatal("gacha claim 1", err)
	}
	for _, shipId := range shipIds {
		ship, err := repository.CreateShip(b.DB, shipId)
		if err != nil {
			log.Fatal("gacha claim", err)
		}

		shipList = append(shipList, ship)
	}

	sc := &protobuf.SC_12026{
		Result:   proto.Uint32(0),
		ShipList: shipList,
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12026
}
