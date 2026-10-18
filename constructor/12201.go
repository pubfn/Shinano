package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12201(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12201{
		SkinList:          []*protobuf.IDTIMEINFO{},
		ForbiddenSkinList: []uint32{},
		ForbiddenSkinType: []uint32{},
		ForbiddenList:     []*protobuf.SKIN_FORBIDDEN{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12201
}
