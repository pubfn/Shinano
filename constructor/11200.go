package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11200(msg *[]byte) ([]byte, int) {
	sc_11200 := &protobuf.SC_11200{
		ActivityList: []*protobuf.ACTIVITYINFO{
			{
				Id: proto.Uint32(1000140), StopTime: proto.Uint32(1792652399),
				Data1: nil, Data2: nil, Data3: nil, Data4: nil,
				Data1List: []uint32{}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
			{
				Id: proto.Uint32(1000112), StopTime: proto.Uint32(1792047599),
				Data1: nil, Data2: nil, Data3: nil, Data4: nil,
				Data1List: []uint32{}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
			{
				Id: proto.Uint32(1000111), StopTime: proto.Uint32(1792652399),
				Data1: nil, Data2: nil, Data3: nil, Data4: nil,
				Data1List: []uint32{}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
			{
				Id: proto.Uint32(1000110), StopTime: proto.Uint32(1792047599),
				Data1: proto.Uint32(0), Data2: nil, Data3: nil, Data4: nil,
				Data1List: []uint32{}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
			{
				Id: proto.Uint32(1000109), StopTime: proto.Uint32(1792047599),
				Data1: proto.Uint32(2), Data2: proto.Uint32(1791874800),
				Data3: nil, Data4: nil,
				Data1List: []uint32{2, 1}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
			{
				Id: proto.Uint32(1000108), StopTime: proto.Uint32(1792047599),
				Data1: proto.Uint32(0), Data2: proto.Uint32(0),
				Data3: proto.Uint32(1), Data4: nil,
				Data1List: []uint32{}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
			{
				Id: proto.Uint32(1000107), StopTime: proto.Uint32(1792047599),
				Data1: proto.Uint32(3), Data2: proto.Uint32(0),
				Data3: nil, Data4: nil,
				Data1List: []uint32{}, Data2List: []uint32{}, Data3List: []uint32{},
				Date1KeyValueList: []*protobuf.KEYVALUELIST_P11{},
				GroupList:         []*protobuf.GROUPINFO_P11{},
				CollectionList:    []*protobuf.COLLECTIONINFO{},
				TaskList:          []*protobuf.TASKINFO{},
				BuffList:          []*protobuf.BENEFITBUFF{},
				StrData1:          nil, Data4List: []uint32{},
			},
		},
		HxList: []uint32{},
	}

	data, err := proto.Marshal(sc_11200)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11200
}
