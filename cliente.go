// By: Alfonso Candiani J.
// https://www.candianijaramillo.com
// Cliente del chat-server

package main 

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial ("tcp", "192.168.1.70:6666")//your ip:port
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("Terminamos")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

}