package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

func server1() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)

		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}
func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}
	c.Close()
}
func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hello, World"
	fmt.Println("Sending message", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}
func main() {
	go server1()
	go client()
	go server()
	go client1()
	var input string
	fmt.Scanln(&input)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
		<head>
		<title>Hello, World</title>
		</head>
		<body>
		Hello, World!
		</body>
		</html>`,
	)
}

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}
func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}

}
func client1() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999)=", result)
	}

}
