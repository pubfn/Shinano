package cfg

import (
	"encoding/json"
	"fmt"
	"os"
)

type Enemy struct {
	Id     int
	Weight int
	Som    int
}

type Grid struct {
	X        int
	Y        int
	Active   bool
	GridType int
}

type chapter struct {
	BossId             []int   `json:"boss_expedition_id"`
	BossSpawnKillCount int     `json:"boss_refresh"`
	Grid               []Grid  `json:"grids"`
	Enemy              []Enemy `json:"expedition_id_weight_list"`
}

var ChapterData = map[uint32]chapter{}

func loadChapterData() error {
	tmp, err := os.ReadFile("cfg/json/chapter_template.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(tmp, &ChapterData)
}

func (e *Enemy) UnmarshalJSON(data []byte) error {
	var raw []int

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 3 {
		return fmt.Errorf("invalid Enemy length: expected 3, got %d", len(raw))
	}

	e.Id = raw[0]
	e.Weight = raw[1]
	e.Som = raw[2]

	return nil
}

func (g *Grid) UnmarshalJSON(data []byte) error {
	var raw []json.RawMessage

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw) != 4 {
		return fmt.Errorf("invalid len", len(raw))
	}

	if err := json.Unmarshal(raw[0], &g.X); err != nil {
		return err
	}

	if err := json.Unmarshal(raw[1], &g.Y); err != nil {
		return err
	}

	if err := json.Unmarshal(raw[2], &g.Active); err != nil {
		g.Active = false
	}

	if err := json.Unmarshal(raw[3], &g.GridType); err != nil {
		return err
	}
	return nil
}
