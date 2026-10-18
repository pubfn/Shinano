package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_22001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_22001{
		OilWellLevel:     proto.Uint32(1),
		OilWellLvUpTime:  proto.Uint32(0),
		GoldWellLevel:    proto.Uint32(1),
		GoldWellLvUpTime: proto.Uint32(0),
		ClassLv:          proto.Uint32(1),
		ClassLvUpTime:    proto.Uint32(0),
		Class: &protobuf.NAVALACADEMY_CLASS{
			Proficiency: proto.Uint32(0),
		},
		SkillClassList:     []*protobuf.SKILL_CLASS{},
		SkillClassNum:      proto.Uint32(2),
		DailyFinishBuffCnt: proto.Uint32(0),
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 22001
}
