package packets

import (
	"encoding/binary"
	"io"
)

type Message struct {
	Len     int
	Id      int
	Index   int
	Message []byte
}

func ReadPacket(data []byte) (Message, error) {
	if len(data) < 5 {
		return Message{}, io.ErrUnexpectedEOF
	}
	msg := Message{
		Len:   int(binary.BigEndian.Uint16(data[0:2])),
		Id:    int(data[2])<<16 | int(data[3])<<8 | int(data[4]),
		Index: int(binary.BigEndian.Uint16(data[5:7])),
	}
	msg.Message = data[7:]
	return msg, nil
}

func BuildPacket(data *[]byte, id int, idx int) []byte {
	len := uint16(len(*data) + 5)
	packet := make([]byte, len+2)
	packet[0], packet[1] = byte(len>>8), byte(len)
	packet[2] = byte(id >> 16)
	packet[3] = byte(id >> 8)
	packet[4] = byte(id)
	packet[5], packet[6] = byte(idx>>8), byte(idx)
	copy(packet[7:], *data)
	return packet
}
