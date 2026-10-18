package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12101(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12101{
		GroupList: []*protobuf.GROUPINFO_P12{
			{Id: proto.Uint32(12), Name: proto.String(""), ShipList: []uint32{}, Commanders: []*protobuf.COMMANDERSINFO{}},
			{Id: proto.Uint32(11), Name: proto.String(""), ShipList: []uint32{}, Commanders: []*protobuf.COMMANDERSINFO{}},
			{Id: proto.Uint32(2), Name: proto.String(""), ShipList: []uint32{}, Commanders: []*protobuf.COMMANDERSINFO{}},
			{Id: proto.Uint32(1), Name: proto.String(""), ShipList: []uint32{1, 3, 2}, Commanders: []*protobuf.COMMANDERSINFO{}},
		},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12101
}
