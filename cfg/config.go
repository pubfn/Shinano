// Package cfg asdl;fjasljf
package cfg

import (
	"encoding/json"
	"log"
	"os"
)

type bufflist struct {
	Id []int `json:"buff_list"`
}

var Skillid = map[uint32]bufflist{}

func LoadGameData() {
	loadskillid()
}

func loadskillid() {
	tmp, err := os.ReadFile("cfg/ship_data_template.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(tmp, &Skillid)
}
