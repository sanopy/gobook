package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

var cities, clocks []string
var mu sync.RWMutex

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: $ %s [[City]=[URL]]...\n", os.Args[0])
		os.Exit(1)
	}

	clocks = make([]string, len(os.Args)-1)

	for i, arg := range os.Args[1:] {
		idx := strings.Index(arg, "=")
		if idx == -1 {
			fmt.Fprintf(os.Stderr, "usage: $ %s [[City]=[URL]]...\n", os.Args[0])
			os.Exit(1)
		}
		cities = append(cities, arg[:idx])
		url := arg[idx+1:]
		go readClock(i, url)
	}

	time.Sleep(1 * time.Second)

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	for _, city := range cities {
		fmt.Fprintf(tw, "%v\t", city)
	}
	fmt.Fprintf(tw, "\n")
	for i := 0; i < len(clocks); i++ {
		mu.RLock()
		fmt.Fprintf(tw, "%v\t", clocks[i])
		mu.RUnlock()
	}
	tw.Flush()

	for {
		fmt.Fprintf(tw, "\r")
		for i := 0; i < len(clocks); i++ {
			mu.RLock()
			fmt.Fprintf(tw, "%v\t", clocks[i])
			mu.RUnlock()
		}
		tw.Flush()
		time.Sleep(50 * time.Millisecond)
	}
}

func readClock(idx int, url string) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		mu.Lock()
		clocks[idx] = string(line)
		mu.Unlock()
	}
}
