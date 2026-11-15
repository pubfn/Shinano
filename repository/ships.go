package repository

import (
	"database/sql"
	"fmt"

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

func AddShip(b DBExecutor, shipId int) error {
	fmt.Println("AddShip called with shipId:", shipId)

	if shipId == 0 {
		return fmt.Errorf("AddShip: shipId is zero")
	}
	_, err := b.Exec("insert into shiplist (templateId, level,exp,islocked,skinid) values (?,?,?,?,?)", shipId, 1, 0, 0, 0)
	if err != nil {
		return err
	}
	return nil
}

func CreateShip(b *sql.DB, shipId int) (*protobuf.SHIPINFO, error) {
	var id int
	ship := proto.Clone(template).(*protobuf.SHIPINFO)

	if err := b.QueryRow(`
    SELECT id FROM Shiplist ORDER BY id DESC LIMIT 1
    `).Scan(&id); err != nil {
		return nil, err
	}

	ship.IsLocked, ship.TemplateId, ship.Level, ship.Exp, ship.SkinId, ship.Id =
		proto.Uint32(0), proto.Uint32(uint32(shipId)), proto.Uint32(1), proto.Uint32(0), proto.Uint32(0), proto.Uint32(uint32(id+1))

	return ship, nil
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
			return nil, err
		}

		skillidlist = append(skillidlist, &protobuf.SHIPSKILL{SkillId: proto.Uint32(skillid), SkillLv: proto.Uint32(skilllvl), SkillExp: proto.Uint32(skillexp)})
	}

	return skillidlist, nil
}
