package main

import (
	"log"
	"net"

	"github.com/minero/minero/proto/packet"
	"github.com/minero/minero/proto/ping"
)

func Server(addr string) {
	log.SetPrefix("serverlistdebug> ")
	log.SetFlags(log.Ltime)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}

		// Serve multiple connections concurrently.
		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()

	log.Println("Got connection from:", c.RemoteAddr())

	// Read first two bytes NMC sends
	var buf = make([]byte, 2)
	_, err := c.Read(buf)
	if err != nil {
		log.Printf("%s: %s\n", c.RemoteAddr(), err)
		return
	}

	// Equal to 0xFE?
	if buf[0] != packet.PacketServerListPing || buf[1] != 0x01 {
		return
	}

	// Send response
	p := ping.Ping(Flags[:])
	p.WriteTo(c)
}
