package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12031(msg *[]byte) ([]byte, int) {
	sc_12031 := &protobuf.SC_12031{
		EnergyAutoIncreaseTime: proto.Uint32(1792009607),
	}

	data, err := proto.Marshal(sc_12031)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12031
}
