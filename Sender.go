package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:2556")
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	buf, err := time.Now().MarshalBinary()
	conn.Write(buf)
	checkError(err)

	os.Exit(0)
}

//Handles error
func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal error ")
		os.Exit(1)
	}
}
