package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/commander"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_10023(msg *[]byte) ([]byte, int) {
	fields := &protobuf.CS_10022{}

	err := proto.Unmarshal(*msg, fields)
	if err != nil {
		log.Fatal(err)
	}

	accountID := int64(fields.GetAccountId())

	if err = commander.CheckCommander(b.DB, accountID); err != nil {
		log.Fatal(err)
	}

	sc_10023 := &protobuf.SC_10023{
		UserId:       proto.Uint32(uint32(accountID)),
		Result:       proto.Uint32(0),
		ServerTicket: proto.String("0"),
		DbLoad:       nil,
		ServerLoad:   nil,
	}

	data, err := proto.Marshal(sc_10023)
	if err != nil {
		log.Fatal(err)
	}

	return data, 10023
}
