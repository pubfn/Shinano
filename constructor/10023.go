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

	accountID := int64(fields.GetAccountId())

	err = b.DB.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM commander WHERE id = ?)",
		accountID,
	).Scan(&exists)

	if !exists {
		db.NewCommader(int(accountID), b.DB)
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
