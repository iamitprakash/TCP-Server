package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func doSomeProcessing(conn net.Conn) {
	buffer := make([]byte, 1024)
	log.Println("Request Processing")
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep((5 * time.Second))
	conn.Write([]byte("HTTP/1.1 200 Ok\r\n\r\nHello, World!\r\n"))
	conn.Close()
}
func main() {
	listner, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("Waiitng for client to connect")
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(conn)
		log.Println("Client connected")
		//This is a toy TCp server for large number of request few optimization can be done
		//Limiting the number of thread for line #36
		//add threading pool to save on thread creation time
		// connection timeout
		// TCP backlog queue configuration
		go doSomeProcessing(conn)
	}
	fmt.Println("Hello, world!", listner)
}
