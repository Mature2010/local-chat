// By: Alfonso Candiani J.
// https://www.candianijaramillo.com

package main 

import (
	"fmt"
	"bufio"
	"log"
	"net"
)

type client chan <- string

var (
	entrantes = make(chan client)
	salientes = make(chan client)
	mensajes = make(chan string)
)

func broadcaster() {

	clients := make(map[client]bool)

	for {
		select {
		case msg := <- mensajes:
			for cli := range clients{
				cli <- msg
			}

		case cli := <- entrantes:
			clients[cli] = true

		case cli := <- salientes:
			delete(clients, cli)
			close(cli)
		}
	}

}

func handleConn (conn net.Conn) {
	ch := make (chan string)

	go clientWriter(conn, ch)

	quien := conn.RemoteAddr().String()
	ch <- "Tú eres " + quien

	mensajes <- quien + "se ha conectado a la sala"
	entrantes <- ch

	input := bufio.NewScanner(conn)
	
	for input.Scan(){
		mensajes <- quien + ":" + input.Text()
	}

	salientes <- ch
	mensajes <- quien + " se ha ido"
	conn.Close()
}


func clientWriter(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "192.168.1.70:6666")//your ip:port
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
