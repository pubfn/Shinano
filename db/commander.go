package db

type CommanderData struct {
	id    int    `json:"id"`
	name  string `json:"name"`
	level int    `json:"level"`
	exp   int    `json:"exp"`
}

var Commander = &CommanderData{}
