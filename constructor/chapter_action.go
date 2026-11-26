package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/story"

	"google.golang.org/protobuf/proto"
)

func Build_13104(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_13103{}

	proto.Unmarshal(*msg, cs)

	movePath := story.StoryAct(cs.GetAct(), cs.GetActArg_1(), cs.GetActArg_2())

	sc := &protobuf.SC_13104{
		Result:           proto.Uint32(0),
		MovePath:         movePath,
		DropList:         []*protobuf.DROPINFO{},
		MapUpdate:        []*protobuf.CHAPTERCELLINFO_P13{},
		ShipUpdate:       []*protobuf.SHIPINCHAPTER_P13{},
		AiList:           []*protobuf.CHAPTERCELLINFO_P13{},
		AiActList:        []*protobuf.AI_ACT_P13{},
		SubmarineActList: []*protobuf.AI_ACT_P13{},
		EscortActList:    []*protobuf.AI_ACT_P13{},
		FleetActList:     []*protobuf.AI_ACT_P13{},
		BuffList:         []uint32{},
		AddFlagList:      []uint32{},
		DelFlagList:      []uint32{},
		CellFlagList:     []*protobuf.CELLFLAG{},
		ExtraDropList:    []*protobuf.DROPINFO{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13104
}
