package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_10023(msg *[]byte) ([]byte, int) {
	sc_10023 := &protobuf.SC_10023{
		Result:       proto.Uint32(0),
		UserId:       proto.Uint32(478112442),
		ServerTicket: proto.String("179200981798c7c4afcb98dd360b35d5e4ed40268b"),
		DbLoad:       nil,
		ServerLoad:   nil,
	}

	data, err := proto.Marshal(sc_10023)
	if err != nil {
		log.Fatal(err)
	}

	return data, 10023
}
