package cfg

import (
	"encoding/json"
	"os"
)

type skin struct {
	Id int `json:"id"`
}

var SkinPool = map[uint32]skin{}

func loadSkinPool() error {
	tmp, err := os.ReadFile("cfg/json/ship_skin_template.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &SkinPool)
}
