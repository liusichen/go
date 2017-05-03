package main

import "net"
import "io"
import "log"
import "time"

import "flag"

var port = flag.String("port", "8000", "the port server use")

func main() {
	ipaddr := "localhost"
	flag.Parse()

	tcpHead := ipaddr + ":" + *port
	listener, err := net.Listen("tcp", tcpHead)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(connect net.Conn) {
	defer connect.Close()
	for {
		_, err := io.WriteString(connect, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
