// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	addr string
	name string
	ch   chan<- string // an outgoing message channel
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				select {
				case cli.ch <- msg:
				default:
				}
			}

		case cli := <-entering:
			msg := "Now connecting:"
			for c := range clients {
				msg += fmt.Sprintf(" %s,", c.name)
			}
			cli.ch <- msg
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func handleConn(conn net.Conn) {
	requested := make(chan struct{})
	go autoClose(conn, requested)

	var name string
	conn.Write([]byte("name: "))

	input := bufio.NewScanner(conn)
	if input.Scan() {
		name = input.Text()
	}

	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{who, name, ch}

	ch <- "hello " + name
	messages <- name + " has arrived"
	entering <- cli

	for input.Scan() {
		requested <- struct{}{}
		messages <- name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- cli
	messages <- name + " has left"
	conn.Close()
}

func autoClose(c net.Conn, req chan struct{}) {
	for {
		select {
		case <-time.After(5 * time.Minute):
			c.Close()
			return
		case _, ok := <-req:
			if !ok {
				return
			}
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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
