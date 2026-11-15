package zero_proto

import (
	"log"
	"time"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11002(msg *[]byte) ([]byte, int) {
	sc_11002 := &protobuf.SC_11002{
		Timestamp:               proto.Uint32(uint32(time.Now().Unix())),
		Monday_0OclockTimestamp: proto.Uint32(1606114800),
		ShipCount:               proto.Uint32(0), //  just do 0. there check.
	}

	data, err := proto.Marshal(sc_11002)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11002
}
