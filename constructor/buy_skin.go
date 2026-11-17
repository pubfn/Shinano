package constructor

import (
	"fmt"
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_16002(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_16001{}

	proto.Unmarshal(*msg, cs)

	fmt.Println(cs.GetId())

	sc := &protobuf.SC_16002{
		Result:   proto.Uint32(0),
		DropList: []*protobuf.DROPINFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 16002
}
