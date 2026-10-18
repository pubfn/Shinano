package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12001(msg *[]byte) ([]byte, int) {
	sc_12001 := &protobuf.SC_12001{
		Shiplist: []*protobuf.SHIPINFO{
			{
				Id: proto.Uint32(8), TemplateId: proto.Uint32(204021),
				Level: proto.Uint32(1), Exp: proto.Uint32(0),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(24100), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(21200), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(150),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(0),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(10280), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1791797103),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(7), TemplateId: proto.Uint32(103031),
				Level: proto.Uint32(1), Exp: proto.Uint32(0),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(13000), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(6000), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(150),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(0),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(1002), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1791797103),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(6), TemplateId: proto.Uint32(303031),
				Level: proto.Uint32(1), Exp: proto.Uint32(0),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(33000), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(35200), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(150),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(0),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(2011), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1791797103),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(5), TemplateId: proto.Uint32(103071),
				Level: proto.Uint32(1), Exp: proto.Uint32(0),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(13000), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(16100), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(150),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(1),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(10120), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
					{SkillId: proto.Uint32(4021), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1791365784),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(4), TemplateId: proto.Uint32(201101),
				Level: proto.Uint32(1), Exp: proto.Uint32(0),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(21300), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(5100), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(150),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(1),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(4071), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1790707624),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(3), TemplateId: proto.Uint32(201241),
				Level: proto.Uint32(4), Exp: proto.Uint32(62),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(21400), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(146),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(1),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(5031), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5010), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1790705339),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(1), TemplateId: proto.Uint32(201211),
				Level: proto.Uint32(5), Exp: proto.Uint32(324),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(21400), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(5201), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(146),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(0),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(10250), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(9000), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1790705327),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
			{
				Id: proto.Uint32(2), TemplateId: proto.Uint32(106011),
				Level: proto.Uint32(4), Exp: proto.Uint32(392),
				EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
					{Id: proto.Uint32(17000), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(19000), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
					{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
				},
				Energy:        proto.Uint32(146),
				State:         &protobuf.SHIPSTATE{State: proto.Uint32(0)},
				IsLocked:      proto.Uint32(0),
				TransformList: []*protobuf.TRANSFORM_INFO{},
				SkillIdList: []*protobuf.SHIPSKILL{
					{SkillId: proto.Uint32(3001), SkillLv: proto.Uint32(1), SkillExp: proto.Uint32(0)},
				},
				Intimacy: proto.Uint32(5014), Proficiency: proto.Uint32(0),
				StrengthList: []*protobuf.STRENGTH_INFO{},
				CreateTime:   proto.Uint32(1790705327),
				SkinId:       proto.Uint32(0), Propose: proto.Uint32(0),
				Name: proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
				MaxLevel: proto.Uint32(70), CommonFlag: proto.Uint32(0),
				ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
				SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
			},
		},
	}

	data, err := proto.Marshal(sc_12001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12001
}
