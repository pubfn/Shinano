package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11003(msg *[]byte) ([]byte, int) {
	sc_11003 := &protobuf.SC_11003{
		Id:    proto.Uint32(478112442),
		Name:  proto.String("Nigger"),
		Level: proto.Uint32(3),
		Exp:   proto.Uint32(20),
		ResourceList: []*protobuf.RESOURCE{
			{
				Type: proto.Uint32(14),
				Num:  proto.Uint32(1488),
			},
			{
				Type: proto.Uint32(1),
				Num:  proto.Uint32(1337),
			},
		},
		AttackCount: proto.Uint32(2),
		WinCount:    proto.Uint32(2),
		Adv:         proto.String(""),
		Character: []*protobuf.KVDATA{
			{
				Key:   proto.Uint32(1),
				Value: proto.Uint32(0),
			},
		},
		ShipBagMax:         proto.Uint32(150),
		EquipBagMax:        proto.Uint32(300),
		GmFlag:             proto.Uint32(0),
		Rank:               proto.Uint32(1),
		PvpAttackCount:     proto.Uint32(0),
		CollectAttackCount: proto.Uint32(0),
		GuideIndex:         proto.Uint32(1),
		BuyOilCount:        proto.Uint32(0),
		ChatRoomId:         proto.Uint32(1),
		CardList:           []*protobuf.CARDINFO{},
		MaxRank:            proto.Uint32(0),
		RegisterTime:       proto.Uint32(1790705327),
		ShipCount:          proto.Uint32(8),
		AccPayLv:           proto.Uint32(0),
		StoryList: []uint32{
			90019,
			90018,
			2400002,
			2400001,
			90049,
			2,
			1,
			90012,
		},
		GuildWaitTime:  proto.Uint32(0),
		ChatMsgBanTime: proto.Uint32(0),
		FlagList:       []uint32{1000002},
		MedalId:        []uint32{},
		IconFrameList:  []*protobuf.IDTIMEINFO{},
		ChatFrameList:  []*protobuf.IDTIMEINFO{},
		Display: &protobuf.DISPLAYINFO{
			Icon:          proto.Uint32(201211),
			Skin:          proto.Uint32(0),
			IconFrame:     proto.Uint32(0),
			ChatFrame:     proto.Uint32(0),
			IconTheme:     proto.Uint32(0),
			MarryFlag:     proto.Uint32(0),
			TransformFlag: proto.Uint32(0),
		},
		Rmb: proto.Uint32(0),
		Appreciation: &protobuf.APPRECIATIONINFO{
			Gallerys:      []uint32{},
			Musics:        []uint32{},
			FavorGallerys: []uint32{},
			FavorMusics:   []uint32{},
			MusicNo:       proto.Uint32(0),
			MusicMode:     proto.Uint32(0),
		},
		ThemeUploadNotAllowedTime: proto.Uint32(0),
		RefundShopInfoList:        []*protobuf.REFUND_SHOPINFO{},
		CartoonReadMark:           []uint32{},
		CartoonCollectMark:        []uint32{},
		RandomShipMode:            proto.Uint32(0),
		RandomShipList:            []*protobuf.KVDATA{},
		MarryShip:                 proto.Uint32(0),
		TakingShipList:            []*protobuf.SHIP_TAKING_DATA{},
		Soundstory:                []uint32{},
		ChildDisplay:              proto.Uint32(0),
		Cover: &protobuf.LIVINGAREA_COVER{
			Id:     proto.Uint32(0),
			Covers: []uint32{},
		},
		MailStoreroomLv: proto.Uint32(1),
		BattleUiList:    []uint32{},
		PvpWinCount:     proto.Uint32(0),
		CdList:          []*protobuf.COOLDOWN{},
		CommanderBagMax: proto.Uint32(0),
		BattleUi:        proto.Uint32(0),
		ActivityMedals: []*protobuf.KVDATA{
			{
				Key:   proto.Uint32(5115402),
				Value: proto.Uint32(1791366312),
			},
		},
		NewGuideIndex: proto.Uint32(19),
	}

	data, err := proto.Marshal(sc_11003)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11003
}
