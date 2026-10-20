package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_23431(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_23431{
		Gold:          proto.Uint32(0),
		BuyNum:        proto.Uint32(0),
		MaxProfit:     proto.Uint32(0),
		AccProfit:     proto.Int32(0),
		AccBuyPrice:   proto.Uint32(0),
		PreBuyState:   proto.Uint32(0),
		ItemList:      []uint32{},
		PreTimestamp:  proto.Uint32(0),
		MatchTime:     proto.Uint32(0),
		IsForbidden:   proto.Uint32(0),
		GameNum:       proto.Uint32(0),
		InactiveNum:   proto.Uint32(0),
		InactiveState: proto.Uint32(0),
		BackForbidden: proto.Uint32(0),
		AccItemPrice:  proto.Uint32(0),
		GetReliefNum:  proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 23431
}
