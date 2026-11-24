package ship

import (
	"database/sql"

	"Shinano/cfg"
	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func loadSkills(b *sql.DB, id uint32) ([]*protobuf.SHIPSKILL, error) {
	skill, err := b.Query("select skillid, skilllvl, skillexp from skillidlist where shipid = ?", id)
	if err != nil {
		return nil, err
	}

	defer skill.Close()

	skillidlist := []*protobuf.SHIPSKILL{}

	for skill.Next() {
		var skillid, skilllvl, skillexp uint32

		if err = skill.Scan(&skillid, &skilllvl, &skillexp); err != nil {
			return nil, err
		}

		skillidlist = append(skillidlist, &protobuf.SHIPSKILL{SkillId: proto.Uint32(skillid), SkillLv: proto.Uint32(skilllvl), SkillExp: proto.Uint32(skillexp)})
	}

	return skillidlist, nil
}

func addSkills(b DBExecutor, shipId int, Id int) error {
	for _, v := range cfg.SkillID[uint32(shipId)].ID {
		_, err := b.Exec("insert into skillidlist (skillid, skilllvl, skillexp, shipid) values (?,?,?,?)", v, 1, 0, Id)
		if err != nil {
			return err
		}
	}
	return nil
}
