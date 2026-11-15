package constructor

import (
	"fmt"
	"log"
	"time"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_12003(msg *[]byte) ([]byte, int) {
	var i int
	buildData := protobuf.CS_12002{}
	curTime := time.Now().Unix()
	if err := proto.Unmarshal(*msg, &buildData); err != nil {
		log.Fatal(err)
	}

	sc_12002 := &protobuf.SC_12003{
		Result:    proto.Uint32(0),
		BuildInfo: []*protobuf.BUILDINFO{},
	}

	buildId := buildData.GetId()
	fmt.Println("count ", buildData.GetCount())
	for range buildData.GetCount() {
		fmt.Println("func raz skolko", i)
		build := protobuf.BUILDINFO{
			Time:       proto.Uint32(uint32(curTime)),
			FinishTime: proto.Uint32(uint32(curTime + 3)),
			BuildId:    &buildId,
		}
		i++
		sc_12002.BuildInfo = append(sc_12002.BuildInfo, &build)
	}

	data, err := proto.Marshal(sc_12002)
	if err != nil {
		log.Fatal(err)
	}

	return data, 12003
}
