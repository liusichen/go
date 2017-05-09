package main

import (
	"io"
	"net"
	"os"
	"time"
)

func main() {
	argv := os.Args[1:]

	for _, local := range argv {
		go func() {
			conn, _ := net.Dial("tcp", local)
			defer conn.Close()
			mustCopy(os.Stdout, conn)
		}()
	}
	time.Sleep(1e11)
}
func mustCopy(dst io.Writer, src io.Reader) {
	io.Copy(dst, src)
}
