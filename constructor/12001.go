package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12001(msg *[]byte) ([]byte, int) {
	var Id, TemplateId, Level, Exp, IsLocked, SkinId uint32
	var skillid, skilllvl, skillexp uint32
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

	ship, err := b.DB.Query("select * from Shiplist")

	sc_12001 := &protobuf.SC_12001{Shiplist: []*protobuf.SHIPINFO{}}

	defer ship.Close()
	for ship.Next() {
		buf := proto.Clone(template).(*protobuf.SHIPINFO)

		ship.Scan(&Id, &TemplateId, &Level, &Exp, &IsLocked, &SkinId)

		skill, err := b.DB.Query("select skillid, skilllvl, skillexp from skillidlist where shipid = ?", TemplateId)
		if err != nil {
			log.Fatal(err)
		}

		skillidlist := []*protobuf.SHIPSKILL{}

		defer skill.Close()
		for skill.Next() {
			err = skill.Scan(&skillid, &skilllvl, &skillexp)
			if err != nil {
				log.Fatal(err)
			}

			skillidlist = append(skillidlist, &protobuf.SHIPSKILL{SkillId: proto.Uint32(skillid), SkillLv: proto.Uint32(skilllvl), SkillExp: proto.Uint32(skillexp)})
		}

		buf.Id, buf.TemplateId, buf.Level, buf.Exp, buf.IsLocked, buf.SkinId, buf.SkillIdList =
			&Id, &TemplateId, &Level, &Exp, &IsLocked, &SkinId, skillidlist

		sc_12001.Shiplist = append(sc_12001.Shiplist, buf)
	}

	data, err := proto.Marshal(sc_12001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12001
}
