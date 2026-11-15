package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12001(msg *[]byte) ([]byte, int) {
	ships, err := repository.LoadShips(b.DB)
	if err != nil {
		log.Fatal(err)
	}

	sc_12001 := &protobuf.SC_12001{Shiplist: ships}

	data, err := proto.Marshal(sc_12001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12001
}
