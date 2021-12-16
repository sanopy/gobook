package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

func toAscii(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	sr := bytes.NewReader(b)
	r := bufio.NewReader(sr)

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		buf.WriteString(fmt.Sprintf("%s\r\n", line))
	}

	return buf.Bytes(), nil
}
