package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11203(msg *[]byte) ([]byte, int) {
	sc_11200 := &protobuf.SC_11203{
		Result:         proto.Uint32(0),
		AwardList:      []*protobuf.DROPINFO{},
		Build:          []*protobuf.BUILDINFO{},
		Number:         []uint32{},
		ReturnUserList: []*protobuf.RETURN_USER_INFO{},
		CollectionList: []*protobuf.COLLECTIONINFO{},
		TaskList:       []*protobuf.TASKINFO{},
	}

	data, err := proto.Marshal(sc_11200)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11203
}
