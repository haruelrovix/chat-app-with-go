package lib

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = "8080"

// RunHost takes an IP as argument
// and listens for connections on that IP
func RunHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error", listenErr)
	}
	fmt.Println("Listening on ", ipAndPort)

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error", acceptErr)
	}
	fmt.Println("New connections accepted")

	for {
		handleHost(conn)
	}
}

func handleHost(conn net.Conn) {
	reader := bufio.NewReader(conn)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error", readErr)
	}
	fmt.Println("Message received: ", message)

	fmt.Print("Send message: ")
	replyReader := bufio.NewReader(os.Stdin)
	replyMessage, replyErr := replyReader.ReadString('\n')
	if replyErr != nil {
		log.Fatal("Error", replyErr)
	}
	fmt.Fprint(conn, replyMessage)
}

// RunGuest takes destination IP as argument
// and connects to that IP
func RunGuest(ip string) {
	ipAndPort := ip + ":" + port
	conn, dialErr := net.Dial("tcp", ipAndPort)
	if dialErr != nil {
		log.Fatal("Error", dialErr)
	}
	fmt.Print("Send message: ")
	reader := bufio.NewReader(os.Stdin)
	message, readErr := reader.ReadString('\n')
	if readErr != nil {
		log.Fatal("Error", readErr)
	}
	fmt.Fprint(conn, message)
}
