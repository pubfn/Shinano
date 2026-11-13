package constructor

import (
	"log"

	"Shinano/cfg"
	"Shinano/protobuf"

	"google.golang.org/protobuf/proto"
)

var (
	idCount uint32 = 0
	msgTime uint32 = 1774512000
)

func Build_11706(msg *[]byte) ([]byte, int) {
	ids := protobuf.CS_11705{}

	err := proto.Unmarshal(*msg, &ids)

	sc_11604 := &protobuf.SC_11706{
		InsMessageList: []*protobuf.INS_MESSAGE{},
	}

	for _, idlist := range ids.GetIdList() {
		if instMsg := parseId(idlist); instMsg != nil {
			sc_11604.InsMessageList = append(sc_11604.InsMessageList, parseId(idlist))
		}
	}

	data, err := proto.Marshal(sc_11604)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11706
}

func parseId(id uint32) *protobuf.INS_MESSAGE {
	data := cfg.InstMsgID[id]
	if data.Text == "" {
		return nil
	}

	playerDisc := &protobuf.INS_PLAYER{
		Id:       proto.Uint32(id),
		Time:     proto.Uint32(0),
		TextList: []string{},
		Text:     proto.String(""),
		NpcReply: []uint32{},
	}
	npcDisc := []*protobuf.INS_NPC{}
	npcRepl := []*protobuf.INS_NPC{}

	for _, playerText := range data.Op {
		playerDisc.TextList = append(playerDisc.TextList, playerText.Discuss)
	}

	for _, npcDiscuss := range data.Discuss {
		idCount++
		disc := &protobuf.INS_NPC{
			Id:       proto.Uint32(idCount),
			Time:     proto.Uint32(msgTime),
			Text:     &npcDiscuss.Discuss,
			NpcReply: []uint32{},
		}

		for _, repl := range npcDiscuss.Reply {
			idCount++
			reply := &protobuf.INS_NPC{
				Id:       proto.Uint32(idCount),
				Time:     proto.Uint32(msgTime),
				Text:     &repl,
				NpcReply: []uint32{},
			}
			disc.NpcReply = append(disc.NpcReply, idCount)

			npcRepl = append(npcRepl, reply)
		}
		npcDisc = append(npcDisc, disc)
	}

	return &protobuf.INS_MESSAGE{
		Id:            proto.Uint32(id),
		Time:          proto.Uint32(msgTime),
		Text:          &data.Text,
		Picture:       proto.String(""),
		NpcDiscuss:    npcDisc,
		NpcReply:      npcRepl,
		PlayerDiscuss: []*protobuf.INS_PLAYER{playerDisc},
		Good:          proto.Uint32(322),
		IsGood:        proto.Uint32(0),
		IsRead:        proto.Uint32(1),
		OalistPic:     proto.String(""),
	}
}
