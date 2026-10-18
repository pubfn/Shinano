package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_15001(msg *[]byte) ([]byte, int) {
	sc_15001 := &protobuf.SC_15001{
		ItemList: []*protobuf.ITEMINFO{
			{Id: proto.Uint32(30011), Count: proto.Uint32(2)},
			{Id: proto.Uint32(50001), Count: proto.Uint32(10)},
			{Id: proto.Uint32(50002), Count: proto.Uint32(20)},
			{Id: proto.Uint32(59010), Count: proto.Uint32(1)},
			{Id: proto.Uint32(20001), Count: proto.Uint32(15)},
			{Id: proto.Uint32(17021), Count: proto.Uint32(4)},
			{Id: proto.Uint32(15003), Count: proto.Uint32(14)},
		},
		LimitList: []*protobuf.ITEMINFO{
			{Id: proto.Uint32(59010), Count: proto.Uint32(0)},
		},
		ItemMiscList: []*protobuf.ITEMMISC{},
	}

	data, err := proto.Marshal(sc_15001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 15001
}
