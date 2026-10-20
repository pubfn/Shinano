package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11015(msg *[]byte) ([]byte, int) {
	sc_11015 := &protobuf.SC_11015{
		BuffList: []*protobuf.BENEFITBUFF{},
	}

	data, err := proto.Marshal(sc_11015)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11015
}
