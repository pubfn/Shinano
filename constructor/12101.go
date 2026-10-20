package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12101(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12101{
		GroupList: []*protobuf.GROUPINFO_P12{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12101
}
