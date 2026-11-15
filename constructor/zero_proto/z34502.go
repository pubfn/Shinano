package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_34502(msg *[]byte) ([]byte, int) {
	sc_33114 := &protobuf.SC_34502{
		FightCount:           proto.Uint32(0),
		FightCountUpdateTime: proto.Uint32(0),
		OtherBoss:            []*protobuf.WORLDBOSS_INFO_P34{},
		SummonPt:             proto.Uint32(0),
		SummonPtOld:          proto.Uint32(0),
		SummonPtDailyAcc:     proto.Uint32(0),
		SummonPtOldDailyAcc:  proto.Uint32(0),
		SummonFree:           proto.Uint32(0),
		AutoFightFinishTime:  proto.Uint32(0),
		DefaultBossId:        proto.Uint32(0),
		AutoFightMaxDamage:   proto.Uint32(0),
		GuildSupport:         proto.Uint32(0),
		FriendSupport:        proto.Uint32(0),
		WorldSupport:         proto.Uint32(0),
		SelfBossLv:           proto.Uint32(0),
	}

	data, err := proto.Marshal(sc_33114)
	if err != nil {
		log.Fatal(err)
	}

	return data, 34502
}
