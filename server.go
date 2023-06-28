package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

const (
	SNET   = "tcp"
	SADDR  = "0.0.0.0:8000"
	SDELIM = '\n'
)

func main() {
	listener, err := net.Listen(SNET, SADDR)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleRequest(conn)
	}
}

// handleRequest ...
func handleRequest(conn net.Conn) {
	defer conn.Close()
	cReader := bufio.NewReader(conn)

	for {
		// Waiting for the client request
		cReq, err := cReader.ReadString(SDELIM)

		switch err {
		case nil:
			cReq := strings.TrimSpace(cReq)
			log.Printf("received from - %s: %s\n", conn.RemoteAddr(), cReq)

		case io.EOF:
			log.Println("connection closed by: ", conn.RemoteAddr())
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		if _, err = conn.Write([]byte("Server Received\n")); err != nil {
			log.Printf("unable to respond to %s: %v\n", conn.RemoteAddr(), err)
		}
	}
}
