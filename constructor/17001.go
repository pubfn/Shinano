package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_17001(msg *[]byte) ([]byte, int) {
	sc := &protobuf.SC_17001{
		FinishList: []uint32{},
		ProgressList: []*protobuf.ACHIEVEMENT_INFO{
			{Id: proto.Uint32(10011), Progress: proto.Uint32(2), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10021), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10031), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10041), Progress: proto.Uint32(14), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10051), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10061), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100611), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100621), Progress: proto.Uint32(5014), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100631), Progress: proto.Uint32(992), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10071), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100711), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100721), Progress: proto.Uint32(19010), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100731), Progress: proto.Uint32(1986), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10081), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100811), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100821), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100831), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10091), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100911), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100921), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(100931), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(10101), Progress: proto.Uint32(8), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200001), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200011), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200021), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200031), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200041), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200051), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200061), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200071), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200081), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200091), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200101), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200111), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200121), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200131), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200141), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200151), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200161), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200171), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200181), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200191), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200201), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200211), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200221), Progress: proto.Uint32(0), Timestamp: proto.Uint32(0)},
			{Id: proto.Uint32(200231), Progress: proto.Uint32(1), Timestamp: proto.Uint32(0)},
		},
		ShipInfoList: []*protobuf.SHIP_STATISTICS_INFO{
			{Id: proto.Uint32(20121), Star: proto.Uint32(2), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(60), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(9000), LvMax: proto.Uint32(5)},
			{Id: proto.Uint32(10601), Star: proto.Uint32(2), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(68), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5014), LvMax: proto.Uint32(4)},
			{Id: proto.Uint32(20124), Star: proto.Uint32(2), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(14), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5010), LvMax: proto.Uint32(4)},
			{Id: proto.Uint32(20110), Star: proto.Uint32(1), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(39), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5000), LvMax: proto.Uint32(1)},
			{Id: proto.Uint32(10307), Star: proto.Uint32(2), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(45), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5000), LvMax: proto.Uint32(1)},
			{Id: proto.Uint32(30303), Star: proto.Uint32(1), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(20), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5000), LvMax: proto.Uint32(1)},
			{Id: proto.Uint32(10303), Star: proto.Uint32(2), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(16), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5000), LvMax: proto.Uint32(1)},
			{Id: proto.Uint32(20402), Star: proto.Uint32(2), HeartFlag: proto.Uint32(0), HeartCount: proto.Uint32(44), MarryFlag: proto.Uint32(0), IntimacyMax: proto.Uint32(5000), LvMax: proto.Uint32(1)},
		},
		ShipAwardList: []*protobuf.SHIP_STATISTICS_AWARD{},
		DailyDiscuss:  proto.Uint32(0),
		TransformList: []uint32{},
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 17001
}
