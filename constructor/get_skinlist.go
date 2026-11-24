package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/ship"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12201(msg *[]byte) ([]byte, int) {
	skinList := []*protobuf.IDTIMEINFO{}

	skinsIds, err := ship.GetAllSkins(b.DB)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range skinsIds {
		skin := &protobuf.IDTIMEINFO{
			Id:   proto.Uint32(v),
			Time: proto.Uint32(0),
		}

		skinList = append(skinList, skin)
	}

	sc := &protobuf.SC_12201{
		SkinList:          skinList,
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
