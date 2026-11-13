package repository

import (
	"database/sql"
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func LoadShips(b *sql.DB, template *protobuf.SHIPINFO) ([]*protobuf.SHIPINFO, error) {
	ship, err := b.Query("select * from Shiplist")
	if err != nil {
		return nil, err
	}
	shipList := []*protobuf.SHIPINFO{}

	defer ship.Close()
	for ship.Next() {
		var Id, TemplateId, Level, Exp, IsLocked, SkinId uint32
		buf := proto.Clone(template).(*protobuf.SHIPINFO)

		if err = ship.Scan(&Id, &TemplateId, &Level, &Exp, &IsLocked, &SkinId); err != nil {
			return nil, err
		}

		skillidlist, err := loadSkills(b, TemplateId)
		if err != nil {
			return nil, err
		}

		buf.Id, buf.TemplateId, buf.Level, buf.Exp, buf.IsLocked, buf.SkinId, buf.SkillIdList =
			&Id, &TemplateId, &Level, &Exp, &IsLocked, &SkinId, skillidlist

		shipList = append(shipList, buf)
	}

	return shipList, nil
}

func loadSkills(b *sql.DB, shipid uint32) ([]*protobuf.SHIPSKILL, error) {
	skill, err := b.Query("select skillid, skilllvl, skillexp from skillidlist where shipid = ?", shipid)
	if err != nil {
		return nil, err
	}

	defer skill.Close()

	skillidlist := []*protobuf.SHIPSKILL{}

	for skill.Next() {
		var skillid, skilllvl, skillexp uint32

		if err = skill.Scan(&skillid, &skilllvl, &skillexp); err != nil {
			log.Fatal(err)
		}

		skillidlist = append(skillidlist, &protobuf.SHIPSKILL{SkillId: proto.Uint32(skillid), SkillLv: proto.Uint32(skilllvl), SkillExp: proto.Uint32(skillexp)})
	}

	return skillidlist, nil
}
