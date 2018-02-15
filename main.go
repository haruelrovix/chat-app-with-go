package main

import (
	"flag"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listen on the specified IP address")
	flag.Parse()

	if isHost {
		// go run main.go -listen <ip>
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIP)
	}
}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	_, listenErr := net.Listen("tcp", ipAndPort)
	if listenErr != nil {
		log.Fatal("Error", listenErr)
	}
}

func runGuest(ip string) {

}
