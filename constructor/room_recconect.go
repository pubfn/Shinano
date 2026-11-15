package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11402(msg *[]byte) ([]byte, int) {
	roomId := protobuf.CS_11401{}

	if err := proto.Unmarshal(*msg, &roomId); err != nil {
		log.Fatal(err)
	}

	sc_11402 := &protobuf.SC_11402{
		Result: proto.Uint32(0),
		RoomId: proto.Uint32(roomId.GetRoomId()),
	}

	data, err := proto.Marshal(sc_11402)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11402
}
