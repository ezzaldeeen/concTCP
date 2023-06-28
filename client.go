package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	CNET   = "tcp"
	CADDR  = "0.0.0.0:8000"
	CDELIM = '\n'
)

func main() {
	conn, err := net.Dial(CNET, CADDR)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	cReader := bufio.NewReader(os.Stdin)
	sReader := bufio.NewReader(conn)

	for {
		// waiting client request
		cReq, err := cReader.ReadString(CDELIM)

		switch err {
		case nil:
			cReq := strings.TrimSpace(cReq)
			b, err := conn.Write([]byte(cReq + "\n"))
			if err != nil {
				log.Printf("unable to send the client request: %v\n", err)
			}
			log.Printf("%s -> %s\n", conn.LocalAddr(), conn.RemoteAddr())
			log.Printf("# of written bytes: %d\n", b)
		case io.EOF:
			log.Println("client closed the connection")
			return
		default:
			log.Printf("client error: %v\n", err)
			return
		}

		// waiting server response
		sRes, err := sReader.ReadString(CDELIM)

		switch err {
		case nil:
			log.Printf("recevied: %s\n", strings.TrimSpace(sRes))
		case io.EOF:
			log.Println("server closed the connection")
			return
		default:
			log.Printf("server error: %v\n", err)
			return
		}
	}
}
