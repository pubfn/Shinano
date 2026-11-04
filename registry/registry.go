package registry

import (
	"fmt"
	"net"

	"Shinano/constructor"
	"Shinano/message"
)

type Handler func(data *[]byte) ([]byte, int)

var registry = map[int][]Handler{
	10022: {
		constructor.Build_10023, // recive accId
	},
	11001: {
		constructor.Build_11000,
		constructor.Build_11003,
		constructor.Build_11015,
		constructor.Build_11752,
		constructor.Build_22001,
		constructor.Build_26120,
		constructor.Build_25001,
		constructor.Build_17001,
		constructor.Build_22300,
		constructor.Build_12024,
		constructor.Build_12001, // character info
		constructor.Build_12101,
		constructor.Build_12201,
		constructor.Build_63000,
		constructor.Build_63100,
		constructor.Build_64000,
		constructor.Build_13001,
		constructor.Build_13000,
		constructor.Build_13002,
		constructor.Build_13201,
		constructor.Build_16200,
		constructor.Build_33114,
		constructor.Build_14001,
		constructor.Build_14101,
		constructor.Build_15001, // inventory mb
		constructor.Build_20001,
		constructor.Build_20101,
		constructor.Build_11210,
		constructor.Build_20201,
		constructor.Build_19001,
		constructor.Build_12031,
		constructor.Build_28000,
		constructor.Build_30001,
		constructor.Build_30101,
		constructor.Build_50000,
		constructor.Build_11200,
		constructor.Build_11213,
		constructor.Build_11300, // patch note
		constructor.Build_11002,
	},
	25026: {
		constructor.Build_25027,
	},
	34501: {
		constructor.Build_34502,
	},
	11202: {
		constructor.Build_11203,
	},
	23430: {
		constructor.Build_23431,
	},
	34001: {
		constructor.Build_34002,
	},
	18001: {
		constructor.Build_18002,
	},
	60037: {
		constructor.Build_60000,
	},
	62100: {
		constructor.Build_62101,
	},
	26101: {
		constructor.Build_26102,
	},
	11601: {
		constructor.Build_11602,
	},
	24020: {
		constructor.Build_24021,
	},
	27000: {
		constructor.Build_27001,
	},
	29001: {
		constructor.Build_29002,
	},
	11710: {
		constructor.Build_11711,
	},
	11603: {
		constructor.Build_11604,
	},
	60102: {
		constructor.Build_60103,
	},
	27010: {
		constructor.Build_27011,
	},
	12406: {
		constructor.Build_12407, // here
	},
	16104: {
		constructor.Build_16105,
	},
	13505: {
		constructor.Build_13506,
	},
	15008: {
		constructor.Build_15009,
	},
	11722: {
		constructor.Build_11723,
	},
}

func Hande(id int, data *[]byte, c net.Conn, idx int) {
	handlers := registry[id]

	for _, handler := range handlers {
		data, id := handler(data)
		msg := message.BuildPacket(&data, id, idx)
		fmt.Println("Sent message:", id)
		c.Write(msg)
	}
}
