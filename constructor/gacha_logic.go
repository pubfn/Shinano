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

	for range cs.GetPosList() {
		shipId := repository.Casino(int(cs.GetType()))
		ship, err := repository.CreateShip(b.DB, shipId)
		if err != nil {
			log.Fatal(err)
		}
		shipList = append(shipList, ship)
		if err := repository.AddShip(b.DB, shipId); err != nil {
			log.Fatal(err)
		}
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
