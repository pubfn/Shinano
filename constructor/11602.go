package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11602(msg *[]byte) ([]byte, int) {
	sc_11752 := &protobuf.SC_11602{
		EmojiList: []uint32{},
	}

	data, err := proto.Marshal(sc_11752)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11602
}
