package handler

import (
	"log"
	"net"
)

func Echo(conn net.Conn) {
	defer conn.Close()

	log.Println("Handling connection:", conn.RemoteAddr())

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Connection closed:", conn.RemoteAddr())
			return
		}

		msg := string(buf[:n])
		log.Println("Received:", msg)

		conn.Write(buf[:n])
	}
}
