package constructor

import (
	"log"
	"time"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_10801(date *[]byte) ([]byte, int) {
	time := time.Now().UTC().Unix()

	sc_10801 := &protobuf.SC_10801{
		GatewayIp:   proto.String("blhxusgate.yo-star.com"),
		GatewayPort: proto.Uint32(80),
		Url:         proto.String("https://play.google.com/store/apps/details?id=com.YoStarEN.AzurLane"),
		Version: []string{
			"$azhash$9$3$386$25468c0968cd7359",
			"$cvhash$846$28d5465204659433",
			"$l2dhash$967$6eae69163599c07d",
			"$pichash$41$45514169a0d52eff",
			"$bgmhash$35$0d91c3fc54424fd2",
			"$paintinghash$664$9149a172c89700a2",
			"$mangahash$146$782e623940d73fef",
			"$cipherhash$83$8b7a7d757503adfb",
			"$dormhash$303$17b3cf8706ab77ec",
			"$maphash$97$af5821919b63a11b",
			"count-5",
			"dTag-1",
		},
		ProxyIp:                 proto.String("blhxusproxy.yo-star.com"),
		ProxyPort:               proto.Uint32(20000),
		IsTs:                    proto.Uint32(0),
		Timestamp:               proto.Uint32(uint32(time)),
		Monday_0OclockTimestamp: proto.Uint32(1606114800),
		CdnList:                 []string{},
	}

	data, err := proto.Marshal(sc_10801)
	if err != nil {
		log.Fatal(err)
	}

	return data, 10801
}
