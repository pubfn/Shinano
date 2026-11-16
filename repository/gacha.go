package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand/v2"

	"Shinano/cfg"
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

// 1 special
// 2 light
// 3 heavy
func Casino(buildType int) int {
	var shipId int
	if buildType == 1 {
		rarity := rarityCheck(rand.Float64())
		shipId = cfg.GachaPool.Special[rarity].Id[rand.Uint32N(uint32(len(cfg.GachaPool.Special[rarity].Id)))]

		if shipId == 0 {
			fmt.Println(buildType)
			fmt.Println(rarity)
		}
	} else if buildType == 2 {
		rarity := rarityCheckLight(rand.Float64())
		shipId = cfg.GachaPool.Light[rarity].Id[rand.Uint32N(uint32(len(cfg.GachaPool.Light[rarity].Id)))]

		if shipId == 0 {
			fmt.Println(buildType)
			fmt.Println(rarity)
		}
	} else if buildType == 3 {
		rarity := rarityCheck(rand.Float64())
		shipId = cfg.GachaPool.Heavy[rarity].Id[rand.Uint32N(uint32(len(cfg.GachaPool.Heavy[rarity].Id)))]
		if shipId == 0 {
			fmt.Println(buildType)
			fmt.Println(rarity)
		}
	}

	if shipId == 0 {

		fmt.Println("casino", buildType)
		fmt.Println()
	}
	return shipId
}

func rarityCheckLight(num float64) uint32 {
	if num < 0.07 {
		return 4
	} else if num < 0.19 {
		return 3
	} else if num < 0.45 {
		return 2
	}
	return 1
}

func rarityCheck(num float64) uint32 {
	if num < 0.012 {
		return 5
	} else if num < 0.082 {
		return 4
	} else if num < 0.202 {
		return 3
	} else if num < 0.712 {
		return 2
	}
	return 1
}
