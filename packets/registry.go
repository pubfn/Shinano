package packets

import (
	"fmt"
	"net"

	"Shinano/constructor"
)

func Hande(id int, data *[]byte, c net.Conn, idx int) {
	handlers := constructor.Registry[id]

	for _, handler := range handlers {
		data, id := handler(data)
		msg := BuildPacket(&data, id, idx)
		fmt.Println("Sent message:", id)
		c.Write(msg)
	}
}
