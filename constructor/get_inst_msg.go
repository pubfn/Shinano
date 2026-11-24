package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/insta"

	"google.golang.org/protobuf/proto"
)

func Build_11702(msg *[]byte) ([]byte, int) {
	id := &protobuf.CS_11701{}

	if err := proto.Unmarshal(*msg, id); err != nil {
		log.Fatal(err)
	}

	instMsg := insta.ParseId(id.GetId())

	sc_11604 := &protobuf.SC_11702{
		Result: proto.Uint32(0),
		Data:   instMsg,
	}

	data, err := proto.Marshal(sc_11604)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11702
}
