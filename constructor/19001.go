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
		ShipIdList:           []uint32{},
		FurnitureIdList: []*protobuf.FURNITUREINFO{
			{Id: proto.Uint32(1002), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1001), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1111), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1104), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1301), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1306), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1107), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1102), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1103), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1304), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1302), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1105), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1106), Count: proto.Uint32(1), GetTime: nil},
			{Id: proto.Uint32(1112), Count: proto.Uint32(1), GetTime: nil},
		},
		FloorNum:         proto.Uint32(1),
		ExpPos:           proto.Uint32(2),
		FurniturePutList: []*protobuf.FURFLOORPUTINFO{},
		NextTimestamp:    proto.Uint32(0),
		LoadExp:          proto.Uint32(0),
		LoadFood:         proto.Uint32(0),
		LoadTime:         proto.Uint32(1790705327),
		Name:             proto.String(""),
	}

	data, err := proto.Marshal(sc_19001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 19001
}
