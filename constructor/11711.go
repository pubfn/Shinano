package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11711(msg *[]byte) ([]byte, int) {
	sc_11711 := &protobuf.SC_11711{
		Groups: []*protobuf.JUUS_GROUP{},
	}

	data, err := proto.Marshal(sc_11711)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11711
}
