package ship

import (
	"database/sql"
	"time"

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
