package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12026(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_12025{}
	proto.Unmarshal(*msg, cs)

	shipList := []*protobuf.SHIPINFO{}

	for _, v := range cs.GetPosList() {
		shipId := repository.Casino(int(v))
		ship, err := repository.CreateShip(b.DB, shipId)
		if err != nil {
			log.Fatal(err)
		}
		shipList = append(shipList, ship)
		repository.AddShip(b.DB, shipId)
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
