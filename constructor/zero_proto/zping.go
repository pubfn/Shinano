package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_10101(msg *[]byte) ([]byte, int) {
	sc_10100 := &protobuf.SC_10101{
		State: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_10100)
	if err != nil {
		log.Fatal(err)
	}

	return data, 10101
}

func Build_11020(msg *[]byte) ([]byte, int) {
	sc_11018 := &protobuf.SC_11020{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_11018)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11020
}

func Build_12046(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_12046{
		Result: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12046
}
