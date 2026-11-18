package repository

import (
	"database/sql"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func UpdateFleet(b *sql.DB, fleetId int, shipList []uint32) error {
	tx, err := b.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("delete from fleet where id = ?", fleetId)
	if err != nil {
		return err
	}

	for _, ship := range shipList {
		_, err = tx.Exec("insert into fleet (id, ship) values (?, ?)", fleetId, ship)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func GetFleet(b *sql.DB) ([]*protobuf.GROUPINFO_P12, error) {
	ids := []int{1, 2, 11, 12} // fleet formation ids
	fleet := []*protobuf.GROUPINFO_P12{}

	for _, id := range ids {
		err := func() error {
			rows, err := b.Query("select ship from fleet where id = ?", id)
			if err != nil {
				return err
			}
			defer rows.Close()

			buf := &protobuf.GROUPINFO_P12{
				Id:         proto.Uint32(uint32(id)),
				ShipList:   []uint32{},
				Name:       proto.String(""),
				Commanders: []*protobuf.COMMANDERSINFO{},
			}

			for rows.Next() {
				var shipid uint32
				err := rows.Scan(&shipid)
				if err != nil {
					return err
				}

				buf.ShipList = append(buf.ShipList, shipid)
			}

			fleet = append(fleet, buf)
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}

	return fleet, nil
}
