package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunHost takes an IP as an argument
// and listens for connections on that IP.
func RunHost(ip string) {
	ipAndPort := fmt.Sprintf("%s:%s", ip, port)

	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error: ", listenErr)
	}
	fmt.Println("Listening on", ipAndPort)

	defer listener.Close()

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error: ", acceptErr)
	}
	fmt.Println("New connection accepted")

	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewReader(os.Stdin)

	for {
		getMessage(conn, *reader)
		sendMessage(conn, *writer)
	}
}

// RunGuest takes a destination IP as an
// argument and connects to that IP.
func RunGuest(ip string) {
	ipAndPort := fmt.Sprintf("%s:%s", ip, port)

	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error: ", dialErr)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewReader(os.Stdin)

	for {
		sendMessage(conn, *writer)
		getMessage(conn, *reader)
	}
}

func getMessage(conn net.Conn, reader bufio.Reader) {
	message, receiveErr := reader.ReadString('\n')
	if receiveErr != nil {
		log.Fatal("Error: ", receiveErr)
	}
	fmt.Println("Message received: ", message)
}

func sendMessage(conn net.Conn, writer bufio.Reader) {
	fmt.Print("Send message: ")
	message, sendErr := writer.ReadString('\n')
	if sendErr != nil {
		log.Fatal("Error: ", sendErr)
	}
	fmt.Fprint(conn, message)
}
