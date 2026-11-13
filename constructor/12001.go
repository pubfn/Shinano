package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12001(msg *[]byte) ([]byte, int) {
	template := &protobuf.SHIPINFO{
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
	}

	ships, err := repository.LoadShips(b.DB, template)
	if err != nil {
		log.Fatal(err)
	}

	sc_12001 := &protobuf.SC_12001{Shiplist: ships}

	data, err := proto.Marshal(sc_12001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12001
}
