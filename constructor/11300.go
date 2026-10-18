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
				BtnTitle:   proto.String("10/15 Update"),
				Title:      proto.String("Patch Note"),
				TitleImage: proto.String("https://blhxusstatic.yo-star.com/notice/2026101501.png"),
				TimeDesc:   proto.String("10/14/2026"),
				Content: proto.String(`<size=35>List of New Contents</size>

		<size=35>War Archives Update</size>
		    Tempesta and the Sleeping Sea has been added to the War Archives.
		    Tempesta and Islas de Libertád has been added to the War Archives.

		<size=35>New Event</size>
		    Available for a limited time between 10/15 - 11/4, 11:59 P.M. (UTC-7)
		        The Strict Maid Isn't So Strict (2 event missions unlock each day until 10/21. Complete all 7 days' worth of missions to get Algérie's limited skin, "The Strict Maid Isn't So Strict".)

		<size=35>Construction Pools Update</size>
		    Available permanently in Light Construction Pool
		        Lyme
		        Royal James
		        Fancy
		        Portsmouth Adventure

		    Available permanently in Heavy Construction Pool
		        Queen Anne's Revenge
		        Pearl
		        Ganj-i-Sawai

		    Available permanently in Special Construction Pool
		        Amity

		<size=35>Rerun Skins</size>
		    Available permanently
		        Pearl - Fallen Angel in the Dark Castle (L2D+)
		        Queen Anne's Revenge - Call of the Creeping Embrace (L2D)
		        Lyme - Corrupting Trap (Dynamic)
		        Adventure - Master of the Safe House (Dynamic)
		        Kronshtadt - Animal Agent (Dynamic)
		        Izumo - Tail-Stirred Cocktail (Dynamic)
		        Royal James - The Forbidden Slurp
		        Harutsuki - Little Steamy Sprite
		        Owari - My Wish is For Love (L2D)
		        Ganj-i-Sawai - Mermaid in Distress (Dynamic)
		        Amity - Friendly Franken (Dynamic)
		        Atago - Full Moon's Feral Wolf (Dynamic)
		        Fancy - Grave of Your Liking
		        Portsmouth Adventure - Adventurous Surprise
		        Dolphin - Incantation of Revival
		        Bristol - "Head" Investigator

		<size=35>New Feature</size>
		    Added the Fleetbuilding Guide function to some Main Campaign stages. Commanders can tap it on the corresponding chapter screen to view recommended fleet and gears;
		        ※Tap on a character to view a recommended example gear setup. Feel free to make adjustments based on the shipgirls you have as well as their stats.
		        ※Fleet and gear recommendations are for reference only.

		<size=35>Rerun Furniture</size>
		    Available permanently
		        Temple of Phantom Light
		        Shadowy Port

		<size=35>New Items in Shop</size>
		    Available for a limited time between 10/15 - 10/21, 11:59 P.M. (UTC-7)
		        Limited Build Supplies (1 purchase limited)
		        Limited Strategic Supplies (5 purchases limited)
		        Decor Tokens Pack (2 purchases limited)
		        Cognitive Chip Pack (2 purchases limited)

		<size=35>Core Data Shop Update</size>
		    Available permanently in Core Shop (Ltd.)
		        Meowfficer - Teach
		        Meowfficer - Bonnet
		        Meowfficer - Every
		        Meowfficer - Roberts

		<size=35>Merit Shop Update</size>
		    Available permanently
		        Adventure
		        Dolphin

		<size=35>Private Quarters Update</size>
		    Available permanently
		        Admiral Nakhimov's exclusive furniture - Nyakhimov's Dream Home

		    Available for a limited time between 10/15 - 10/21, 11:59 P.M. (UTC-7)
		        Taihou's exclusive furniture - Floralized Rhythm

		<size=35>CV Update</size>
		    Skin
		        Algérie - The Strict Maid Isn't So Strict - Shizuka Ito

		<size=35>System Optimization</size>
		1. Comics were updated;
		2. Added a play-while-downloading feature, allowing game content to be accessed before all game resources have been downloaded;
		3. Fixed an issue where the facial expression variations in Kearsarge's skin "Springtime Monitoring" differed between the dynamic artwork and static artwork;
		4. Fixed an issue where game resources could be downloaded abnormally under certain circumstances.`),
				TagType:   proto.Uint32(1),
				Icon:      proto.Uint32(2),
				Track:     proto.String(""),
				Priority:  proto.Uint32(135),
				NeedLevel: proto.Uint32(0),
			},
		},
	}

	data, err := proto.Marshal(sc_11300)
	if err != nil {
		log.Fatal(err)
	}

	return data, 11300
}
