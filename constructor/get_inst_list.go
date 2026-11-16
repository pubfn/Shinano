package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository"

	"google.golang.org/protobuf/proto"
)

func Build_11706(msg *[]byte) ([]byte, int) {
	ids := protobuf.CS_11705{}

	err := proto.Unmarshal(*msg, &ids)

	sc_11604 := &protobuf.SC_11706{
		InsMessageList: []*protobuf.INS_MESSAGE{},
	}

	for _, idlist := range ids.GetIdList() {
		if instMsg := repository.ParseId(idlist); instMsg != nil {
			sc_11604.InsMessageList = append(sc_11604.InsMessageList, repository.ParseId(idlist))
		}
	}

	data, err := proto.Marshal(sc_11604)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11706
}
