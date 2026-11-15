package cfg

import (
	"encoding/json"
	"os"
)

type buffList struct {
	ID []int `json:"buff_list"`
}

type group struct {
	Id []int `json:"id"`
}

type instMsg struct {
	Text    string         `json:"text"`
	Discuss map[uint32]msg `json:"discuss"`
	Op      map[uint32]msg `json:"op"`
}
type msg struct {
	Discuss string   `json:"discuss"`
	Reply   []string `json:"reply"`
}

type gachaPool struct {
	Light   map[uint32]group `json:"light"`
	Heavy   map[uint32]group `json:"heavy"`
	Special map[uint32]group `json:"special"`
}

var InstMsgID = map[uint32]instMsg{}

var SkillID = map[uint32]buffList{}

var GachaPool = gachaPool{}

func LoadGameData() error {
	funcs := []func() error{
		loadskillid,
		loadInstData,
		loadGachaPool,
	}

	for _, action := range funcs {
		if err := action(); err != nil {
			return err
		}
	}
	return nil
}

func loadGachaPool() error {
	tmp, err := os.ReadFile("cfg/azur_lane_pool.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &GachaPool)
}

func loadInstData() error {
	tmp, err := os.ReadFile("cfg/activity_ins_language.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &InstMsgID)
}

func loadskillid() error {
	tmp, err := os.ReadFile("cfg/ship_data_template.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &SkillID)
}
