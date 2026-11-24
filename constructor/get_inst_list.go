package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/insta"

	"google.golang.org/protobuf/proto"
)

func Build_11706(msg *[]byte) ([]byte, int) {
	ids := protobuf.CS_11705{}

	err := proto.Unmarshal(*msg, &ids)

	sc_11604 := &protobuf.SC_11706{
		InsMessageList: []*protobuf.INS_MESSAGE{},
	}

	for _, idlist := range ids.GetIdList() {
		if instMsg := insta.ParseId(idlist); instMsg != nil {
			sc_11604.InsMessageList = append(sc_11604.InsMessageList, insta.ParseId(idlist))
		}
	}

	data, err := proto.Marshal(sc_11604)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11706
}
