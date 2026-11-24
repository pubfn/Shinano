package cfg

import (
	"encoding/json"
	"os"
)

type group struct {
	Id []int `json:"id"`
}

type gachaPool struct {
	Light   map[uint32]group `json:"light"`
	Heavy   map[uint32]group `json:"heavy"`
	Special map[uint32]group `json:"special"`
}

var GachaPool = gachaPool{}

func loadGachaPool() error {
	tmp, err := os.ReadFile("cfg/json/azur_lane_pool.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &GachaPool)
}
