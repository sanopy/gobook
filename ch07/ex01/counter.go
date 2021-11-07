package counter

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	in := bufio.NewScanner(bytes.NewReader(p))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		*c++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	in := bufio.NewScanner(bytes.NewReader(p))
	in.Split(bufio.ScanLines)
	for in.Scan() {
		*c++
	}
	return len(p), nil
}
