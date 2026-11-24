package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/inventory"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_15001(msg *[]byte) ([]byte, int) {
	itemList, err := inventory.GetInventory(b.DB)
	if err != nil {
		log.Fatal(err)
	}
	sc_15001 := &protobuf.SC_15001{
		ItemList:     itemList,
		LimitList:    []*protobuf.ITEMINFO{},
		ItemMiscList: []*protobuf.ITEMMISC{},
	}

	data, err := proto.Marshal(sc_15001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 15001
}
