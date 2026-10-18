package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_20201(msg *[]byte) ([]byte, int) {
	sc_20201 := &protobuf.SC_20201{
		Info: []*protobuf.ACT_TASK_INIT_LIST{
			{
				ActId: proto.Uint32(7502),
				Tasks: []*protobuf.ACT_TASK{
					{Id: proto.Uint32(23846), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23847), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23848), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23849), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23850), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23851), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23852), Progress: proto.Uint32(0)},
				},
				FinishIds: []uint32{},
			},
			{
				ActId: proto.Uint32(7501),
				Tasks: []*protobuf.ACT_TASK{
					{Id: proto.Uint32(20611), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(20612), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(20613), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(20614), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(20615), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(20616), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(20617), Progress: proto.Uint32(0)},
				},
				FinishIds: []uint32{},
			},
			{
				ActId: proto.Uint32(7505),
				Tasks: []*protobuf.ACT_TASK{
					{Id: proto.Uint32(23920), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23921), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23922), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23923), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23924), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23925), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(23926), Progress: proto.Uint32(0)},
				},
				FinishIds: []uint32{},
			},
			{
				ActId: proto.Uint32(7503),
				Tasks: []*protobuf.ACT_TASK{
					{Id: proto.Uint32(24481), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(24482), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(24483), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(24484), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(24485), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(24486), Progress: proto.Uint32(0)},
					{Id: proto.Uint32(24487), Progress: proto.Uint32(0)},
				},
				FinishIds: []uint32{},
			},
		},
	}

	data, err := proto.Marshal(sc_20201)
	if err != nil {
		log.Fatal(err)
	}

	return data, 20201
}
