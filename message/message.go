// Package message suck dick
package message

import (
	"encoding/binary"
)

type Message struct {
	Len     int
	Id      int
	index   int
	Message []byte
}

func ReadMsg(data []byte) Message {
	msg := Message{
		Len:   int(binary.BigEndian.Uint16(data[0:2])),
		Id:    int(data[2])<<16 | int(data[3])<<8 | int(data[4]),
		index: int(binary.BigEndian.Uint16(data[5:7])),
	}

	msg.Message = data[7:msg.Len]
	return msg
}
