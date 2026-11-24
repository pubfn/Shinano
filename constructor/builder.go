package constructor

import (
	"database/sql"

	"Shinano/constructor/zero_proto"
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
			b.Build_10023,
		},
		10100: {
			zero_proto.Build_10101,
		},
		11001: {
			zero_proto.Build_11000,
			b.Build_11003,
			zero_proto.Build_11015,
			zero_proto.Build_11752,
			zero_proto.Build_22001,
			zero_proto.Build_26120,
			zero_proto.Build_25001,
			zero_proto.Build_17001,
			zero_proto.Build_22300,
			Build_12024,
			b.Build_12001,
			b.Build_12101,
			b.Build_12201,
			zero_proto.Build_63000,
			zero_proto.Build_63100,
			zero_proto.Build_64000,
			b.Build_13001,
			zero_proto.Build_13000,
			zero_proto.Build_13002,
			zero_proto.Build_13201,
			zero_proto.Build_16200,
			zero_proto.Build_33114,
			zero_proto.Build_14001,
			zero_proto.Build_14101,
			b.Build_15001,
			zero_proto.Build_20001,
			zero_proto.Build_20101,
			zero_proto.Build_11210,
			zero_proto.Build_20201,
			zero_proto.Build_19001,
			zero_proto.Build_12031,
			zero_proto.Build_28000,
			zero_proto.Build_30001,
			zero_proto.Build_30101,
			zero_proto.Build_50000,
			zero_proto.Build_11200,
			zero_proto.Build_11213,
			Build_11300,
			zero_proto.Build_11002,
		},
		11017: {
			zero_proto.Build_11018,
		},
		13101: {
			Build_13102,
		},
		11401: {
			Build_11402,
		},
		11701: {
			Build_11702,
		},
		11705: {
			Build_11706,
		},
		11019: {
			zero_proto.Build_11020,
		},
		12002: {
			b.Build_12003,
		},
		12022: {
			b.Build_12023,
		},
		12043: {
			b.Build_12044,
		},
		12045: {
			zero_proto.Build_12046,
		},
		12025: {
			b.Build_12026,
		},
		25026: {
			zero_proto.Build_25027,
		},
		34501: {
			zero_proto.Build_34502,
		},
		11202: {
			zero_proto.Build_11203,
		},
		23430: {
			zero_proto.Build_23431,
		},
		34001: {
			zero_proto.Build_34002,
		},
		18001: {
			zero_proto.Build_18002,
		},
		60037: {
			zero_proto.Build_60000,
		},
		62100: {
			zero_proto.Build_62101,
		},
		26101: {
			zero_proto.Build_26102,
		},
		11601: {
			zero_proto.Build_11602,
		},
		24020: {
			zero_proto.Build_24021,
		},
		27000: {
			zero_proto.Build_27001,
		},
		29001: {
			zero_proto.Build_29002,
		},
		11710: {
			zero_proto.Build_11711,
		},
		11603: {
			zero_proto.Build_11604,
		},
		60102: {
			zero_proto.Build_60103,
		},
		63317: {
			Build_63318,
		},
		27010: {
			zero_proto.Build_27011,
		},
		12406: {
			zero_proto.Build_12407,
		},
		12102: {
			b.Build_12103,
		},
		16104: {
			zero_proto.Build_16105,
		},
		12202: {
			b.Build_12203,
		},
		13505: {
			zero_proto.Build_13506,
		},
		15008: {
			zero_proto.Build_15009,
		},
		16001: {
			b.Build_16002,
		},
		11722: {
			zero_proto.Build_11723,
		},
		13107: {
			zero_proto.Build_13108,
		},
		13103: {
			Build_13104,
		},
	}
}
