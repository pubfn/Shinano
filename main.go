package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"Shinano/message"
	"Shinano/registry"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		con, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go conn(con)
	}
}

func conn(c net.Conn) {
	reader := bufio.NewReader(c)
	for {
		len, err := reader.Peek(2)
		if err != nil {
			log.Fatal(err)
		}
		data := make([]byte, int(binary.BigEndian.Uint16(len[0:2]))+2)
		_, _ = io.ReadFull(reader, data)
		msg := message.ReadMsg(data)
		fmt.Println("Message Id", msg.Id)
		registry.Hande(msg.Id, &msg.Message, c, msg.Index)
	}
}
