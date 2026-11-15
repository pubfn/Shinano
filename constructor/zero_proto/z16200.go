package zero_proto

import (
	"log"
	"time"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_16200(msg *[]byte) ([]byte, int) {
	sc_16200 := &protobuf.SC_16200{
		CoreShopList:   []*protobuf.SHOPINFO{},
		BlueShopList:   []*protobuf.SHOPINFO{},
		NormalShopList: []*protobuf.SHOPINFO{},
		Month:          proto.Uint32(uint32(time.Now().Month())),
	}

	data, err := proto.Marshal(sc_16200)
	if err != nil {
		log.Fatal(err)
	}

	return data, 16200
}
