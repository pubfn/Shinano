package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_26102(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_26102{
		Hubs: []*protobuf.MINIGAMEHUB{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 26102
}
