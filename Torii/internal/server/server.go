package server

import (
	"log"
	"net"
)

func StartServer(address string, handler func(net.Conn)) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	log.Println("Server running on", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}

		log.Println("New connection from", conn.RemoteAddr())

		go handler(conn)
	}
}
