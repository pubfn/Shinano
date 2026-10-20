package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_27001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_27001{
		Result: proto.Uint32(0),
		Child: &protobuf.CHILD_INFO{
			Tid: proto.Uint32(0),

			CurTime: &protobuf.CHILD_TIME{
				Month: proto.Uint32(0),
				Week:  proto.Uint32(0),
				Day:   proto.Uint32(0),
			},

			Mood:       proto.Uint32(0),
			Money:      proto.Uint32(0),
			SiteNumber: proto.Uint32(0),

			Attrs: []*protobuf.CHILD_ATTR{},
			Favor: &protobuf.CHILD_FAVOR{
				Lv:  proto.Uint32(0),
				Exp: proto.Uint32(0),
			},
			Items:       []*protobuf.CHILD_ITEM{},
			PlanHistory: []*protobuf.CHILD_PLAN_HISTORY{},
			Memorys:     []uint32{},
			Plans:       []*protobuf.CHILD_PLAN_CELL{},
			Polaroids:   []*protobuf.CHILD_POLAROID{},

			Target:              proto.Uint32(0),
			Tasks:               []*protobuf.CHILD_TASK{},
			RealizedWish:        []uint32{},
			Buffs:               []*protobuf.CHILD_BUFF{},
			UserName:            proto.String(""),
			SpecEvents:          []uint32{},
			CanTriggerHomeEvent: proto.Uint32(0),
			HomeEvents:          []uint32{},
			DiscountEventId:     []uint32{},
			Shop:                []*protobuf.CHILD_SHOP_DATA{},
			OptionRecords:       []*protobuf.CHILD_OPTION_RECORD{},
			FavorAwardHistory:   []uint32{},

			IsEnding:                proto.Uint32(0),
			NewGamePlusCount:        proto.Uint32(0),
			HadTargetStageAward:     proto.Uint32(0),
			HadAdjustment:           proto.Uint32(0),
			IsSpecialSecretaryValid: proto.Uint32(0),
			EndingBuyCount:          proto.Uint32(0),
			MemoryBuyCount:          proto.Uint32(0),
			PolaroidBuyCount:        proto.Uint32(0),
		},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 27001
}
