package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_34002(msg *[]byte) ([]byte, int) {
	sc_33114 := &protobuf.SC_34002{
		MetaShipList: []*protobuf.META_SHIP_INFO{},
	}
	client := &protobuf.CS_34001{}
	err := proto.Unmarshal(*msg, client)
	fields := client.GroupId
	for _, v := range fields {
		sc_33114.MetaShipList = append(sc_33114.MetaShipList, &protobuf.META_SHIP_INFO{
			GroupId:   proto.Uint32(v),
			Pt:        proto.Uint32(0),
			FetchList: []uint32{},
		})
	}

	data, err := proto.Marshal(sc_33114)
	if err != nil {
		log.Fatal(err)
	}

	return data, 34002
}
