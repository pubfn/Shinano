package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_62101(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_62101{
		Technologys: []*protobuf.GUILD_TECHNOLOGY{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 62101
}
