package repository

import (
	"database/sql"
	"fmt"
	"math/rand/v2"

	"Shinano/cfg"
)

func GetBuildInfo(b *sql.DB, id int) (*int, error) {
	var buildType int
	if err := b.QueryRow("select type from gacha where id = ?", id).Scan(&buildType); err != nil {
		return nil, err
	}

	return &buildType, nil
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
