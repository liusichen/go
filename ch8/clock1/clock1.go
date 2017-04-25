package main

import "net"
import "io"
import "log"
import "time"

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(connect net.Conn) {
	defer connect.Close()
	for {
		_, err := io.WriteString(connect, time.Now().Format("15:04:01\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
