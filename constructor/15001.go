package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_15001(msg *[]byte) ([]byte, int) {
	sc_15001 := &protobuf.SC_15001{
		ItemList:     []*protobuf.ITEMINFO{},
		LimitList:    []*protobuf.ITEMINFO{},
		ItemMiscList: []*protobuf.ITEMMISC{},
	}

	data, err := proto.Marshal(sc_15001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 15001
}
