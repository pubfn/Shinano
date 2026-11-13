package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12026(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12026{
		Result: proto.Uint32(0),
		ShipList: []*protobuf.SHIPINFO{
			{
				Id:         proto.Uint32(2),
				TemplateId: proto.Uint32(307164),
				SkinId:     proto.Uint32(0),
				Level:      proto.Uint32(52),
				Exp:        proto.Uint32(0),
				IsLocked:   proto.Uint32(0),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},

					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(146),
				State:         proto.Uint32(0),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				Intimacy:      proto.Uint32(9000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1790705327),
				Propose:      proto.Uint32(1),
				Name:         proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(100), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
		},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12026
}
