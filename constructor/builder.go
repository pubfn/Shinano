package constructor

import (
	"database/sql"
)

type Handler func(data *[]byte) ([]byte, int)

var Registry map[int][]Handler

type Builder struct {
	DB *sql.DB
}

func NewBuilder(db *sql.DB) *Builder {
	return &Builder{
		DB: db,
	}
}

func NewRegistry(b *Builder) {
	Registry = map[int][]Handler{
		10022: {
			b.Build_10023, // recive accId
		},
		10100: {
			Build_10101, // ping
		},
		11001: {
			Build_11000,
			b.Build_11003,
			Build_11015,
			Build_11752,
			Build_22001,
			Build_26120,
			Build_25001,
			Build_17001,
			Build_22300,
			Build_12024,
			b.Build_12001, // ship list info
			Build_12101,
			Build_12201,
			Build_63000,
			Build_63100,
			Build_64000,
			Build_13001,
			Build_13000,
			Build_13002,
			Build_13201,
			Build_16200,
			Build_33114,
			Build_14001,
			Build_14101,
			Build_15001, // inventory mb
			Build_20001,
			Build_20101,
			Build_11210,
			Build_20201,
			Build_19001,
			Build_12031,
			Build_28000,
			Build_30001,
			Build_30101,
			Build_50000,
			Build_11200,
			Build_11213,
			Build_11300, // patch note
			Build_11002,
		},
		25026: {
			Build_25027,
		},
		34501: {
			Build_34502,
		},
		11202: {
			Build_11203,
		},
		23430: {
			Build_23431,
		},
		34001: {
			Build_34002,
		},
		18001: {
			Build_18002,
		},
		60037: {
			Build_60000,
		},
		62100: {
			Build_62101,
		},
		26101: {
			Build_26102,
		},
		11601: {
			Build_11602,
		},
		24020: {
			Build_24021,
		},
		27000: {
			Build_27001,
		},
		29001: {
			Build_29002,
		},
		11710: {
			Build_11711,
		},
		11603: {
			Build_11604,
		},
		60102: {
			Build_60103,
		},
		27010: {
			Build_27011,
		},
		12406: {
			Build_12407, // here
		},
		16104: {
			Build_16105,
		},
		13505: {
			Build_13506,
		},
		15008: {
			Build_15009,
		},
		11722: {
			Build_11723,
		},
	}
}
