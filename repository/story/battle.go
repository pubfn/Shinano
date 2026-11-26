package story

import "Shinano/protobuf"

func GetMvp(statistic []*protobuf.STATISTICSINFO) uint32 {
	var mvp uint32

	var buf uint32
	for _, ship := range statistic {
		if ship.GetDamageCause() > buf {
			buf = ship.GetDamageCause()
			mvp = ship.GetShipId()
		}
	}

	return mvp
}
