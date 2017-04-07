package main

import (
	"flag"
	"os"

	"github.com/travisty-/go-chat/chat"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		chat.RunHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		chat.RunGuest(connIP)
	}
}
