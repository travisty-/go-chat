package main

import (
	"flag"

	"github.com/travisty-/go-chat/chat"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	connIP := flag.Arg(0)

	if isHost {
		// go run main.go -listen <ip>
		chat.RunHost(connIP)
	} else {
		// go run main.go <ip>
		chat.RunGuest(connIP)
	}
}
