package server

import (
	"log"

	"github.com/minero/minero/constants"
	"github.com/minero/minero/proto/auth"
	"github.com/minero/minero/proto/packet"
	"github.com/minero/minero/server/player"
)

// Handle02 handles incoming requests of packet 0x02: Handshake
func Handle02(server *Server, sender *player.Player) {
	pkt := new(packet.Handshake)
	pkt.ReadFrom(sender.Conn)

	log.Printf("Handshake from: %q [%s]", pkt.Username, sender.RemoteAddr())

	if pkt.Version != constants.ProtoNum {
		log.Printf("Wrong Protocol version. Player: %d, Server: %d\n",
			pkt.Version, constants.ProtoNum)
		return
	}

	// Save player to list
	sender.Name = pkt.Username
	server.AddPlayer(sender)

	log.Println("online_mode =", server.config.Get("server.online_mode"))

	if server.config.Get("server.online_mode") == "true" {
		// Succesful handshake, prepare Encryption Request
		r := packet.EncryptionKeyRequest{
			ServerId:  server.Id(),
			PublicKey: server.PublicKey(),
			Token:     auth.EncryptionBytes(),
		}
		r.WriteTo(sender.Conn)
		sender.Token = r.Token
	} else {
		// BUG(toqueteos): server: Add online_mode=false support.
	}
}
