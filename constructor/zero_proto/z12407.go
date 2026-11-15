package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12407(msg *[]byte) ([]byte, int) {
	sc_12407 := &protobuf.SC_12407{
		ConvertedList:       []*protobuf.PT_OLD_LOVER_ITEM{},
		RewardedList:        []uint32{},
		MedalList:           []*protobuf.PT_LOVE_LETTER_MEDAL{},
		LetterList:          []*protobuf.PT_SHIP_LOVE_LETTER{},
		ConvertedLetterList: []*protobuf.PT_SHIP_LOVE_LETTER{},
	}

	data, err := proto.Marshal(sc_12407)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12407
}
