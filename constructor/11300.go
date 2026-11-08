package constructor

import (
	"log"

	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

func Build_11300(msg *[]byte) ([]byte, int) {
	sc_11300 := &protobuf.SC_11300{
		NoticeList: []*protobuf.NOTICEINFO_P11{
			{
				Id:         proto.Uint32(6),
				Version:    proto.String("1"),
				BtnTitle:   proto.String("hihihihih"),
				Title:      proto.String("Patch Note"),
				TitleImage: proto.String("https://blhxusstatic.yo-star.com/pic/260326_OA4bigldklYOjgVP5SyeJa.png"),
				TimeDesc:   proto.String("10/67/2026"),
				Content:    proto.String("hahah im in"),
				TagType:    proto.Uint32(1),
				Icon:       proto.Uint32(2),
				Track:      proto.String(""),
				Priority:   proto.Uint32(135),
				NeedLevel:  proto.Uint32(0),
			},
		},
	}

	data, err := proto.Marshal(sc_11300)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11300
}
