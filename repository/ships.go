package repository

import (
	"database/sql"
	"time"

	"Shinano/cfg"
	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

type DBExecutor interface {
	Exec(query string, args ...any) (sql.Result, error)
}

var template = &protobuf.SHIPINFO{
	EquipInfoList: []*protobuf.EQUIPSKIN_INFO{
		{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
		{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
		{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
		{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
		{Id: proto.Uint32(0), SkinId: proto.Uint32(0)},
	},
	Energy:        proto.Uint32(150),
	State:         proto.Uint32(0),
	TransformList: []*protobuf.TRANSFORM_INFO{},
	Intimacy:      proto.Uint32(9000), Proficiency: proto.Uint32(0),
	StrengthList: []*protobuf.STRENGTH_INFO{},
	CreateTime:   proto.Uint32(uint32(time.Now().Unix())),
	Propose:      proto.Uint32(1),
	Name:         proto.String(""), ChangeNameTimestamp: proto.Uint32(0),
	MaxLevel: proto.Uint32(100), CommonFlag: proto.Uint32(0),
	ActivityNpc: proto.Uint32(0), MetaRepairList: []uint32{},
	SkinShadowList: []*protobuf.KVDATA{}, CharRandomFlag: []uint32{},
}

func AddShip(b DBExecutor, shipId int) error {
	result, err := b.Exec("insert or ignore into shiplist (templateId, level,exp,islocked,skinid) values (?,?,?,?,?)", shipId, 1, 0, 0, 0)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if err := addSkills(b, shipId, int(id)); err != nil {
		return err
	}
	return nil
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

func GetAllSkins(b *sql.DB) ([]uint32, error) {
	rows, err := b.Query("select * from skinlist")
	skins := []uint32{}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var skinid uint32
		if err := rows.Scan(&skinid); err != nil {
			return nil, err
		}
		skins = append(skins, skinid)
	}

	return skins, err
}

func BuySkin(b *sql.DB, skinId int) error {
	_, err := b.Exec("insert into skinlist (skinid) values (?)", skinId)
	if err != nil {
		return err
	}

	return nil
}

func CreateShip(b *sql.DB, shipId int) (*protobuf.SHIPINFO, error) {
	var id int
	ship := proto.Clone(template).(*protobuf.SHIPINFO)

	if err := AddShip(b, shipId); err != nil {
		return nil, err
	}

	if err := b.QueryRow(`
    SELECT id FROM Shiplist ORDER BY id DESC LIMIT 1
    `).Scan(&id); err != nil {
		return nil, err
	}

	skills, err := loadSkills(b, uint32(id))
	if err != nil {
		return nil, err
	}

	ship.IsLocked, ship.TemplateId, ship.Level, ship.Exp, ship.SkinId, ship.Id, ship.SkillIdList =
		proto.Uint32(0), proto.Uint32(uint32(shipId)), proto.Uint32(1), proto.Uint32(0), proto.Uint32(0),
		proto.Uint32(uint32(id)), skills
	return ship, nil
}

func LockShip(b *sql.DB, id uint32, IsLocked int) error {
	_, err := b.Exec("update shiplist set islocked = ? where id = ?", IsLocked, id)
	if err != nil {
		return err
	}
	return nil
}

func ApplySkin(b *sql.DB, id int, skinId int) error {
	_, err := b.Exec("update shiplist set skinid = ? where id = ?", skinId, id)
	if err != nil {
		return err
	}
	return nil
}

func LoadShips(b *sql.DB) ([]*protobuf.SHIPINFO, error) {
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

		skillidlist, err := loadSkills(b, Id)
		if err != nil {
			return nil, err
		}

		buf.Id, buf.TemplateId, buf.Level, buf.Exp, buf.IsLocked, buf.SkinId, buf.SkillIdList =
			&Id, &TemplateId, &Level, &Exp, &IsLocked, &SkinId, skillidlist

		shipList = append(shipList, buf)
	}

	return shipList, nil
}

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
