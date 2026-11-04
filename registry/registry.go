package registry

import (
	"fmt"
	"net"

	"Shinano/constructor"
	"Shinano/message"
)

func Hande(id int, data *[]byte, c net.Conn, idx int) {
	handlers := constructor.Registry[id]

	for _, handler := range handlers {
		data, id := handler(data)
		msg := message.BuildPacket(&data, id, idx)
		fmt.Println("Sent message:", id)
		c.Write(msg)
	}
}
