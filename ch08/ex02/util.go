package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func (c *ftpConn) reply(code int, message string) {
	s := fmt.Sprintf("%d %s\r\n", code, message)
	c.Ctrl.Write([]byte(s))
}

func (c *ftpConn) parse() error {
	r := bufio.NewReader(c.Ctrl)
	line, _, err := r.ReadLine()
	if err != nil {
		return err
	}

	token := strings.Split(string(line), " ")
	c.cmd = strings.ToUpper(token[0])
	c.args = token[1:]

	if c.cmd == "PASS" { // due to security reasons
		log.Print("request: [PASS ********]")
	} else {
		log.Printf("request: [%s]", line)
	}

	return nil
}
