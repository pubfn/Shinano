package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_28000(msg *[]byte) ([]byte, int) {
	sc_28000 := &protobuf.SC_28000{
		Gifts:              []*protobuf.APARTMENT_GIFT{},
		Ships:              []*protobuf.APARTMENT_SHIP{},
		GiftDaily:          []*protobuf.APARTMENT_GIFT_SHOP{},
		GiftPermanent:      []*protobuf.APARTMENT_GIFT_SHOP{},
		FurnitureDaily:     []*protobuf.APARTMENT_GIFT_SHOP{},
		FurniturePermanent: []*protobuf.APARTMENT_GIFT_SHOP{},
		DailyVigorMax:      proto.Uint32(0),
		Rooms:              []*protobuf.APARTMENT_ROOM{},
		Ins:                []*protobuf.APARTMENT_INS{},
	}

	data, err := proto.Marshal(sc_28000)
	if err != nil {
		log.Fatal(err)
	}

	return data, 28000
}
