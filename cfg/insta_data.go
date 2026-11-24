package cfg

import (
	"encoding/json"
	"os"
)

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

func loadInstData() error {
	tmp, err := os.ReadFile("cfg/json/activity_ins_language.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &InstMsgID)
}
