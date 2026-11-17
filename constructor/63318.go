package constructor

import (
	"fmt"
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_63318(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_63317{}

	if err := proto.Unmarshal(*msg, cs); err != nil {
		log.Fatal(err)
	}

	for _, v := range cs.GetShipIdList() {
		fmt.Println(v)
	}

	sc_14101 := &protobuf.SC_63318{
		InfoList: []*protobuf.META_SKILL_SIMPLE_INFO{},
	}

	data, err := proto.Marshal(sc_14101)
	if err != nil {
		log.Fatal(err)
	}

	return data, 63318
}
