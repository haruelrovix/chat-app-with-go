package main

import (
	"flag"
	"fmt"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listen on the specified IP address")
	flag.Parse()

	if isHost {
		fmt.Println("is host")
	} else {
		fmt.Println("is guest")
	}
}
