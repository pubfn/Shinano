package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_19001(msg *[]byte) ([]byte, int) {
	sc_19001 := &protobuf.SC_19001{
		Lv:                   proto.Uint32(1),
		Food:                 proto.Uint32(0),
		FoodMaxIncrease:      proto.Uint32(0),
		FoodMaxIncreaseCount: proto.Uint32(0),
		FurnitureIdList:      []*protobuf.FURNITUREINFO{},
		FloorNum:             proto.Uint32(1),
		ExpPos:               proto.Uint32(2),
		FurniturePutList:     []*protobuf.FURFLOORPUTINFO{},
		NextTimestamp:        proto.Uint32(0),
		LoadExp:              proto.Uint32(0),
		LoadFood:             proto.Uint32(0),
		LoadTime:             proto.Uint32(1790705327),
		Name:                 proto.String(""),
	}

	data, err := proto.Marshal(sc_19001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 19001
}
