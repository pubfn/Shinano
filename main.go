package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"Shinano/cfg"
	"Shinano/constructor"
	"Shinano/db"
	"Shinano/packets"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	builder := constructor.Builder{DB: db.Init()}
	constructor.NewRegistry(&builder)
	err = cfg.LoadGameData()
	if err != nil {
		log.Fatal(err)
	}

	for {
		con, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go conn(con)
	}
}

func conn(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		len, err := reader.Peek(2)
		if err != nil {
			log.Println("error:", err)
			return
		}

		data := make([]byte, int(binary.BigEndian.Uint16(len[0:2]))+2)
		_, err = io.ReadFull(reader, data)
		if err != nil {
			log.Println("error:", err)
			return
		}

		msg, err := packets.ReadPacket(data)
		if err != nil {
			log.Println("error:", err)
			return
		}

		fmt.Println("Message Id", msg.Id)
		packets.Hande(msg.Id, &msg.Message, c, msg.Index)
	}
}
