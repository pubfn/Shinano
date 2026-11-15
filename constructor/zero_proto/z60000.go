package zero_proto

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_60000(msg *[]byte) ([]byte, int) {
	sc_50000 := &protobuf.SC_60000{
		Guild: &protobuf.GUILD_INFO{
			Base: &protobuf.GUILD_BASE_INFO{
				Id:              proto.Uint32(0),
				Policy:          proto.Uint32(1),
				Faction:         proto.Uint32(1),
				Name:            proto.String(""),
				Level:           proto.Uint32(1),
				Announce:        proto.String(""),
				Manifesto:       proto.String(""),
				Exp:             proto.Uint32(0),
				MemberCount:     proto.Uint32(0),
				ChangeFactionCd: proto.Uint32(0),
				KickLeaderCd:    proto.Uint32(0),
			},

			Member: []*protobuf.MEMBER_INFO{},
			Log:    []*protobuf.LOG_INFO{},

			GuildEx: &protobuf.GUILD_EXPANSION_INFO{
				Capital: proto.Uint32(0),

				ThisWeeklyTasks: &protobuf.WEEKLY_TASK{
					Id:            proto.Uint32(0),
					Progress:      proto.Uint32(0),
					Monday_0Clock: proto.Uint32(0),
				},

				BenefitFinishTime:     proto.Uint32(0),
				Technologys:           []*protobuf.GUILD_TECHNOLOGY{},
				RetreatCnt:            proto.Uint32(0),
				TechCancelCnt:         proto.Uint32(0),
				LastBenefitFinishTime: proto.Uint32(0),
				ActiveEventCnt:        proto.Uint32(0),
			},
		},
	}

	data, err := proto.Marshal(sc_50000)
	if err != nil {
		log.Fatal(err)
	}

	return data, 60000
}
