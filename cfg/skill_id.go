package cfg

import (
	"encoding/json"
	"os"
)

type buffList struct {
	ID []int `json:"buff_list"`
}

var SkillID = map[uint32]buffList{}

func loadSkillId() error {
	tmp, err := os.ReadFile("cfg/json/ship_data_template.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &SkillID)
}
