package constructor

import (
	"log"

	"Shinano/protobuf"
	"Shinano/repository/story"

	"google.golang.org/protobuf/proto"
)

func (b *Builder) Build_13102(msg *[]byte) ([]byte, int) {
	cs := &protobuf.CS_13101{}
	if err := proto.Unmarshal(*msg, cs); err != nil {
		log.Fatal(err)
	}

	team := cs.GetFleet().MainTeam
	id := cs.GetId()
	b.Chapter = story.InitChapter(id)

	curChapter := b.Chapter.StartChapter(team[0].ShipList, team[0].Id)

	sc := &protobuf.SC_13102{
		Result:         proto.Uint32(0),
		CurrentChapter: curChapter,
	}

	data, err := proto.Marshal(sc)
	if err != nil {
		log.Fatal(err)
	}

	return data, 13102
}
