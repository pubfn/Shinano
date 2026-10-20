package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11200(msg *[]byte) ([]byte, int) {
	sc_11200 := &protobuf.SC_11200{
		ActivityList: []*protobuf.ACTIVITYINFO{},
		HxList:       []uint32{},
	}

	data, err := proto.Marshal(sc_11200)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11200
}
