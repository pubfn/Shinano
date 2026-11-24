package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/gacha"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_12044(msg *[]byte) ([]byte, int) {
	infoList := []*protobuf.BUILD_INFO{}

	buildInfo, err := gacha.GetBuildInfo(b.DB)
	if err != nil {
		log.Fatal("gacha result", err)
	}

	for _, v := range buildInfo {
		infoList = append(infoList, &protobuf.BUILD_INFO{
			Pos: &v.Pos,
			Tid: &v.ShipId,
		})
	}

	sc := &protobuf.SC_12044{
		InfoList: infoList,
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12044
}
