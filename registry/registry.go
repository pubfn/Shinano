package registry

import (
	"fmt"
	"net"

	"Shinano/constructor"
	"Shinano/message"
)

type Handler func(data *[]byte) ([]byte, int)

var registry = map[int][]Handler{
	10800: {
		constructor.Build_10801,
	},
	10022: {
		constructor.Build_10023,
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
		constructor.Build_12001,
		constructor.Build_12101,
		constructor.Build_12201,
		constructor.Build_63000,
		constructor.Build_63100,
		constructor.Build_64000,
		constructor.Build_13001,
		constructor.Build_13000,
		constructor.Build_13201,
		constructor.Build_16200,
		constructor.Build_33114,
		constructor.Build_14001,
		constructor.Build_14101,
		constructor.Build_15001,
		constructor.Build_20101,
		constructor.Build_20201,
		constructor.Build_19001,
		constructor.Build_12031,
		constructor.Build_28000,
		constructor.Build_30001,
		constructor.Build_30101,
		constructor.Build_50000,
		constructor.Build_11200,
		constructor.Build_11213,
		constructor.Build_11300,
		constructor.Build_11002,
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
