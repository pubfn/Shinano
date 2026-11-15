package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_26120(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_26120{
		WeeklyFree:    proto.Uint32(0),
		MonthlyTicket: proto.Uint32(0),
		Rooms:         []*protobuf.GAMEROOM{},
		PayCoinCount:  proto.Uint32(0),
		FirstEnter:    proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 26120
}
