package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp4", ":2556")
	checkError(err)
	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	for {
		handleClient(conn)
	}
}

//Deals with the traffic comming in through the socket
func handleClient(conn *net.UDPConn) {
	var buf []byte = make([]byte, 512)
	Size, Addr, err := conn.ReadFrom(buf)
	if err != nil {
		return
	}

	// Get the timeStamp and calculate the trip time
	sentTime := new(time.Time)
	sentTime.UnmarshalBinary(buf[:Size])
	Trip := time.Now().Sub(*sentTime)

	//outputs
	fmt.Printf("Confirmation of receiving packet from %+v\n", Addr)
	fmt.Printf("Packet seceived At: %s\n", time.Now().String())
	fmt.Printf("Package sent At: %+v\n", sentTime)
	fmt.Printf("Traveling time is: %+v\n", Trip)

}

// function to deal with error
func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal error ")
		os.Exit(1)
	}
}
