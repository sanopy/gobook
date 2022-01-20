package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"path/filepath"
	"strconv"
	"strings"
)

// TODO: Implement authentication
func (c *ftpConn) handleUser() {
	c.reply(331, "User name okay, need password.") // always return okay
	log.Println("user name ok")
}

// TODO: Implement authentication
func (c *ftpConn) handlePass() {
	c.reply(230, "User logged in, proceed.") // always return okay
	log.Println("user logged in")
}

func (c *ftpConn) handleCwd() {
	path := c.args[0]

	err := c.wd.cd(path)
	if err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("cwd failed: %v", err)
		return
	}

	c.reply(250, "Requested file action okay, completed.")
	log.Println("changed working directory")
}

func (c *ftpConn) handleQuit() {
	c.reply(221, "Service closing control connection.")
	log.Println("client quit")
}

func (c *ftpConn) handlePort() {
	hosts := strings.Split(c.args[0], ",")
	if len(hosts) != 6 {
		c.reply(501, "Syntax error in parameters or arguments.")
		log.Println("port parse failed: invalid number of args")
		return
	}

	domain := fmt.Sprintf("%s.%s.%s.%s", hosts[0], hosts[1], hosts[2], hosts[3])

	p1, err := strconv.Atoi(hosts[4])
	if err != nil {
		c.reply(501, "Syntax error in parameters or arguments.")
		log.Printf("port parse failed: %v", err)
		return
	}
	p2, err := strconv.Atoi(hosts[5])
	if err != nil {
		c.reply(501, "Syntax error in parameters or arguments.")
		log.Printf("port parse failed: %v", err)
		return
	}
	port := p1*256 + p2

	addr := fmt.Sprintf("%s:%d", domain, port)
	c.Data, err = net.Dial("tcp", addr)
	if err != nil {
		c.reply(421, "Service not available, closing control connection.")
		log.Fatal(err)
	}

	c.reply(200, "Command okay.")
	log.Printf("connected to %s", addr)
}

// TODO: Implement ASCII, EBCDIC mode
func (c *ftpConn) handleRetr() {
	path := c.args[0]

	// file open
	r, err := c.wd.open(path)
	if err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	// transfer content
	c.reply(150, "File status okay; about to open data connection.")
	if _, err := io.Copy(c.Data, r); err != nil { // binary mode only
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("io.Copy failed: %v", err)
	}

	c.reply(226, "Closing data connection.")
	c.Data.Close()
}

// TODO: Implement ASCII, EBCDIC mode
func (c *ftpConn) handleStor() {
	path := c.args[0]

	// make directory
	if err := c.wd.mkdir(filepath.Dir(path)); err != nil {
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("mkdir failed: %v", err)
		return
	}
	log.Printf("success mkdir %s", filepath.Dir(path))

	// file open
	f, err := c.wd.create(path)
	if err != nil {
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("file open failed: %v", err)
		return
	}
	defer f.Close()

	// write content
	c.reply(125, "Data connection already open; transfer starting.")
	if _, err := io.Copy(f, c.Data); err != nil { // binary mode only
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("io.Copy failed: %v", err)
		return
	}
	log.Printf("data copied to %s", path)

	c.reply(226, "Closing data connection.")
	c.Data.Close()
}

func (c *ftpConn) handleDele() {
	path := c.args[0]

	f, err := c.wd.open(path)
	if err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	info, err := f.Stat()
	if err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	if info.IsDir() {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	if err := c.wd.remove(path); err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	c.reply(250, "Requested file action okay, completed.")
}

func (c *ftpConn) handleRmd() {
	path := c.args[0]

	f, err := c.wd.open(path)
	if err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	info, err := f.Stat()
	if err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	if !info.IsDir() {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	if err := c.wd.remove(path); err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("file open failed: %v", err)
		return
	}

	c.reply(250, "Requested file action okay, completed.")
}

func (c *ftpConn) handleMkd() {
	path := c.args[0]

	// make directory
	if err := c.wd.mkdir(path); err != nil {
		c.reply(550, "Requested action not taken.")
		log.Printf("mkdir failed: %v", err)
		return
	}
	log.Printf("success mkdir %s", path)

	c.reply(257, fmt.Sprintf("\"%s\" created.", path))
}

func (c *ftpConn) handlePwd() {
	wd := c.wd.pwd()

	c.reply(250, fmt.Sprintf("\"%s\"", wd))
	log.Println("response working directory")
}

func (c *ftpConn) handleList() {
	path := "./"
	if len(c.args) >= 1 {
		path += c.args[0]
	}

	out, err := c.wd.ls(path)
	if err != nil {
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("io.Copy failed: %v", err)
	}

	// write content
	c.reply(125, "Data connection already open; transfer starting.")
	ascii, err := toAscii(out)
	if err != nil {
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("io.Copy failed: %v", err)
	}
	r := bytes.NewReader(ascii)
	if _, err := io.Copy(c.Data, r); err != nil { // binary mode only
		c.reply(451, "Requested action aborted. Local error in processing.")
		log.Printf("io.Copy failed: %v", err)
	}
	log.Println("response list")

	c.reply(226, "Closing data connection.")
	c.Data.Close()
}

func (c *ftpConn) handleSyst() {
	c.reply(215, "UNIX system type.")
	log.Println("response system")
}

func (c *ftpConn) handleNoop() {
	c.reply(200, "Command okay.")
	log.Println("noop")
}

func (c *ftpConn) handleEprt() {
	tokens := strings.Split(c.args[0][1:len(c.args[0])-1], c.args[0][0:1])
	if len(tokens) != 3 {
		c.reply(501, "Syntax error in parameters or arguments.")
		log.Println("port parse failed: invalid number of args")
		return
	}

	var addr string
	switch tokens[0] {
	case "1":
		addr = fmt.Sprintf("%s:%s", tokens[1], tokens[2])
	case "2":
		addr = fmt.Sprintf("[%s]:%s", tokens[1], tokens[2])
	default:
		c.reply(501, "Syntax error in parameters or arguments.")
		log.Println("port parse failed: invalid number of args")
		return
	}

	var err error
	c.Data, err = net.Dial("tcp", addr)
	if err != nil {
		c.reply(421, "Service not available, closing control connection.")
		log.Fatal(err)
	}

	c.reply(200, "Command okay.")
	log.Printf("connected to %s", addr)
}
