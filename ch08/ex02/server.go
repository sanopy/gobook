package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type ftpConn struct {
	Ctrl net.Conn
	Data net.Conn
	wd   Wd
	cmd  string
	args []string
}

const port = 2100

var initialDir string

func main() {
	var err error
	initialDir, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(&ftpConn{Ctrl: conn, wd: Wd(initialDir)})
	}
}

func handleConn(c *ftpConn) {
	defer c.Ctrl.Close()

	c.reply(220, "Service ready for new user.")
	log.Println("server ready")

	for {
		err := c.parse()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("request parse failed: %v", err)
		}

		switch c.cmd {
		case "USER":
			c.handleUser()
		case "PASS":
			c.handlePass()
		case "CWD":
			c.handleCwd()
		case "QUIT":
			c.handleQuit()
		case "PORT":
			c.handlePort()
		// case "TYPE":
		// case "MODE":
		// case "STRU":
		case "RETR":
			c.handleRetr()
		case "STOR":
			c.handleStor()
		case "PWD":
			c.handlePwd()
		case "LIST":
			c.handleList()
		case "SYST":
			c.handleSyst()
		case "NOOP":
			c.handleNoop()
		case "EPRT":
			c.handleEprt()
		default:
			c.reply(500, "Syntax error, command unrecognized.")
			log.Println("command unrecognized")
		}
	}
}
