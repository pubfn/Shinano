package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_14001(msg *[]byte) ([]byte, int) {
	sc_14001 := &protobuf.SC_14001{
		EquipList:       []*protobuf.EQUIPINFO{},
		ShipIdList:      []uint32{},
		SpweaponList:    []*protobuf.SPWEAPONINFO{},
		SpweaponBagSize: proto.Uint32(150),
	}

	data, err := proto.Marshal(sc_14001)
	if err != nil {
		log.Fatal(err)
	}

	return data, 14001
}
