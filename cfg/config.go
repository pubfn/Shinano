// Package cfg asdl;fjasljf
package cfg

import (
	"encoding/json"
	"os"
)

type buffList struct {
	ID []int `json:"buff_list"`
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

var InstMsgID = map[uint32]instMsg{}

var SkillID = map[uint32]buffList{}

func LoadGameData() error {
	funcs := []func() error{
		loadskillid,
		loadInstData,
	}

	for _, action := range funcs {
		if err := action(); err != nil {
			return err
		}
	}
	return nil
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
