package gacha

import (
	"database/sql"
	"errors"
	"fmt"
)

type buildInfo struct {
	Pos    uint32
	ShipId uint32
}

func GetBuildInfo(b *sql.DB) ([]buildInfo, error) {
	build := make([]buildInfo, 0, 10)

	rows, err := b.Query(`
        SELECT shipid
        FROM gacha
        ORDER BY id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pos uint32

	for rows.Next() {
		var shipId uint32

		if err := rows.Scan(&shipId); err != nil {
			return nil, err
		}

		pos++

		build = append(build, buildInfo{
			Pos:    pos,
			ShipId: shipId,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return build, nil
}

func DrawBuild(b *sql.DB, pos []uint32) ([]int, error) {
	tx, err := b.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	shipIds := make([]int, 0, len(pos))
	ids := make([]int64, 0, len(pos))

	for _, position := range pos {

		var id int64
		var shipId int

		err := tx.QueryRow(`
            SELECT id, shipid
            FROM gacha
            ORDER BY id
            LIMIT 1 OFFSET ?
        `, position-1).Scan(&id, &shipId)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("позиция %d не найдена", position)
		}

		if err != nil {
			return nil, err
		}

		ids = append(ids, id)
		shipIds = append(shipIds, shipId)
	}

	for _, id := range ids {
		_, err := tx.Exec(
			"DELETE FROM gacha WHERE id = ?",
			id,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return shipIds, nil
}

func AddBuild(b *sql.DB, shipId int) error {
	_, err := b.Exec("insert into gacha (shipid) values (?)", shipId)
	if err != nil {
		return err
	}
	return nil
}
