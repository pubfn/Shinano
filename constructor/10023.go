package constructor

import (
	"log"

	"Shinano/db"
	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_10023(msg *[]byte) ([]byte, int) {
	var exists bool
	fields := &protobuf.CS_10022{}

	err := proto.Unmarshal(*msg, fields)
	if err != nil {
		log.Fatal(err)
	}

	err = b.DB.QueryRow("SELECT EXISTS(SELECT id FROM commander WHERE id = ?)", fields.GetAccountId()).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	if !exists {
		db.NewCommader(int(fields.GetAccountId()), b.DB)
	}

	sc_10023 := &protobuf.SC_10023{
		UserId:       proto.Uint32(fields.GetAccountId()),
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
