package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_29002(msg *[]byte) ([]byte, int) {
	sc_29002 := &protobuf.SC_29002{
		Result: proto.Uint32(0),
		Tb: &protobuf.TBINFO{
			Id: proto.Uint32(0),
			Fsm: &protobuf.TBFSM{
				SystemNo:     proto.Uint32(0),
				CurrentNode:  proto.Uint32(0),
				Cache:        []*protobuf.TBFSMCACHE{},
				PriorityFsm:  []*protobuf.TBFSM{},
				TarotSelects: []uint32{},
			},
			Round: &protobuf.TBROUND{
				Round:     proto.Uint32(0),
				InTemp:    proto.Uint32(0),
				TempRound: proto.Uint32(0),
			},
			Res: &protobuf.TBRES{
				Attrs:    []*protobuf.KVDATA{},
				Resource: []*protobuf.KVDATA{},
			},
			Talent: &protobuf.TBTALENT{
				Talents: []uint32{},
			},
			Plan: &protobuf.TBPLAN{
				PlanUpgrade: []uint32{},
			},
			Site: &protobuf.TBSITE{
				Characters:   []uint32{},
				WorkCounter:  []*protobuf.KVDATA{},
				Works:        []uint32{},
				EventCounter: []*protobuf.KVDATA{},
			},
			Evaluations: []*protobuf.KVDATA{},
			Name:        proto.String(""),
			FavorLv:     proto.Uint32(0),
			Benefit: &protobuf.TBBENEFIT{
				Actives: []*protobuf.TBBF{},
			},
			Difficulty: proto.Uint32(0),
			EvalFail:   proto.Uint32(0),
			Display: &protobuf.TBDISPLAY{
				BenefitDisplay:   []*protobuf.TBDROP{},
				DollarNumDisplay: []*protobuf.TBBENEFITVAL{},
				Counter:          []*protobuf.TBBFCOUNTER{},
			},
		},
		Permanent: &protobuf.TBPERMANENT{
			NgPlusCount:   proto.Uint32(0),
			Polaroids:     []uint32{},
			Endings:       []uint32{},
			ActiveEndings: []uint32{},
			TarotArchive:  []uint32{},
			MaxRound:      proto.Uint32(0),
		},
	}
	data, err := proto.Marshal(sc_29002)
	if err != nil {
		log.Fatal(err)
	}

	return data, 29002
}
