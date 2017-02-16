package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/ecs0/OCWTF/traffic"
)

func main() {
	// Create packet instance
	var a traffic.AckPacket
	conn := Send("127.0.0.1:13337")

	lis, err := net.Listen("tcp", conn.LocalAddr().String())
	CheckError(err)
	// Close the lister after function main return
	defer lis.Close()

	con, err := lis.Accept()
	CheckError(err)
	decoder := json.NewDecoder(con)

	decoder.Decode(&a)

	fmt.Print(a, "\n")

}

func Send(SinkAddr string) net.Conn {
	connection, _ := net.Dial("tcp", SinkAddr)
	encoder := json.NewEncoder(connection)

	p := traffic.IntentionPacketFactory{"abc123"}
	Intent := p.Build(5, 20, 20)

	err := encoder.Encode(Intent)
	if err != nil {
		log.Fatal(err)
	}

	return connection
}

// check for error
func CheckError(err error) {
	if err != nil {
		fmt.Print("Error: ", err)
	}
}
