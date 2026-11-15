package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_50000(msg *[]byte) ([]byte, int) {
	sc_50000 := &protobuf.SC_50000{
		FriendList:  []*protobuf.FRIEND_INFO{},
		RequestList: []*protobuf.MSG_INFO_P50{},
	}

	data, err := proto.Marshal(sc_50000)
	if err != nil {
		log.Fatal(err)
	}

	return data, 50000
}
